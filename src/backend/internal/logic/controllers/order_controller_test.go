package controllers

import (
	"context"
	"reflect"
	"testing"

	myContext "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/context"
	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/interfaces"
	myErrors "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/logic/errors"
	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/logic/models"
	mocks2 "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/mocks"
	"github.com/gojuno/minimock/v3"
	"github.com/google/uuid"
)

func TestOrder_Create(t *testing.T) {
	mc := minimock.NewController(t)
	type fields struct {
		orderRep interfaces.IOrderRepository
		userRep  interfaces.IUserRepository
		billRep  interfaces.IBillRepository
	}
	type args struct {
		ctx   context.Context
		order *models.Order
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "successful create order",
			fields: fields{
				orderRep: mocks2.NewIOrderRepositoryMock(mc).InsertMock.Return(nil).
					GetByUserInProcessMock.Return(nil, myErrors.ErrNotFound),
				billRep: nil,
				userRep: nil,
			},
			args: args{
				ctx: myContext.UserToContext(context.Background(), &models.User{
					ID:       uuid.New(),
					Login:    "obrevskaya",
					Password: "qwerty1234",
					Fio:      "Obrevskaya Veronika",
					Email:    "obrevskaya.vera@mail.ru",
					Points:   0,
					Status:   models.Customer,
				}),
				order: &models.Order{
					IDUser:     uuid.New(),
					TotalPrice: 100,
					IsPoints:   false,
					Status:     models.ProcessOrder,
				},
			},
			wantErr: false,
		},
		{
			name: "order already exist",
			fields: fields{
				orderRep: mocks2.NewIOrderRepositoryMock(mc).InsertMock.Return(nil).
					GetByUserInProcessMock.Return(nil, nil),
				billRep: nil,
				userRep: nil,
			},
			args: args{
				ctx: myContext.UserToContext(context.Background(), &models.User{
					ID:       uuid.New(),
					Login:    "obrevskaya",
					Password: "qwerty1234",
					Fio:      "Obrevskaya Veronika",
					Email:    "obrevskaya.vera@mail.ru",
					Points:   0,
					Status:   models.Customer,
				}),
				order: &models.Order{
					IDUser:     uuid.New(),
					TotalPrice: 100,
					IsPoints:   false,
					Status:     models.ProcessOrder,
				},
			},
			wantErr: true,
		},
		{
			name: "error get order in db",
			fields: fields{
				orderRep: mocks2.NewIOrderRepositoryMock(mc).InsertMock.Return(nil).
					GetByUserInProcessMock.Return(nil, myErrors.ErrGet),
				billRep: nil,
				userRep: nil,
			},
			args: args{
				ctx: myContext.UserToContext(context.Background(), &models.User{
					ID:       uuid.New(),
					Login:    "obrevskaya",
					Password: "qwerty1234",
					Fio:      "Obrevskaya Veronika",
					Email:    "obrevskaya.vera@mail.ru",
					Points:   0,
					Status:   models.Customer,
				}),
				order: &models.Order{
					IDUser:     uuid.New(),
					TotalPrice: 100,
					IsPoints:   false,
					Status:     models.ProcessOrder,
				},
			},
			wantErr: true,
		},
		{
			name: "no context",
			fields: fields{
				orderRep: mocks2.NewIOrderRepositoryMock(mc).InsertMock.Return(nil).
					GetByUserInProcessMock.Return(nil, myErrors.ErrNotFound),
				billRep: nil,
				userRep: nil,
			},
			args: args{
				ctx: context.Background(),
				order: &models.Order{
					IDUser:     uuid.New(),
					TotalPrice: 100,
					IsPoints:   false,
					Status:     models.ProcessOrder,
				},
			},
			wantErr: true,
		},
		{
			name: "error create order",
			fields: fields{
				orderRep: mocks2.NewIOrderRepositoryMock(mc).InsertMock.Return(myErrors.ErrInsert).
					GetByUserInProcessMock.Return(nil, myErrors.ErrNotFound),
				billRep: nil,
				userRep: nil,
			},
			args: args{
				ctx: myContext.UserToContext(context.Background(), &models.User{
					ID:       uuid.New(),
					Login:    "obrevskaya",
					Password: "qwerty1234",
					Fio:      "Obrevskaya Veronika",
					Email:    "obrevskaya.vera@mail.ru",
					Points:   0,
					Status:   models.Customer,
				}),
				order: &models.Order{
					IDUser:     uuid.New(),
					TotalPrice: 100,
					IsPoints:   false,
					Status:     models.ProcessOrder,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OrderController{
				orderRep: tt.fields.orderRep,
				userRep:  tt.fields.userRep,
				billRep:  tt.fields.billRep,
			}
			if err := o.Create(tt.args.ctx, tt.args.order); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOrder_Delete(t *testing.T) {
	mc := minimock.NewController(t)
	type fields struct {
		orderRep interfaces.IOrderRepository
		userRep  interfaces.IUserRepository
		billRep  interfaces.IBillRepository
	}
	type args struct {
		ctx context.Context
		ID  uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "successful delete order",
			fields: fields{
				orderRep: mocks2.NewIOrderRepositoryMock(mc).DeleteMock.Return(nil),
				billRep:  nil,
				userRep:  nil,
			},
			args: args{
				ctx: myContext.UserToContext(context.Background(), &models.User{
					ID:       uuid.New(),
					Login:    "obrevskaya",
					Password: "qwerty1234",
					Fio:      "Obrevskaya Veronika",
					Email:    "obrevskaya.vera@mail.ru",
					Points:   0,
					Status:   models.Admin,
				}),
				ID: uuid.New(),
			},
			wantErr: false,
		},
		{
			name: "no context",
			fields: fields{
				orderRep: mocks2.NewIOrderRepositoryMock(mc).DeleteMock.Return(nil),
				billRep:  nil,
				userRep:  nil,
			},
			args: args{
				ctx: context.Background(),
				ID:  uuid.New(),
			},
			wantErr: true,
		},
		{
			name: "no access rights",
			fields: fields{
				orderRep: mocks2.NewIOrderRepositoryMock(mc).DeleteMock.Return(nil),
				billRep:  nil,
				userRep:  nil,
			},
			args: args{
				ctx: myContext.UserToContext(context.Background(), &models.User{
					ID:       uuid.New(),
					Login:    "obrevskaya",
					Password: "qwerty1234",
					Fio:      "Obrevskaya Veronika",
					Email:    "obrevskaya.vera@mail.ru",
					Points:   0,
					Status:   models.Customer,
				}),
				ID: uuid.New(),
			},
			wantErr: true,
		},
		{
			name: "error delete order",
			fields: fields{
				orderRep: mocks2.NewIOrderRepositoryMock(mc).DeleteMock.Return(myErrors.ErrDelete),
				billRep:  nil,
				userRep:  nil,
			},
			args: args{
				ctx: myContext.UserToContext(context.Background(), &models.User{
					ID:       uuid.New(),
					Login:    "obrevskaya",
					Password: "qwerty1234",
					Fio:      "Obrevskaya Veronika",
					Email:    "obrevskaya.vera@mail.ru",
					Points:   0,
					Status:   models.Admin,
				}),
				ID: uuid.New(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OrderController{
				orderRep: tt.fields.orderRep,
				userRep:  tt.fields.userRep,
				billRep:  tt.fields.billRep,
			}
			if err := o.Delete(tt.args.ctx, tt.args.ID); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOrder_Get(t *testing.T) {
	mc := minimock.NewController(t)

	o := &models.Order{
		ID:         uuid.New(),
		IDUser:     uuid.New(),
		TotalPrice: 100,
		IsPoints:   false,
		Status:     models.ProcessOrder,
	}
	type fields struct {
		orderRep interfaces.IOrderRepository
		userRep  interfaces.IUserRepository
		billRep  interfaces.IBillRepository
	}
	type args struct {
		ctx context.Context
		ID  uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Order
		wantErr bool
	}{
		{
			name: "successful get order",
			fields: fields{
				orderRep: mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, nil),
				billRep:  nil,
				userRep:  nil,
			},
			args: args{
				ctx: myContext.UserToContext(context.Background(), &models.User{
					ID:       uuid.New(),
					Login:    "obrevskaya",
					Password: "qwerty1234",
					Fio:      "Obrevskaya Veronika",
					Email:    "obrevskaya.vera@mail.ru",
					Points:   0,
					Status:   models.Admin,
				}),
				ID: uuid.New(),
			},
			want:    o,
			wantErr: false,
		},
		{
			name: "no context",
			fields: fields{
				orderRep: mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, nil),
				billRep:  nil,
				userRep:  nil,
			},
			args: args{
				ctx: context.Background(),
				ID:  uuid.New(),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "error get order",
			fields: fields{
				orderRep: mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, myErrors.ErrGetDB),
				billRep:  nil,
				userRep:  nil,
			},
			args: args{
				ctx: myContext.UserToContext(context.Background(), &models.User{
					ID:       uuid.New(),
					Login:    "obrevskaya",
					Password: "qwerty1234",
					Fio:      "Obrevskaya Veronika",
					Email:    "obrevskaya.vera@mail.ru",
					Points:   0,
					Status:   models.Admin,
				}),
				ID: uuid.New(),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OrderController{
				orderRep: tt.fields.orderRep,
				userRep:  tt.fields.userRep,
				billRep:  tt.fields.billRep,
			}
			got, err := o.GetByID(tt.args.ctx, tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderController_Update(t *testing.T) {
	mc := minimock.NewController(t)

	u := &models.User{
		ID:       uuid.New(),
		Login:    "wera",
		Password: "qwertyuiop",
		Fio:      "wera wera",
		Email:    "rty6@mail.ru",
		Points:   10000,
		Status:   models.Customer,
	}
	w := &models.Wine{
		ID:       uuid.New(),
		Name:     "q",
		Count:    100,
		Year:     2000,
		Strength: 6,
		Price:    10,
		Type:     "red",
	}
	elems := []*models.OrderElement{
		{ID: uuid.New(),
			IDOrder: uuid.New(),
			IDWine:  uuid.New(),
			Count:   1,
		}}
	orderProcess := &models.Order{
		ID:         uuid.New(),
		IDUser:     uuid.New(),
		TotalPrice: 100,
		IsPoints:   false,
		Status:     models.ProcessOrder,
	}
	orderPlaced := &models.Order{
		ID:         uuid.New(),
		IDUser:     uuid.New(),
		TotalPrice: 100,
		IsPoints:   false,
		Status:     models.PlacedOrder,
	}

	type fields struct {
		orderRep interfaces.IOrderRepository
		elemRep  interfaces.IOrderElementRepository
		userRep  interfaces.IUserRepository
		billRep  interfaces.IBillRepository
		wineRep  interfaces.IWineRepository
	}
	type args struct {
		ctx   context.Context
		order *models.Order
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "successful update",
			fields: fields{
				orderRep: mocks2.NewIOrderRepositoryMock(mc).UpdateMock.Return(nil).GetMock.Return(orderProcess, nil),
				elemRep:  mocks2.NewIOrderElementRepositoryMock(mc).GetByOrderMock.Return(elems, nil),
				userRep:  mocks2.NewIUserRepositoryMock(mc).GetMock.Return(u, nil),
				billRep:  mocks2.NewIBillRepositoryMock(mc).InsertMock.Return(nil),
				wineRep:  mocks2.NewIWineRepositoryMock(mc).ReduceWinesMock.Return(nil),
			},
			args: args{
				ctx: myContext.UserToContext(
					context.Background(),
					&models.User{
						ID:       uuid.New(),
						Login:    "wera",
						Password: "qwertyuiop",
						Fio:      "wera wera",
						Email:    "rty6@mail.ru",
						Points:   10000,
						Status:   models.Admin,
					}),
				order: &models.Order{
					ID:         uuid.New(),
					IDUser:     uuid.New(),
					IsPoints:   false,
					TotalPrice: 100,
					Status:     models.PlacedOrder,
				},
			},
			wantErr: false,
		},
		{
			name: "no context",
			fields: fields{
				orderRep: mocks2.NewIOrderRepositoryMock(mc).UpdateMock.Return(nil).InsertMock.Return(nil),
				elemRep:  mocks2.NewIOrderElementRepositoryMock(mc).GetByOrderMock.Return(elems, nil),
				userRep:  mocks2.NewIUserRepositoryMock(mc).GetMock.Return(u, nil),
				billRep:  mocks2.NewIBillRepositoryMock(mc).InsertMock.Return(nil),
				wineRep:  mocks2.NewIWineRepositoryMock(mc).GetWineMock.Return(w, nil).UpdateMock.Return(nil),
			},
			args: args{
				ctx: context.Background(),
				order: &models.Order{
					ID:         uuid.New(),
					IDUser:     uuid.New(),
					IsPoints:   false,
					TotalPrice: 100,
					Status:     models.PlacedOrder,
				},
			},
			wantErr: true,
		},
		{
			name: "no access rights",
			fields: fields{
				orderRep: mocks2.NewIOrderRepositoryMock(mc).UpdateMock.Return(nil).GetMock.Return(orderProcess, nil),
				elemRep:  mocks2.NewIOrderElementRepositoryMock(mc).GetByOrderMock.Return(elems, nil),
				userRep:  mocks2.NewIUserRepositoryMock(mc).GetMock.Return(u, nil),
				billRep:  mocks2.NewIBillRepositoryMock(mc).InsertMock.Return(nil),
				wineRep:  mocks2.NewIWineRepositoryMock(mc).GetWineMock.Return(w, nil).UpdateMock.Return(nil),
			},
			args: args{
				ctx: myContext.UserToContext(
					context.Background(),
					&models.User{
						ID:       uuid.New(),
						Login:    "wera",
						Password: "qwertyuiop",
						Fio:      "wera wera",
						Email:    "rty6@mail.ru",
						Points:   10000,
						Status:   models.Customer,
					}),
				order: &models.Order{
					ID:         uuid.New(),
					IDUser:     uuid.New(),
					IsPoints:   false,
					TotalPrice: 100,
					Status:     models.PlacedOrder,
				},
			},
			wantErr: true,
		},
		{
			name: "error get user",
			fields: fields{
				orderRep: mocks2.NewIOrderRepositoryMock(mc).UpdateMock.Return(nil).GetMock.Return(orderProcess, nil),
				elemRep:  mocks2.NewIOrderElementRepositoryMock(mc).GetByOrderMock.Return(elems, nil),
				userRep:  mocks2.NewIUserRepositoryMock(mc).GetMock.Return(u, myErrors.ErrGet),
				billRep:  mocks2.NewIBillRepositoryMock(mc).InsertMock.Return(nil),
				wineRep:  mocks2.NewIWineRepositoryMock(mc).GetWineMock.Return(w, nil).UpdateMock.Return(nil),
			},
			args: args{
				ctx: myContext.UserToContext(
					context.Background(),
					&models.User{
						ID:       uuid.New(),
						Login:    "wera",
						Password: "qwertyuiop",
						Fio:      "wera wera",
						Email:    "rty6@mail.ru",
						Points:   10000,
						Status:   models.Admin,
					}),
				order: &models.Order{
					ID:         uuid.New(),
					IDUser:     uuid.New(),
					IsPoints:   false,
					TotalPrice: 100,
					Status:     models.PlacedOrder,
				},
			},
			wantErr: true,
		},
		{
			name: "error balance",
			fields: fields{
				orderRep: mocks2.NewIOrderRepositoryMock(mc).UpdateMock.Return(nil).GetMock.Return(orderProcess, nil),
				elemRep:  mocks2.NewIOrderElementRepositoryMock(mc).GetByOrderMock.Return(elems, nil),
				userRep:  mocks2.NewIUserRepositoryMock(mc).GetMock.Return(u, nil),
				billRep:  mocks2.NewIBillRepositoryMock(mc).InsertMock.Return(nil),
				wineRep:  mocks2.NewIWineRepositoryMock(mc).GetWineMock.Return(w, nil).UpdateMock.Return(nil),
			},
			args: args{
				ctx: myContext.UserToContext(
					context.Background(),
					&models.User{
						ID:       uuid.New(),
						Login:    "wera",
						Password: "qwertyuiop",
						Fio:      "wera wera",
						Email:    "rty6@mail.ru",
						Points:   10000,
						Status:   models.Admin,
					}),
				order: &models.Order{
					ID:         uuid.New(),
					IDUser:     uuid.New(),
					IsPoints:   true,
					TotalPrice: 100000,
					Status:     models.PlacedOrder,
				},
			},
			wantErr: true,
		},
		{
			name: "error order update",
			fields: fields{
				orderRep: mocks2.NewIOrderRepositoryMock(mc).UpdateMock.Return(myErrors.ErrUpdate).GetMock.Return(orderProcess, nil),
				elemRep:  mocks2.NewIOrderElementRepositoryMock(mc).GetByOrderMock.Return(elems, nil),
				userRep:  mocks2.NewIUserRepositoryMock(mc).GetMock.Return(u, nil),
				billRep:  mocks2.NewIBillRepositoryMock(mc).InsertMock.Return(nil),
				wineRep:  mocks2.NewIWineRepositoryMock(mc).ReduceWinesMock.Return(nil),
			},
			args: args{
				ctx: myContext.UserToContext(
					context.Background(),
					&models.User{
						ID:       uuid.New(),
						Login:    "wera",
						Password: "qwertyuiop",
						Fio:      "wera wera",
						Email:    "rty6@mail.ru",
						Points:   10000,
						Status:   models.Admin,
					}),
				order: &models.Order{
					ID:         uuid.New(),
					IDUser:     uuid.New(),
					IsPoints:   false,
					TotalPrice: 100,
					Status:     models.PlacedOrder,
				},
			},
			wantErr: true,
		},
		{
			name: "error get order",
			fields: fields{
				orderRep: mocks2.NewIOrderRepositoryMock(mc).UpdateMock.Return(nil).GetMock.Return(orderProcess, myErrors.ErrGet),
				elemRep:  mocks2.NewIOrderElementRepositoryMock(mc).GetByOrderMock.Return(elems, nil),
				userRep:  mocks2.NewIUserRepositoryMock(mc).GetMock.Return(u, nil),
				billRep:  mocks2.NewIBillRepositoryMock(mc).InsertMock.Return(nil),
				wineRep:  mocks2.NewIWineRepositoryMock(mc).ReduceWinesMock.Return(nil),
			},
			args: args{
				ctx: myContext.UserToContext(
					context.Background(),
					&models.User{
						ID:       uuid.New(),
						Login:    "wera",
						Password: "qwertyuiop",
						Fio:      "wera wera",
						Email:    "rty6@mail.ru",
						Points:   10000,
						Status:   models.Admin,
					}),
				order: &models.Order{
					ID:         uuid.New(),
					IDUser:     uuid.New(),
					IsPoints:   false,
					TotalPrice: 100,
					Status:     models.PlacedOrder,
				},
			},
			wantErr: true,
		},
		{
			name: "order already placed",
			fields: fields{
				orderRep: mocks2.NewIOrderRepositoryMock(mc).UpdateMock.Return(nil).GetMock.Return(orderPlaced, nil),
				elemRep:  mocks2.NewIOrderElementRepositoryMock(mc).GetByOrderMock.Return(elems, nil),
				userRep:  mocks2.NewIUserRepositoryMock(mc).GetMock.Return(u, nil),
				billRep:  mocks2.NewIBillRepositoryMock(mc).InsertMock.Return(nil),
				wineRep:  mocks2.NewIWineRepositoryMock(mc).ReduceWinesMock.Return(nil),
			},
			args: args{
				ctx: myContext.UserToContext(
					context.Background(),
					&models.User{
						ID:       uuid.New(),
						Login:    "wera",
						Password: "qwertyuiop",
						Fio:      "wera wera",
						Email:    "rty6@mail.ru",
						Points:   10000,
						Status:   models.Admin,
					}),
				order: &models.Order{
					ID:         uuid.New(),
					IDUser:     uuid.New(),
					IsPoints:   false,
					TotalPrice: 100,
					Status:     models.PlacedOrder,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OrderController{
				orderRep: tt.fields.orderRep,
				elemRep:  tt.fields.elemRep,
				userRep:  tt.fields.userRep,
				billRep:  tt.fields.billRep,
				wineRep:  tt.fields.wineRep,
			}
			if err := o.Update(tt.args.ctx, tt.args.order); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOrderController_GetByUserInProcess(t *testing.T) {
	mc := minimock.NewController(t)

	o := &models.Order{
		ID:         uuid.New(),
		IDUser:     uuid.New(),
		TotalPrice: 100,
		IsPoints:   false,
		Status:     models.ProcessOrder,
	}

	type fields struct {
		orderRep interfaces.IOrderRepository
		elemRep  interfaces.IOrderElementRepository
		userRep  interfaces.IUserRepository
		billRep  interfaces.IBillRepository
		wineRep  interfaces.IWineRepository
	}
	type args struct {
		ctx context.Context
		ID  uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Order
		wantErr bool
	}{
		{
			name: "successful test",
			fields: fields{
				orderRep: mocks2.NewIOrderRepositoryMock(mc).GetByUserInProcessMock.Return(o, nil),
				elemRep:  nil,
				userRep:  nil,
				billRep:  nil,
				wineRep:  nil,
			},
			args: args{
				ctx: myContext.UserToContext(context.Background(), &models.User{
					ID:       uuid.New(),
					Login:    "obrevskaya",
					Password: "qwerty1234",
					Fio:      "Obrevskaya Veronika",
					Email:    "obrevskaya.vera@mail.ru",
					Points:   0,
					Status:   models.Admin,
				}),
				ID: uuid.New(),
			},
			want:    o,
			wantErr: false,
		},
		{
			name: "no context",
			fields: fields{
				orderRep: mocks2.NewIOrderRepositoryMock(mc).GetByUserInProcessMock.Return(o, nil),
				elemRep:  nil,
				userRep:  nil,
				billRep:  nil,
				wineRep:  nil,
			},
			args: args{
				ctx: context.Background(),
				ID:  uuid.New(),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "error get by user",
			fields: fields{
				orderRep: mocks2.NewIOrderRepositoryMock(mc).GetByUserInProcessMock.Return(o, myErrors.ErrGet),
				elemRep:  nil,
				userRep:  nil,
				billRep:  nil,
				wineRep:  nil,
			},
			args: args{
				ctx: myContext.UserToContext(context.Background(), &models.User{
					ID:       uuid.New(),
					Login:    "obrevskaya",
					Password: "qwerty1234",
					Fio:      "Obrevskaya Veronika",
					Email:    "obrevskaya.vera@mail.ru",
					Points:   0,
					Status:   models.Admin,
				}),
				ID: uuid.New(),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OrderController{
				orderRep: tt.fields.orderRep,
				elemRep:  tt.fields.elemRep,
				userRep:  tt.fields.userRep,
				billRep:  tt.fields.billRep,
				wineRep:  tt.fields.wineRep,
			}
			got, err := o.GetByUserInProcess(tt.args.ctx, tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByUserInProcess() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByUserInProcess() got = %v, want %v", got, tt.want)
			}
		})
	}
}
