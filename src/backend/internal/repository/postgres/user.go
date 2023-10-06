package postgres

import (
	"context"
	"fmt"

	logicModels "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/logic/models"
	repoModels "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/repository/postgres/models"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUR(db *gorm.DB) UserRepository {
	return UserRepository{db: db}
}

func (u UserRepository) Insert(ctx context.Context, user *logicModels.User) error {
	user.ID = uuid.New()
	userDB := &repoModels.User{
		ID:       user.ID,
		Login:    user.Login,
		Password: user.Password,
		Fio:      user.Fio,
		Email:    user.Email,
		Points:   user.Points,
		Status:   user.Status,
	}
	res := u.db.WithContext(ctx).Table("users").Create(userDB)
	if res.Error != nil {
		return fmt.Errorf("insert: %w", res.Error)
	}
	return nil
}

func (u UserRepository) Authorize(ctx context.Context, login string, password string) (*logicModels.User, error) {
	user := repoModels.User{}
	var count int64
	res := u.db.WithContext(ctx).Table("users").Where("login = ? and password = ?", login, password).Take(&user).Count(&count)
	if res.Error != nil {
		return nil, fmt.Errorf("select: %w", res.Error)
	}
	if count != 1 {
		return nil, fmt.Errorf("error authorize")
	}
	resUser := logicModels.User{
		ID:       user.ID,
		Login:    user.Login,
		Password: user.Password,
		Fio:      user.Fio,
		Email:    user.Email,
		Points:   user.Points,
		Status:   user.Status,
	}

	return &resUser, nil
}

func (u UserRepository) Get(ctx context.Context, ID uuid.UUID) (*logicModels.User, error) {
	user := repoModels.User{}

	res := u.db.WithContext(ctx).Table("users").Where("id = ?", ID).Take(&user)
	if res.Error != nil {
		return nil, fmt.Errorf("select: %w", res.Error)
	}

	resUser := logicModels.User{
		ID:       user.ID,
		Login:    user.Login,
		Password: user.Password,
		Fio:      user.Fio,
		Email:    user.Email,
		Points:   user.Points,
		Status:   user.Status,
	}

	return &resUser, nil

}

func (u UserRepository) UpdateUserPoints(ctx context.Context, ID uuid.UUID, points int) error {
	userDB := &repoModels.User{}
	userDBOld, err := u.Get(ctx, ID)
	if err != nil {
		return fmt.Errorf("update: %w", err)
	}
	err = copier.Copy(userDB, userDBOld)
	if err != nil {
		return fmt.Errorf("update: %w", err)
	}
	userDB.Points = points
	res := u.db.WithContext(ctx).Table("users").Save(userDB)
	if res.Error != nil {
		return fmt.Errorf("update: %w", res.Error)
	}

	return nil
}
