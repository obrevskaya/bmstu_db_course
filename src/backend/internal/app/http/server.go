package http

import (
	"context"
	"net/http"
	"strconv"

	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/interfaces"
	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/logic/models"
	openapi "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/server"
	"github.com/google/uuid"
)

const no = "false"
const yes = "true"

type Server struct {
	openapi.DefaultAPIService
	userLogic     interfaces.IUserController
	billLogic     interfaces.IBillController
	orderLogic    interfaces.IOrderController
	elemLogic     interfaces.IOrderElementController
	wineLogic     interfaces.IWineController
	userWineLogic interfaces.IUserWinesController
}

func NewServer(u interfaces.IUserController, b interfaces.IBillController,
	o interfaces.IOrderController, el interfaces.IOrderElementController,
	w interfaces.IWineController, uw interfaces.IUserWinesController) *Server {
	return &Server{
		userLogic:     u,
		billLogic:     b,
		orderLogic:    o,
		elemLogic:     el,
		wineLogic:     w,
		userWineLogic: uw,
	}
}

func (s *Server) Authorize(ctx context.Context, request openapi.AuthRequest) (openapi.ImplResponse, error) {
	user, err := s.userLogic.Authorize(ctx, request.Login, request.Password)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusUnauthorized,
			Body: openapi.ErrorResponse{
				Message:       "Authorization failed.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.User{
			Id:       user.ID.String(),
			Login:    user.Login,
			Password: user.Password,
			Fio:      user.Fio,
			Email:    user.Email,
			Points:   strconv.Itoa(user.Points),
			Status:   strconv.Itoa(user.Status),
		},
	}, nil
}

func (s *Server) AddElem(ctx context.Context, s2 string) (openapi.ImplResponse, error) {
	id, err := uuid.Parse(s2)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect uuid.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	err = s.elemLogic.Add(ctx, id)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "can't add.",
				SystemMessage: err.Error(),
			},
		}, nil
	}
	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.AddElemResponse{Added: true},
	}, nil

}

func (s *Server) AddWine(ctx context.Context, wine openapi.AddWineRequest) (openapi.ImplResponse, error) {
	count, err := strconv.Atoi(wine.Count)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect count." + wine.Count,
				SystemMessage: err.Error(),
			},
		}, nil
	}

	price, err := strconv.Atoi(wine.Price)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect price." + wine.Count,
				SystemMessage: err.Error(),
			},
		}, nil
	}

	wineModel := &models.Wine{
		Name:     wine.Name,
		Count:    count,
		Year:     int(wine.Year),
		Strength: int(wine.Strength),
		Price:    price,
		Type:     wine.Type,
	}

	err = s.wineLogic.Create(ctx, wineModel)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "can't create wine.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.AddWineResponse{Added: true},
	}, nil
}

func (s *Server) CreateElem(ctx context.Context, elem openapi.CreateElemRequest) (openapi.ImplResponse, error) {
	wineID, err := uuid.Parse(elem.IdWine)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect wine uuid.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	elemModel := &models.OrderElement{
		IDWine: wineID,
		Count:  int(elem.Count),
	}

	err = s.elemLogic.Create(ctx, elemModel)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "can't create elem.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.CreateElemResponse{Created: true},
	}, nil
}

func (s *Server) DecreaseElem(ctx context.Context, s2 string) (openapi.ImplResponse, error) {
	id, err := uuid.Parse(s2)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect uuid.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	err = s.elemLogic.Decrease(ctx, id)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "can't decrease.",
				SystemMessage: err.Error(),
			},
		}, nil
	}
	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.DecreaseElemResponse{Decreased: true},
	}, nil

}

func (s *Server) DeleteElem(ctx context.Context, s2 string) (openapi.ImplResponse, error) {
	id, err := uuid.Parse(s2)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect uuid.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	err = s.elemLogic.Delete(ctx, id)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "can't delete elem.",
				SystemMessage: err.Error(),
			},
		}, nil
	}
	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.DeleteElemResponse{Deleted: true},
	}, nil
}

