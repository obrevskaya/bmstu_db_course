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

type UserWineRepository struct {
	db *gorm.DB
}

func NewUWR(db *gorm.DB) UserWineRepository {
	return UserWineRepository{db: db}
}

func (uw UserWineRepository) Insert(ctx context.Context, IDUser uuid.UUID, IDWine uuid.UUID) error {
	userWine := &repoModels.UserWine{
		IDUser: IDUser,
		IDWine: IDWine,
	}
	res := uw.db.WithContext(ctx).Table("user_wines").Create(userWine)
	if res.Error != nil {
		return fmt.Errorf("insert: %w", res.Error)
	}
	return nil
}

func (uw UserWineRepository) DeleteWine(ctx context.Context, IDUser uuid.UUID, IDWine uuid.UUID) error {
	res := uw.db.WithContext(ctx).Table("user_wines").Where("id_user = ? and id_wine = ?", IDUser, IDWine).Delete(&repoModels.UserWine{})
	if res.Error != nil {
		return fmt.Errorf("delete: %w", res.Error)
	}

	return nil
}

func (uw UserWineRepository) GetByUser(ctx context.Context, IDUser uuid.UUID) ([]*logicModels.UserWine, error) {
	var userWinesDB []*repoModels.UserWine

	res := uw.db.WithContext(ctx).Table("user_wines").Where("id_user = ?", IDUser).Find(&userWinesDB)
	if res.Error != nil {
		return nil, fmt.Errorf("select: %w", res.Error)
	}

	userWinesLogic := make([]*logicModels.UserWine, 0, len(userWinesDB))
	for _, userWineOld := range userWinesDB {
		userWine := &logicModels.UserWine{}
		err := copier.Copy(userWine, userWineOld)
		if err != nil {
			return nil, err
		}

		userWinesLogic = append(userWinesLogic, userWine)
	}

	return userWinesLogic, nil
}
