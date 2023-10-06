package postgres

import (
	"context"
	"errors"
	"fmt"

	myErrors "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/logic/errors"
	logicModels "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/logic/models"
	repoModels "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/repository/postgres/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOR(db *gorm.DB) OrderRepository {
	return OrderRepository{db: db}
}

func (o OrderRepository) Insert(ctx context.Context, order *logicModels.Order) error {

	order.ID = uuid.New()
	orderDB := &repoModels.Order{
		ID:         order.ID,
		IDUser:     order.IDUser,
		TotalPrice: order.TotalPrice,
		IsPoints:   order.IsPoints,
		Status:     order.Status,
	}
	res := o.db.WithContext(ctx).Table("orders").Create(orderDB)
	if res.Error != nil {
		return fmt.Errorf("insert: %w", res.Error)
	}
	return nil
}

func (o OrderRepository) Get(ctx context.Context, ID uuid.UUID) (*logicModels.Order, error) {
	order := repoModels.Order{}

	res := o.db.WithContext(ctx).Table("orders").Where("id = ?", ID).Take(&order)
	if res.Error != nil {
		if errors.Is(res.Error, myErrors.ErrNotFound) {
			return nil, myErrors.ErrNotFound
		}
		return nil, fmt.Errorf("select: %w", res.Error)
	}

	resOrder := logicModels.Order{
		ID:         order.ID,
		IDUser:     order.IDUser,
		TotalPrice: order.TotalPrice,
		IsPoints:   order.IsPoints,
		Status:     order.Status,
	}

	return &resOrder, nil
}

func (o OrderRepository) GetByUserInProcess(ctx context.Context, ID uuid.UUID) (*logicModels.Order, error) {
	order := repoModels.Order{}

	res := o.db.WithContext(ctx).Table("orders").Where("id_user = ?", ID).Where("status = ?", logicModels.ProcessOrder).Take(&order)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, myErrors.ErrNotFound
		}
		return nil, fmt.Errorf("select: %w", res.Error)
	}

	resOrder := logicModels.Order{
		ID:         order.ID,
		IDUser:     order.IDUser,
		TotalPrice: order.TotalPrice,
		IsPoints:   order.IsPoints,
		Status:     order.Status,
	}

	return &resOrder, nil
}

func (o OrderRepository) Delete(ctx context.Context, ID uuid.UUID) error {
	res := o.db.WithContext(ctx).Table("orders").Where("id = ?", ID).Delete(&repoModels.Order{})
	if res.Error != nil {
		return fmt.Errorf("delete: %w", res.Error)
	}

	return nil
}

func (o OrderRepository) Update(ctx context.Context, order *logicModels.Order) error {
	err := o.db.Transaction(func(tx *gorm.DB) error {
		if order.Status == logicModels.PlacedOrder {

			var elems []*repoModels.OrderElement

			res := tx.WithContext(ctx).Table("order_elements").Where("id_order = ?", order.ID).Find(&elems)
			if res.Error != nil {
				return fmt.Errorf("select: %w", res.Error)
			}

			for _, el := range elems {
				wine := repoModels.Wine{}

				res := tx.WithContext(ctx).Table("wines").Where("id = ?", el.IDWine).Take(&wine)

				if res.Error != nil {
					return fmt.Errorf("error get %d wine: %w", el.IDWine, res.Error)
				}

				if wine.Count < el.Count {
					return fmt.Errorf("error cnt %s wines: %w", wine.Name, myErrors.ErrCntWine)
				}
				wine.Count -= el.Count

				res = tx.WithContext(ctx).Table("wines").Save(wine)
				if res.Error != nil {
					return fmt.Errorf("update: %w", res.Error)
				}
			}

		}

		orderDB := &repoModels.Order{
			ID:         order.ID,
			IDUser:     order.IDUser,
			TotalPrice: order.TotalPrice,
			IsPoints:   order.IsPoints,
			Status:     order.Status,
		}

		res := tx.WithContext(ctx).Table("orders").Save(orderDB)
		if res.Error != nil {
			return fmt.Errorf("update: %w", res.Error)
		}

		if order.Status == logicModels.PlacedOrder {
			bill := &repoModels.Bill{
				ID:      uuid.New(),
				IDOrder: order.ID,
				Price:   order.TotalPrice,
				Status:  logicModels.NotPaidBill,
			}

			res := tx.WithContext(ctx).Table("bills").Create(bill)
			if res.Error != nil {
				return fmt.Errorf("insert: %w", res.Error)
			}

			if order.IsPoints == true {
				user := &repoModels.User{}
				tx.WithContext(ctx).Table("users").Where("id = ?", order.IDUser).Take(user)
				if user.Points < order.TotalPrice {
					return fmt.Errorf("not enough points for order")
				}
				user.Points -= order.TotalPrice
				tx.WithContext(ctx).Table("users").Save(user)

			}

		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("error transaction: %w", err)
	}
	return nil
}
