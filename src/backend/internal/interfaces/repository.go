package interfaces

import (
	"context"

	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/logic/models"
	"github.com/google/uuid"
)

type IUserRepository interface {
	Insert(ctx context.Context, user *models.User) error
	Authorize(ctx context.Context, login string, password string) (*models.User, error)
	Get(ctx context.Context, ID uuid.UUID) (*models.User, error)
	UpdateUserPoints(ctx context.Context, ID uuid.UUID, points int) error
}

type IWineRepository interface {
	Insert(ctx context.Context, wine *models.Wine) error
	GetWine(ctx context.Context, ID uuid.UUID) (*models.Wine, error)
	GetWines(ctx context.Context, limit int, skip int) ([]*models.Wine, error)
	Delete(ctx context.Context, ID uuid.UUID) error
	Update(ctx context.Context, wine *models.Wine) error
	ReduceWines(ctx context.Context, els []*models.OrderElement) error
}

type IOrderRepository interface {
	Insert(ctx context.Context, order *models.Order) error
	Get(ctx context.Context, ID uuid.UUID) (*models.Order, error)
	GetByUserInProcess(ctx context.Context, ID uuid.UUID) (*models.Order, error)
	Delete(ctx context.Context, ID uuid.UUID) error
	Update(ctx context.Context, order *models.Order) error
}

type IOrderElementRepository interface {
	Insert(ctx context.Context, elem *models.OrderElement, order *models.Order) error
	GetByID(ctx context.Context, ID uuid.UUID) (*models.OrderElement, error)
	GetByOrder(ctx context.Context, IDOrder uuid.UUID) ([]*models.OrderElement, error)
	Add(ctx context.Context, elem *models.OrderElement, order *models.Order) error
	Decrease(ctx context.Context, elem *models.OrderElement, order *models.Order) error
	Delete(ctx context.Context, ID uuid.UUID, order *models.Order) error
	WineInOrder(ctx context.Context, IDWine uuid.UUID, IDOrder uuid.UUID) error
}

type IBillRepository interface {
	Insert(ctx context.Context, bill *models.Bill) error
	Get(ctx context.Context, ID uuid.UUID) (*models.Bill, error)
	UpdateBillStatus(ctx context.Context, ID uuid.UUID, status string, points int) error
}

type IUserWinesRepository interface {
	Insert(ctx context.Context, IDUser uuid.UUID, IDWine uuid.UUID) error
	GetByUser(ctx context.Context, IDUser uuid.UUID) ([]*models.UserWine, error)
	DeleteWine(ctx context.Context, IDUser uuid.UUID, IDWine uuid.UUID) error
}
