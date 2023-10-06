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

type OrderElRepository struct {
	db *gorm.DB
}

func NewOElR(db *gorm.DB) OrderElRepository {
	return OrderElRepository{db: db}
}

func (el OrderElRepository) Insert(ctx context.Context, elem *logicModels.OrderElement, order *logicModels.Order) error {
	err := el.db.Transaction(func(tx *gorm.DB) error {
		elem.ID = uuid.New()
		orderElemDB := &repoModels.OrderElement{}
		err := copier.Copy(orderElemDB, elem)
		if err != nil {
			return err
		}

		res := tx.WithContext(ctx).Table("order_elements").Create(orderElemDB)
		if res.Error != nil {
			return fmt.Errorf("insert: %w", res.Error)
		}

		orderDB := &repoModels.Order{}
		err = copier.Copy(orderDB, order)
		if err != nil {
			return fmt.Errorf("copy order: %w", err)
		}

		res = tx.WithContext(ctx).Table("orders").Save(orderDB)
		if res.Error != nil {
			return fmt.Errorf("update: %w", res.Error)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("error transaction: %w", err)
	}
	return nil
}

func (el OrderElRepository) GetByID(ctx context.Context, ID uuid.UUID) (*logicModels.OrderElement, error) {
	elem := &repoModels.OrderElement{}

	res := el.db.WithContext(ctx).Table("order_elements").Where("id = ?", ID).Take(elem)
	if res.Error != nil {
		return nil, fmt.Errorf("select: %w", res.Error)
	}

	resElem := &logicModels.OrderElement{}
	err := copier.Copy(resElem, elem)
	if err != nil {
		return nil, err
	}
	return resElem, nil
}

func (el OrderElRepository) GetByOrder(ctx context.Context, IDOrder uuid.UUID) ([]*logicModels.OrderElement, error) {
	var elemsDB []*repoModels.OrderElement

	res := el.db.WithContext(ctx).Table("order_elements").Where("id_order = ?", IDOrder).Find(&elemsDB)
	if res.Error != nil {
		return nil, fmt.Errorf("select: %w", res.Error)
	}

	elemsLogic := make([]*logicModels.OrderElement, 0, len(elemsDB))
	for _, elemOld := range elemsDB {
		elem := &logicModels.OrderElement{}
		err := copier.Copy(elem, elemOld)
		if err != nil {
			return nil, err
		}

		elemsLogic = append(elemsLogic, elem)
	}

	return elemsLogic, nil
}

func (el OrderElRepository) Add(ctx context.Context, elem *logicModels.OrderElement, order *logicModels.Order) error {
	err := el.db.Transaction(func(tx *gorm.DB) error {
		elemDB := &repoModels.OrderElement{}
		elem.Count += 1
		err := copier.Copy(elemDB, elem)
		if err != nil {
			return fmt.Errorf("add copy: %w", err)
		}

		res := tx.WithContext(ctx).Table("order_elements").Save(elemDB)
		if res.Error != nil {
			return fmt.Errorf("add save: %w", res.Error)
		}

		orderDB := &repoModels.Order{}
		err = copier.Copy(orderDB, order)
		if err != nil {
			return fmt.Errorf("copy order: %w", err)
		}
		res = tx.WithContext(ctx).Table("orders").Save(orderDB)
		if res.Error != nil {
			return fmt.Errorf("update: %w", res.Error)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("error transaction: %w", err)
	}
	return nil
}

func (el OrderElRepository) Decrease(ctx context.Context, elem *logicModels.OrderElement, order *logicModels.Order) error {
	err := el.db.Transaction(func(tx *gorm.DB) error {
		elem.Count -= 1
		elemDB := &repoModels.OrderElement{}
		err := copier.Copy(elemDB, elem)
		if err != nil {
			return fmt.Errorf("copy elem: %w", err)
		}
		if elemDB.Count == 0 {
			return fmt.Errorf("error of count, need delete")
		}
		res := tx.WithContext(ctx).Table("order_elements").Save(elemDB)
		if res.Error != nil {
			return fmt.Errorf("decrease save: %w", res.Error)
		}

		orderDB := &repoModels.Order{}
		err = copier.Copy(orderDB, order)
		if err != nil {
			return fmt.Errorf("error in copy order: %w", err)
		}

		res = tx.WithContext(ctx).Table("orders").Save(orderDB)
		if res.Error != nil {
			return fmt.Errorf("update: %w", res.Error)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("error transaction: %w", err)
	}
	return nil
}

func (el OrderElRepository) Delete(ctx context.Context, ID uuid.UUID, order *logicModels.Order) error {
	err := el.db.Transaction(func(tx *gorm.DB) error {

		res := tx.WithContext(ctx).Table("order_elements").Where("id = ?", ID).Delete(&repoModels.OrderElement{})
		if res.Error != nil {
			return fmt.Errorf("delete: %w", res.Error)
		}

		orderDB := &repoModels.Order{}
		err := copier.Copy(orderDB, order)
		if err != nil {
			return fmt.Errorf("error in copy order: %w", err)
		}

		res = tx.WithContext(ctx).Table("orders").Save(orderDB)
		if res.Error != nil {
			return fmt.Errorf("update: %w", res.Error)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("error transaction: %w", err)
	}
	return nil
}

func (el OrderElRepository) WineInOrder(ctx context.Context, IDWine uuid.UUID, IDOrder uuid.UUID) error {
	elem := &repoModels.OrderElement{}
	res := el.db.WithContext(ctx).Table("order_elements").Where("id_wine = ? and id_order = ?", IDWine, IDOrder).Take(elem)
	if res.Error != nil {
		return fmt.Errorf("error of get order: %w", res.Error)
	}
	return nil
}
