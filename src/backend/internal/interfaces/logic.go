package interfaces

import (
	"context"
	"time"

	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/logic/models"
	"github.com/google/uuid"
)

type IUserController interface {
	Create(ctx context.Context, user *models.User) error
	Authorize(ctx context.Context, login string, password string) (*models.User, error)
	Get(ctx context.Context, ID uuid.UUID) (*models.User, error)
	UpdateUserPoints(ctx context.Context, ID uuid.UUID, points int) error
}

type IWineController interface {
	Create(ctx context.Context, wine *models.Wine) error
	GetWine(ctx context.Context, ID uuid.UUID) (*models.Wine, error)
	GetWines(ctx context.Context, limit int, skip int) ([]*models.Wine, error)
	Delete(ctx context.Context, ID uuid.UUID) error
	Update(ctx context.Context, wine *models.Wine) error
}

type IOrderController interface {
	Create(ctx context.Context, order *models.Order) error
	GetByID(ctx context.Context, ID uuid.UUID) (*models.Order, error)
	GetByUserInProcess(ctx context.Context, ID uuid.UUID) (*models.Order, error)
	Delete(ctx context.Context, ID uuid.UUID) error
	Update(ctx context.Context, order *models.Order) error
}

type IOrderElementController interface {
	Create(ctx context.Context, elem *models.OrderElement) error
	GetByID(ctx context.Context, ID uuid.UUID) (*models.OrderElement, error)
	GetByOrder(ctx context.Context, IDOrder uuid.UUID) ([]*models.OrderElement, error)
	Add(ctx context.Context, ID uuid.UUID) error
	Decrease(ctx context.Context, ID uuid.UUID) error
	Delete(ctx context.Context, ID uuid.UUID) error

	ResearchCreate(ctx context.Context, elem *models.OrderElement) (error, time.Duration)
	ResearchCreateTr(ctx context.Context, elem *models.OrderElement) (error, time.Duration)
}

type IBillController interface {
	Create(ctx context.Context, bill *models.Bill) error
	Get(ctx context.Context, ID uuid.UUID) (*models.Bill, error)
	UpdateBillStatus(ctx context.Context, ID uuid.UUID, status string) error
}

type IUserWinesController interface {
	Create(ctx context.Context, IDUser uuid.UUID, IDWine uuid.UUID) error
	GetByUser(ctx context.Context, IDUser uuid.UUID) ([]*models.UserWine, error)
	DeleteWine(ctx context.Context, IDUser uuid.UUID, IDWine uuid.UUID) error
}