func (s *Server) DeleteWine(ctx context.Context, s2 string) (openapi.ImplResponse, error) {
	id, err := uuid.Parse(s2)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect uuid.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	err = s.wineLogic.Delete(ctx, id)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "can't delete wine.",
				SystemMessage: err.Error(),
			},
		}, nil
	}
	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.DeleteWineResponse{Deleted: true},
	}, nil
}

func (s *Server) GetOrder(ctx context.Context, s2 string) (openapi.ImplResponse, error) {
	id, err := uuid.Parse(s2)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect uuid.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	order, err := s.orderLogic.GetByID(ctx, id)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "can't get order.",
				SystemMessage: err.Error(),
			},
		}, nil
	}
	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.Order{
			Id:         order.ID.String(),
			IdUser:     order.IDUser.String(),
			TotalPrice: strconv.Itoa(order.TotalPrice),
			IsPoints:   strconv.FormatBool(order.IsPoints),
			Status:     order.Status,
		},
	}, nil
}

func (s *Server) GetWines(ctx context.Context, request openapi.GetWinesRequest) (openapi.ImplResponse, error) {
	limit, err := strconv.Atoi(request.Limit)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect limit." + request.Limit,
				SystemMessage: err.Error(),
			},
		}, nil
	}

	skip, err := strconv.Atoi(request.Skip)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect skip." + request.Skip,
				SystemMessage: err.Error(),
			},
		}, nil
	}

	wines, err := s.wineLogic.GetWines(ctx, limit, skip)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "can't get wines.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	winesApi := make([]openapi.Wine, len(wines))

	for i, w := range wines {
		winesApi[i] = openapi.Wine{
			Id:       w.ID.String(),
			Name:     w.Name,
			Count:    strconv.Itoa(w.Count),
			Year:     int32(w.Year),
			Strength: int32(w.Strength),
			Price:    strconv.Itoa(w.Price),
			Type:     w.Type,
		}
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.GetWinesResponse{
			Wines: winesApi,
		},
	}, nil
}

func (s *Server) PayBill(ctx context.Context, s2 string) (openapi.ImplResponse, error) {
	id, err := uuid.Parse(s2)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect uuid.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	err = s.billLogic.UpdateBillStatus(ctx, id, models.PaidBill)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "can't update bill.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.PayBillResponse{Payed: true},
	}, nil
}

func (s *Server) PlaceOrder(ctx context.Context, order openapi.Order) (openapi.ImplResponse, error) {
	orderID, err := uuid.Parse(order.Id)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect order uuid.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	userID, err := uuid.Parse(order.IdUser)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect user uuid.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	totalPrice, err := strconv.Atoi(order.TotalPrice)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect total price." + order.TotalPrice,
				SystemMessage: err.Error(),
			},
		}, nil
	}

	var isPoints bool
	if order.IsPoints == yes {
		isPoints = true
	} else if order.IsPoints == no {
		isPoints = false
	} else {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect is points." + order.IsPoints,
				SystemMessage: err.Error(),
			},
		}, nil
	}

	orderModel := &models.Order{
		ID:         orderID,
		IDUser:     userID,
		TotalPrice: totalPrice,
		IsPoints:   isPoints,
		Status:     order.Status,
	}

	err = s.orderLogic.Update(ctx, orderModel)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "can't update order.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.PlaceOrderResponse{Placed: true},
	}, nil
}

func (s *Server) Register(ctx context.Context, user openapi.RegisterRequest) (openapi.ImplResponse, error) {
	status, err := strconv.Atoi(user.Status)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect status." + user.Status,
				SystemMessage: err.Error(),
			},
		}, nil
	}
	userModel := &models.User{
		Login:    user.Login,
		Password: user.Password,
		Fio:      user.Fio,
		Email:    user.Email,
		Points:   int(user.Points),
		Status:   status,
	}

	err = s.userLogic.Create(ctx, userModel)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusConflict,
			Body: openapi.ErrorResponse{
				Message:       "can't create user.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	return openapi.ImplResponse{
		Code: http.StatusCreated,
		Body: openapi.RegisterResponse{Registered: true},
	}, nil
}

