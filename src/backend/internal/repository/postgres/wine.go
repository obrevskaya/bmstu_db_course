package postgres

import (
	"context"
	"fmt"

	myErrors "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/logic/errors"
	logicModels "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/logic/models"
	repoModels "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/repository/postgres/models"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type WineRepository struct {
	db *gorm.DB
}

func NewWR(db *gorm.DB) WineRepository {
	return WineRepository{db: db}
}

func (w WineRepository) Insert(ctx context.Context, wine *logicModels.Wine) error {
	wine.ID = uuid.New()
	wineDB := &repoModels.Wine{
		ID:       wine.ID,
		Name:     wine.Name,
		Count:    wine.Count,
		Year:     wine.Year,
		Strength: wine.Strength,
		Price:    wine.Price,
		Type:     wine.Type,
	}
	res := w.db.WithContext(ctx).Table("wines").Create(wineDB)
	if res.Error != nil {
		return fmt.Errorf("insert: %w", res.Error)
	}
	return nil
}

func (w WineRepository) GetWine(ctx context.Context, ID uuid.UUID) (*logicModels.Wine, error) {
	wine := repoModels.Wine{}

	res := w.db.WithContext(ctx).Table("wines").Where("id = ?", ID).Take(&wine)
	if res.Error != nil {
		return nil, fmt.Errorf("select: %w", res.Error)
	}

	resWine := logicModels.Wine{
		ID:       wine.ID,
		Name:     wine.Name,
		Count:    wine.Count,
		Year:     wine.Year,
		Strength: wine.Strength,
		Price:    wine.Price,
		Type:     wine.Type,
	}
	return &resWine, nil
}
func (w WineRepository) GetWines(ctx context.Context, limit int, skip int) ([]*logicModels.Wine, error) {
	var winesDB []*repoModels.Wine

	res := w.db.WithContext(ctx).Table("wines").Limit(limit).Offset(skip).Find(&winesDB)
	if res.Error != nil {
		return nil, fmt.Errorf("select: %w", res.Error)
	}

	winesLogic := make([]*logicModels.Wine, 0, len(winesDB))
	for _, wineOld := range winesDB {
		wine := &logicModels.Wine{}
		err := copier.Copy(wine, wineOld)
		if err != nil {
			return nil, err
		}

		winesLogic = append(winesLogic, wine)
	}

	return winesLogic, nil
}
func (w WineRepository) Delete(ctx context.Context, ID uuid.UUID) error {
	res := w.db.WithContext(ctx).Table("wines").Where("id = ?", ID).Delete(&repoModels.Wine{})
	if res.Error != nil {
		return fmt.Errorf("delete: %w", res.Error)
	}

	return nil
}

func (w WineRepository) Update(ctx context.Context, wine *logicModels.Wine) error {
	wineDB := &repoModels.Wine{
		ID:       wine.ID,
		Name:     wine.Name,
		Count:    wine.Count,
		Year:     wine.Year,
		Strength: wine.Strength,
		Price:    wine.Price,
		Type:     wine.Type,
	}

	res := w.db.WithContext(ctx).Table("wines").Save(wineDB)
	if res.Error != nil {
		return fmt.Errorf("update: %w", res.Error)
	}

	return nil
}

func (w WineRepository) ReduceWines(ctx context.Context, elems []*logicModels.OrderElement) error {
	err := w.db.Transaction(func(tx *gorm.DB) error {
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
		return nil
	})
	if err != nil {
		return fmt.Errorf("error transaction: %w", err)
	}
	return nil
}
