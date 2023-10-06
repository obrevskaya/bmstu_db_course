package postgres

import (
	"context"
	"fmt"

	logicModels "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/logic/models"
	repoModels "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/repository/postgres/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BillRepository struct {
	db *gorm.DB
}

func NewBR(db *gorm.DB) BillRepository {
	return BillRepository{db: db}
}

func (b BillRepository) Insert(ctx context.Context, bill *logicModels.Bill) error {
	bill.ID = uuid.New()
	billDB := &repoModels.Bill{
		ID:      bill.ID,
		IDOrder: bill.IDOrder,
		Price:   bill.Price,
		Status:  bill.Status,
	}
	res := b.db.WithContext(ctx).Table("bills").Create(billDB)
	if res.Error != nil {
		return fmt.Errorf("insert: %w", res.Error)
	}
	return nil
}

func (b BillRepository) Get(ctx context.Context, ID uuid.UUID) (*logicModels.Bill, error) {
	bill := repoModels.Bill{}

	res := b.db.WithContext(ctx).Table("bills").Where("id = ?", ID).Take(&bill)
	if res.Error != nil {
		return nil, fmt.Errorf("select: %w", res.Error)
	}

	resBill := logicModels.Bill{
		ID:      bill.ID,
		IDOrder: bill.IDOrder,
		Price:   bill.Price,
		Status:  bill.Status,
	}

	return &resBill, nil
}

func (b BillRepository) UpdateBillStatus(ctx context.Context, ID uuid.UUID, status string, points int) error {
	err := b.db.Transaction(func(tx *gorm.DB) error {
		billDB := &repoModels.Bill{}

		res := tx.WithContext(ctx).Table("bills").Where("id = ?", ID).Take(&billDB)
		if res.Error != nil {
			return fmt.Errorf("select: %w", res.Error)
		}

		billDB.Status = status
		res = tx.WithContext(ctx).Table("bills").Save(billDB)
		if res.Error != nil {
			return fmt.Errorf("update: %w", res.Error)
		}

		order := &repoModels.Order{}
		res = tx.WithContext(ctx).Table("orders").Where("id = ?", billDB.IDOrder).Take(order)
		if res.Error != nil {
			return fmt.Errorf("get order: %w", res.Error)
		}

		if order.IsPoints == false {
			user := &repoModels.User{}
			res = tx.WithContext(ctx).Table("users").Where("id = ?", order.IDUser).Take(user)
			if res.Error != nil {
				return fmt.Errorf("get user: %w", res.Error)
			}
			user.Points = points
			res = tx.WithContext(ctx).Table("users").Save(user)
			if res.Error != nil {
				return fmt.Errorf("update user points: %w", res.Error)
			}
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("error transaction: %w", err)
	}
	return nil

}