func (s *Server) UpdatePoints(ctx context.Context, request openapi.UpdatePointsRequest) (openapi.ImplResponse, error) {
	userID, err := uuid.Parse(request.Id)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect uuid.",
				SystemMessage: err.Error(),
			},
		}, nil
	}
	err = s.userLogic.UpdateUserPoints(ctx, userID, int(request.Points))
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "can't update points.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.UpdateWineResponse{Updated: true},
	}, nil
}

func (s *Server) UpdateWine(ctx context.Context, wine openapi.Wine) (openapi.ImplResponse, error) {
	wineID, err := uuid.Parse(wine.Id)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect uuid.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	count, err := strconv.Atoi(wine.Count)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect count." + wine.Count,
				SystemMessage: err.Error(),
			},
		}, nil
	}

	price, err := strconv.Atoi(wine.Price)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect price." + wine.Price,
				SystemMessage: err.Error(),
			},
		}, nil
	}

	err = s.wineLogic.Update(ctx, &models.Wine{
		ID:       wineID,
		Name:     wine.Name,
		Count:    count,
		Year:     int(wine.Year),
		Strength: int(wine.Strength),
		Price:    price,
		Type:     wine.Type,
	})
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "can't update wine.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.UpdateWineResponse{Updated: true},
	}, nil
}

func (s *Server) GetWine(ctx context.Context, id string) (openapi.ImplResponse, error) {
	wineID, err := uuid.Parse(id)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect uuid.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	w, err := s.wineLogic.GetWine(ctx, wineID)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "can't get wine.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.Wine{
			Id:       w.ID.String(),
			Name:     w.Name,
			Count:    strconv.Itoa(w.Count),
			Year:     int32(w.Year),
			Strength: int32(w.Strength),
			Price:    strconv.Itoa(w.Price),
			Type:     w.Type,
		},
	}, nil
}

func (s *Server) GetByOrder(ctx context.Context, request openapi.GetByOrderRequest) (openapi.ImplResponse, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect uuid.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	elems, err := s.elemLogic.GetByOrder(ctx, id)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "can't get by order.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	elemsApi := make([]openapi.Elem, len(elems))

	for i, el := range elems {
		elemsApi[i] = openapi.Elem{
			Id:      el.ID.String(),
			IdOrder: el.IDOrder.String(),
			IdWine:  el.IDWine.String(),
			Count:   int32(el.Count),
		}
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.Elems{
			Elems: elemsApi,
		},
	}, nil
}

func (s *Server) CreateUserWine(ctx context.Context, request openapi.UserWine) (openapi.ImplResponse, error) {
	wineID, err := uuid.Parse(request.IdWine)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect wine uuid.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	userID, err := uuid.Parse(request.IdUser)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect user uuid.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	err = s.userWineLogic.Create(ctx, userID, wineID)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "can't create user wine.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.CreateUserWineResponse{Created: true},
	}, nil
}

func (s *Server) DeleteUserWine(ctx context.Context, request openapi.UserWine) (openapi.ImplResponse, error) {
	wineID, err := uuid.Parse(request.IdWine)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect wine uuid.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	userID, err := uuid.Parse(request.IdUser)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect user uuid.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	err = s.userWineLogic.DeleteWine(ctx, userID, wineID)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "can't delete user wine.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.DeleteUserWineResponse{Deleted: true},
	}, nil
}

func (s *Server) GetUserWines(ctx context.Context, id string) (openapi.ImplResponse, error) {
	userID, err := uuid.Parse(id)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "incorrect user uuid.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	wines, err := s.userWineLogic.GetByUser(ctx, userID)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "can't get user wines.",
				SystemMessage: err.Error(),
			},
		}, nil
	}

	userWinesApi := make([]openapi.UserWine, len(wines))

	for i, w := range wines {
		userWinesApi[i] = openapi.UserWine{
			IdUser: w.IDUser.String(),
			IdWine: w.IDWine.String(),
		}
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.UserWines{
			UserWines: userWinesApi,
		},
	}, nil
}
