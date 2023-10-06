package menu

import (
	"context"
	"errors"
	"fmt"
	myСontext "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/context"
	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/interfaces"
	"github.com/dixonwille/wmenu"
	"github.com/fatih/color"
)

var (
	errExit = errors.New("exit")
)

type Menu struct {
	mainMenu          *wmenu.Menu
	billController    interfaces.IBillController
	orderController   interfaces.IOrderController
	elementController interfaces.IOrderElementController
	userController    interfaces.IUserController
	wineController    interfaces.IWineController

	login    string
	password string
}

func (m *Menu) AddOptions(ctx context.Context) {
	m.mainMenu.Option("Register", ctx, false, m.registerHandler)
	m.mainMenu.Option("Authorize", ctx, false, m.authorizeHandler)
	m.mainMenu.Option("Log out", ctx, false, m.logoutHandler)
	m.mainMenu.Option("Get wines", ctx, false, m.getWinesHandler)
	m.mainMenu.Option("Create order", ctx, false, m.createOrderHandler)
	m.mainMenu.Option("Create element of order with wine", ctx, false, m.createElementHandler)
	m.mainMenu.Option("Add element of order with wine", ctx, false, m.addElementHandler)
	m.mainMenu.Option("Decrease element of order with wine", ctx, false, m.decreaseElementHandler)
	m.mainMenu.Option("Delete element of order with wine", ctx, false, m.deleteElementHandler)
	m.mainMenu.Option("Get order", ctx, false, m.getOrderHandler)
	m.mainMenu.Option("Place an order", ctx, false, m.placeOrderHandler)
	m.mainMenu.Option("Confirm pay bill", ctx, false, m.payBillHandler)
	m.mainMenu.Option("Add wine", ctx, false, m.addWineHandler)
	m.mainMenu.Option("Delete wine", ctx, false, m.deleteWineHandler)
	m.mainMenu.Option("Update wine", ctx, false, m.updateWineHandler)
	m.mainMenu.Option("Update user points", ctx, false, m.updateUserPointsHandler)
	m.mainMenu.Option("Exit", ctx, false, func(_ wmenu.Opt) error {
		return errExit
	})
}

func NewMenu(b interfaces.IBillController, o interfaces.IOrderController, el interfaces.IOrderElementController,
	u interfaces.IUserController, w interfaces.IWineController) *Menu {
	return &Menu{
		billController:    b,
		orderController:   o,
		elementController: el,
		userController:    u,
		wineController:    w,
	}
}

func (m *Menu) RunMenu(ctx context.Context) {
	m.mainMenu = wmenu.NewMenu("Choose an option.")
	m.AddOptions(ctx)

	for {
		err := m.mainMenu.Run()
		fmt.Println()
		if err != nil {
			if errors.Is(err, errExit) {
				break
			}

			c := color.New(color.FgRed)
			c.Printf("ERROR: %s\n\n", err)
		}
	}

	c := color.New(color.FgMagenta)
	c.Printf("Exited menu.\n")
	myСontext.LoggerFromContext(ctx).Infow("exited menu")
}
