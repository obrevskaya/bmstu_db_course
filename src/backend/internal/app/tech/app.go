package tech

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	myContext "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/context"
	controllers2 "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/logic/controllers"
	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/logic/models"
	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/menu"
	postgres2 "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/repository/postgres"
	"github.com/fatih/color"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

type App struct {
	db     *gorm.DB
	menu   *menu.Menu
	cfg    Config
	logger *zap.SugaredLogger
}

func New() *App {
	return &App{}
}

func (a *App) readConfig() error {
	cfgFile := flag.String("cfg", "./configs/config.yaml", "config file name")

	flag.Parse()

	viper.SetConfigName(filepath.Base(*cfgFile))
	viper.SetConfigType("yaml")
	viper.AddConfigPath(filepath.Dir(*cfgFile))

	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("read in config: %w", err)
	}

	err = viper.Unmarshal(&a.cfg)
	if err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}

	return nil
}

func (a *App) initLogger() error {
	lvl, err := zap.ParseAtomicLevel(a.cfg.Logger.Level)
	if err != nil {
		return fmt.Errorf("parse level: %w", err)
	}

	logConfig := zap.Config{
		Level:    lvl,
		Encoding: a.cfg.Logger.Encoding,
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "message",
			LevelKey:     "level",
			TimeKey:      "time",
			CallerKey:    "caller",
			EncodeLevel:  zapcore.CapitalLevelEncoder,
			EncodeTime:   zapcore.RFC3339TimeEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{a.cfg.Logger.File},
		ErrorOutputPaths: []string{a.cfg.Logger.File},
	}

	logger, err := logConfig.Build()
	if err != nil {
		return fmt.Errorf("build logger: %w", err)
	}

	a.logger = logger.Sugar()

	return nil
}

func (a *App) initAdmin(u *controllers2.UserController) error {
	_, err := u.Authorize(context.Background(), "admin", "password")
	if err == nil {
		return nil
	}

	admin := models.User{
		Login:    "admin",
		Password: "password",
		Points:   0,
		Status:   models.Admin,
	}

	fmt.Print("\nEnter admin name: ")
	if _, err := fmt.Scan(&admin.Fio); err != nil {
		return fmt.Errorf("input name: %w", err)
	}

	fmt.Print("Enter admin email: ")
	if _, err := fmt.Scan(&admin.Email); err != nil {
		return fmt.Errorf("input email: %w", err)
	}

	err = u.Create(context.Background(), &admin)
	if err != nil {
		return fmt.Errorf("create: %w", err)
	}

	c := color.New(color.FgHiRed)
	c.Printf("Admin:\nlogin - %s\npassword - %s\n\n", admin.Login, admin.Password)

	return nil
}

func (a *App) Init() error {
	err := a.readConfig()
	if err != nil {
		return fmt.Errorf("read config: %w", err)
	}

	err = a.initLogger()
	if err != nil {
		return fmt.Errorf("init logger: %w", err)
	}

	db, err := gorm.Open(postgres.Open(a.cfg.PG.toDSN()), &gorm.Config{
		Logger: zapgorm2.New(a.logger.Desugar()),
	})
	if err != nil {
		a.logger.Fatalw("cannot open gorm connection", "error", err)
		return fmt.Errorf("gorm open: %w", err)
	}

	a.db = db

	billRep := postgres2.NewBR(db)
	orderRep := postgres2.NewOR(db)
	elemRep := postgres2.NewOElR(db)
	userRep := postgres2.NewUR(db)
	wineRep := postgres2.NewWR(db)

	billController := controllers2.NewBillController(billRep, userRep, orderRep)
	orderController := controllers2.NewOrderController(billRep, elemRep, userRep, orderRep, wineRep)
	elemController := controllers2.NewElemController(elemRep, orderRep, wineRep, orderController)
	userController := controllers2.NewUserController(userRep)
	wineController := controllers2.NewWineController(wineRep)

	err = a.initAdmin(userController)
	if err != nil {
		a.logger.Fatalw("cannot init admin", "error", err)
		return fmt.Errorf("init admin: %w", err)
	}

	a.menu = menu.NewMenu(billController, orderController, elemController, userController, wineController)

	return nil
}

func (a *App) Run(ctx context.Context) {

	a.logger.Infow("started running application")
	ctx = myContext.LoggerToContext(ctx, a.logger)
	a.menu.RunMenu(ctx)
}
