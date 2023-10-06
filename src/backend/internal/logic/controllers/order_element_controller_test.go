package controllers

import (
	"context"
	"reflect"
	"testing"

	myContext "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/context"
	interfaces2 "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/interfaces"
	myErrors "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/logic/errors"
	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/logic/models"
	mocks2 "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/mocks"
	"github.com/gojuno/minimock/v3"
	"github.com/google/uuid"
)

func TestOrderElem_GetByID(t *testing.T) {
	mc := minimock.NewController(t)
	el := &models.OrderElement{
		ID:      uuid.New(),
		IDOrder: uuid.New(),
		IDWine:  uuid.New(),
		Count:   1,
	}

	type fields struct {
		orderElemRep interfaces2.IOrderElementRepository
		orderRep     interfaces2.IOrderRepository
		wineRep      interfaces2.IWineRepository
		orderLogic   interfaces2.IOrderController
	}
	type args struct {
		ctx context.Context
		ID  uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.OrderElement
		wantErr bool
	}{
		{
			name: "successful get element",
			fields: fields{
				orderElemRep: mocks2.NewIOrderElementRepositoryMock(mc).GetByIDMock.Return(el, nil),
				wineRep:      nil,
				orderRep:     nil,
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
			want:    el,
			wantErr: false,
		},
		{
			name: "no context",
			fields: fields{
				orderElemRep: mocks2.NewIOrderElementRepositoryMock(mc).GetByIDMock.Return(el, nil),
				wineRep:      nil,
				orderRep:     nil,
			},
			args: args{
				ctx: context.Background(),
				ID:  uuid.New(),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "error get element",
			fields: fields{
				orderElemRep: mocks2.NewIOrderElementRepositoryMock(mc).GetByIDMock.Return(el, myErrors.ErrGetDB),
				wineRep:      nil,
				orderRep:     nil,
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
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			el := &OrderElemController{
				orderElemRep: tt.fields.orderElemRep,
				orderRep:     tt.fields.orderRep,
				wineRep:      tt.fields.wineRep,
			}
			got, err := el.GetByID(tt.args.ctx, tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderElem_GetByOrder(t *testing.T) {
	mc := minimock.NewController(t)
	el := []*models.OrderElement{{
		ID:      uuid.New(),
		IDOrder: uuid.New(),
		IDWine:  uuid.New(),
		Count:   1,
	}}

	type fields struct {
		orderElemRep interfaces2.IOrderElementRepository
		orderRep     interfaces2.IOrderRepository
		wineRep      interfaces2.IWineRepository
		orderLogic   interfaces2.IOrderController
	}
	type args struct {
		ctx     context.Context
		IDOrder uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*models.OrderElement
		wantErr bool
	}{
		{
			name: "successful get by order element",
			fields: fields{
				orderElemRep: mocks2.NewIOrderElementRepositoryMock(mc).GetByOrderMock.Return(el, nil),
				wineRep:      nil,
				orderRep:     nil,
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
				IDOrder: uuid.New(),
			},
			want:    el,
			wantErr: false,
		},
		{
			name: "no context",
			fields: fields{
				orderElemRep: mocks2.NewIOrderElementRepositoryMock(mc).GetByOrderMock.Return(el, nil),
				wineRep:      nil,
				orderRep:     nil,
			},
			args: args{
				ctx:     context.Background(),
				IDOrder: uuid.New(),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "error get elements order",
			fields: fields{
				orderElemRep: mocks2.NewIOrderElementRepositoryMock(mc).GetByOrderMock.Return(el, myErrors.ErrGetDB),
				wineRep:      nil,
				orderRep:     nil,
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
				IDOrder: uuid.New(),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			el := &OrderElemController{
				orderElemRep: tt.fields.orderElemRep,
				orderRep:     tt.fields.orderRep,
				wineRep:      tt.fields.wineRep,
			}
			got, err := el.GetByOrder(tt.args.ctx, tt.args.IDOrder)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByOrder() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderElemController_Add(t *testing.T) {
	mc := minimock.NewController(t)

	elem := &models.OrderElement{
		ID:      uuid.New(),
		IDOrder: uuid.New(),
		IDWine:  uuid.New(),
		Count:   1,
	}

	elem1 := &models.OrderElement{
		ID:      uuid.New(),
		IDOrder: uuid.New(),
		IDWine:  uuid.New(),
		Count:   100,
	}

	w := &models.Wine{
		ID:       uuid.New(),
		Name:     "q",
		Count:    10,
		Year:     1999,
		Strength: 3,
		Price:    100,
		Type:     "red",
	}

	o := &models.Order{
		ID:         uuid.New(),
		IDUser:     uuid.New(),
		TotalPrice: 100,
		IsPoints:   false,
		Status:     models.ProcessOrder,
	}

	type fields struct {
		orderElemRep interfaces2.IOrderElementRepository
		orderRep     interfaces2.IOrderRepository
		wineRep      interfaces2.IWineRepository
		orderLogic   interfaces2.IOrderController
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
			name: "successful ",
			fields: fields{
				orderElemRep: mocks2.NewIOrderElementRepositoryMock(mc).GetByIDMock.Return(elem, nil).AddMock.Return(nil),
				orderRep:     mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, nil),
				wineRep:      mocks2.NewIWineRepositoryMock(mc).GetWineMock.Return(w, nil),
				orderLogic:   mocks2.NewIOrderControllerMock(mc).UpdateMock.Return(nil),
			},
			args: args{
				ctx: myContext.UserToContext(context.Background(),
					&models.User{
						ID:       uuid.New(),
						Login:    "veronika",
						Password: "qwertyuiop[",
						Fio:      "obrevskaya",
						Email:    "q@gmail.com",
						Points:   10000,
						Status:   models.Admin,
					}),
				ID: uuid.New(),
			},
			wantErr: false,
		},
		{
			name: "no context",
			fields: fields{
				orderElemRep: mocks2.NewIOrderElementRepositoryMock(mc).GetByIDMock.Return(elem, nil).AddMock.Return(nil),
				orderRep:     mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, nil),
				wineRep:      mocks2.NewIWineRepositoryMock(mc).GetWineMock.Return(w, nil),
				orderLogic:   mocks2.NewIOrderControllerMock(mc).UpdateMock.Return(nil),
			},
			args: args{
				ctx: context.Background(),
				ID:  uuid.New(),
			},
			wantErr: true,
		},
		{
			name: "error get elem",
			fields: fields{
				orderElemRep: mocks2.NewIOrderElementRepositoryMock(mc).GetByIDMock.Return(elem, myErrors.ErrGet).AddMock.Return(nil),
				orderRep:     mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, nil),
				wineRep:      mocks2.NewIWineRepositoryMock(mc).GetWineMock.Return(w, nil),
				orderLogic:   mocks2.NewIOrderControllerMock(mc).UpdateMock.Return(nil),
			},
			args: args{
				ctx: myContext.UserToContext(context.Background(),
					&models.User{
						ID:       uuid.New(),
						Login:    "veronika",
						Password: "qwertyuiop[",
						Fio:      "obrevskaya",
						Email:    "q@gmail.com",
						Points:   10000,
						Status:   models.Admin,
					}),
				ID: uuid.New(),
			},
			wantErr: true,
		},
		{
			name: "error get wine",
			fields: fields{
				orderElemRep: mocks2.NewIOrderElementRepositoryMock(mc).GetByIDMock.Return(elem, nil).AddMock.Return(nil),
				orderRep:     mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, nil),
				wineRep:      mocks2.NewIWineRepositoryMock(mc).GetWineMock.Return(w, myErrors.ErrGet),
				orderLogic:   mocks2.NewIOrderControllerMock(mc).UpdateMock.Return(nil),
			},
			args: args{
				ctx: myContext.UserToContext(context.Background(),
					&models.User{
						ID:       uuid.New(),
						Login:    "veronika",
						Password: "qwertyuiop[",
						Fio:      "obrevskaya",
						Email:    "q@gmail.com",
						Points:   10000,
						Status:   models.Admin,
					}),
				ID: uuid.New(),
			},
			wantErr: true,
		},
		{
			name: "error cnt wine",
			fields: fields{
				orderElemRep: mocks2.NewIOrderElementRepositoryMock(mc).GetByIDMock.Return(elem1, nil).AddMock.Return(nil),
				orderRep:     mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, nil),
				wineRep:      mocks2.NewIWineRepositoryMock(mc).GetWineMock.Return(w, nil),
				orderLogic:   mocks2.NewIOrderControllerMock(mc).UpdateMock.Return(nil),
			},
			args: args{
				ctx: myContext.UserToContext(context.Background(),
					&models.User{
						ID:       uuid.New(),
						Login:    "veronika",
						Password: "qwertyuiop[",
						Fio:      "obrevskaya",
						Email:    "q@gmail.com",
						Points:   10000,
						Status:   models.Admin,
					}),
				ID: uuid.New(),
			},
			wantErr: true,
		},
		{
			name: "error get order",
			fields: fields{
				orderElemRep: mocks2.NewIOrderElementRepositoryMock(mc).GetByIDMock.Return(elem, nil).AddMock.Return(nil),
				orderRep:     mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, myErrors.ErrGet),
				wineRep:      mocks2.NewIWineRepositoryMock(mc).GetWineMock.Return(w, nil),
				orderLogic:   mocks2.NewIOrderControllerMock(mc).UpdateMock.Return(nil),
			},
			args: args{
				ctx: myContext.UserToContext(context.Background(),
					&models.User{
						ID:       uuid.New(),
						Login:    "veronika",
						Password: "qwertyuiop[",
						Fio:      "obrevskaya",
						Email:    "q@gmail.com",
						Points:   10000,
						Status:   models.Admin,
					}),
				ID: uuid.New(),
			},
			wantErr: true,
		},
		{
			name: "error add elem",
			fields: fields{
				orderElemRep: mocks2.NewIOrderElementRepositoryMock(mc).GetByIDMock.Return(elem, nil).AddMock.Return(myErrors.ErrAdd),
				orderRep:     mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, nil),
				wineRep:      mocks2.NewIWineRepositoryMock(mc).GetWineMock.Return(w, nil),
				orderLogic:   mocks2.NewIOrderControllerMock(mc).UpdateMock.Return(nil),
			},
			args: args{
				ctx: myContext.UserToContext(context.Background(),
					&models.User{
						ID:       uuid.New(),
						Login:    "veronika",
						Password: "qwertyuiop[",
						Fio:      "obrevskaya",
						Email:    "q@gmail.com",
						Points:   10000,
						Status:   models.Admin,
					}),
				ID: uuid.New(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			el := &OrderElemController{
				orderElemRep: tt.fields.orderElemRep,
				orderRep:     tt.fields.orderRep,
				wineRep:      tt.fields.wineRep,
				orderLogic:   tt.fields.orderLogic,
			}
			if err := el.Add(tt.args.ctx, tt.args.ID); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOrderElemController_Create(t *testing.T) {
	mc := minimock.NewController(t)
	elem := &models.OrderElement{
		ID:      uuid.New(),
		IDOrder: uuid.New(),
		IDWine:  uuid.New(),
		Count:   1,
	}
	elemNew := &models.OrderElement{
		ID:     uuid.New(),
		IDWine: uuid.New(),
		Count:  1,
	}
	w := &models.Wine{
		ID:       uuid.New(),
		Name:     "q",
		Count:    10,
		Year:     1999,
		Strength: 3,
		Price:    100,
		Type:     "red",
	}
	w1 := &models.Wine{
		ID:       uuid.New(),
		Name:     "q",
		Count:    0,
		Year:     1999,
		Strength: 3,
		Price:    100,
		Type:     "red",
	}
	o := &models.Order{
		ID:         uuid.New(),
		IDUser:     uuid.New(),
		TotalPrice: 100,
		IsPoints:   false,
		Status:     models.ProcessOrder,
	}
	type fields struct {
		orderElemRep interfaces2.IOrderElementRepository
		orderRep     interfaces2.IOrderRepository
		wineRep      interfaces2.IWineRepository
		orderLogic   interfaces2.IOrderController
	}
	type args struct {
		ctx  context.Context
		elem *models.OrderElement
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "successful create elem",
			fields: fields{
				orderElemRep: mocks2.NewIOrderElementRepositoryMock(mc).InsertMock.Return(nil).DeleteMock.Return(nil).WineInOrderMock.Return(myErrors.ErrGet),
				orderRep:     mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, nil).InsertMock.Return(nil),
				wineRep:      mocks2.NewIWineRepositoryMock(mc).GetWineMock.Return(w, nil),
				orderLogic:   mocks2.NewIOrderControllerMock(mc).UpdateMock.Return(nil).GetByUserInProcessMock.Return(o, nil),
			},
			args: args{
				ctx: myContext.UserToContext(context.Background(), &models.User{
					ID:       uuid.New(),
					Login:    "veronika",
					Password: "qwertyuiop",
					Fio:      "obrevskaya",
					Points:   10000,
					Email:    "1@mail.ru",
					Status:   models.Customer,
				}),
				elem: elem,
			},
			wantErr: false,
		},
		{
			name: "successful create elem with create order",
			fields: fields{
				orderElemRep: mocks2.NewIOrderElementRepositoryMock(mc).InsertMock.Return(nil).DeleteMock.Return(nil).WineInOrderMock.Return(myErrors.ErrGet),
				orderRep:     mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, nil).InsertMock.Return(nil),
				wineRep:      mocks2.NewIWineRepositoryMock(mc).GetWineMock.Return(w, nil),
				orderLogic:   mocks2.NewIOrderControllerMock(mc).UpdateMock.Return(nil).GetByUserInProcessMock.Return(o, nil),
			},
			args: args{
				ctx: myContext.UserToContext(context.Background(), &models.User{
					ID:       uuid.New(),
					Login:    "veronika",
					Password: "qwertyuiop",
					Fio:      "obrevskaya",
					Points:   10000,
					Email:    "1@mail.ru",
					Status:   models.Customer,
				}),
				elem: elemNew,
			},
			wantErr: false,
		},
		{
			name: "no context",
			fields: fields{
				orderElemRep: mocks2.NewIOrderElementRepositoryMock(mc),
				orderRep:     mocks2.NewIOrderRepositoryMock(mc),
				wineRep:      mocks2.NewIWineRepositoryMock(mc),
				orderLogic:   mocks2.NewIOrderControllerMock(mc),
			},
			args: args{
				ctx:  context.Background(),
				elem: elem,
			},
			wantErr: true,
		},
		{
			name: "error get wine",
			fields: fields{
				orderElemRep: mocks2.NewIOrderElementRepositoryMock(mc).InsertMock.Return(nil).DeleteMock.Return(nil).WineInOrderMock.Return(myErrors.ErrGet),
				orderRep:     mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, nil).InsertMock.Return(nil),
				wineRep:      mocks2.NewIWineRepositoryMock(mc).GetWineMock.Return(w, myErrors.ErrGet),
				orderLogic:   mocks2.NewIOrderControllerMock(mc).UpdateMock.Return(nil).GetByUserInProcessMock.Return(o, nil),
			},
			args: args{
				ctx: myContext.UserToContext(context.Background(), &models.User{
					ID:       uuid.New(),
					Login:    "veronika",
					Password: "qwertyuiop",
					Fio:      "obrevskaya",
					Points:   10000,
					Email:    "1@mail.ru",
					Status:   models.Customer,
				}),
				elem: elem,
			},
			wantErr: true,
		},
		{
			name: "error cnt wine",
			fields: fields{
				orderElemRep: mocks2.NewIOrderElementRepositoryMock(mc).InsertMock.Return(nil).DeleteMock.Return(nil).WineInOrderMock.Return(myErrors.ErrGet),
				orderRep:     mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, nil).InsertMock.Return(nil),
				wineRep:      mocks2.NewIWineRepositoryMock(mc).GetWineMock.Return(w1, nil),
				orderLogic:   mocks2.NewIOrderControllerMock(mc).UpdateMock.Return(nil).GetByUserInProcessMock.Return(o, nil),
			},
			args: args{
				ctx: myContext.UserToContext(context.Background(), &models.User{
					ID:       uuid.New(),
					Login:    "veronika",
					Password: "qwertyuiop",
					Fio:      "obrevskaya",
					Points:   10000,
					Email:    "1@mail.ru",
					Status:   models.Customer,
				}),
				elem: elem,
			},
			wantErr: true,
		},
		{
			name: "error create order",
			fields: fields{
				orderElemRep: mocks2.NewIOrderElementRepositoryMock(mc).InsertMock.Return(nil).DeleteMock.Return(nil).WineInOrderMock.Return(myErrors.ErrGet),
				orderRep:     mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, nil).InsertMock.Return(myErrors.ErrInsert),
				wineRep:      mocks2.NewIWineRepositoryMock(mc).GetWineMock.Return(w, nil),
				orderLogic:   mocks2.NewIOrderControllerMock(mc).UpdateMock.Return(nil).GetByUserInProcessMock.Return(nil, myErrors.ErrGet),
			},
			args: args{
				ctx: myContext.UserToContext(context.Background(), &models.User{
					ID:       uuid.New(),
					Login:    "veronika",
					Password: "qwertyuiop",
					Fio:      "obrevskaya",
					Points:   10000,
					Email:    "1@mail.ru",
					Status:   models.Customer,
				}),
				elem: elemNew,
			},
			wantErr: true,
		},
		{
			name: "error insert elem",
			fields: fields{
				orderElemRep: mocks2.NewIOrderElementRepositoryMock(mc).InsertMock.Return(myErrors.ErrInsert).DeleteMock.Return(nil).WineInOrderMock.Return(myErrors.ErrGet),
				orderRep:     mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, nil).InsertMock.Return(nil),
				wineRep:      mocks2.NewIWineRepositoryMock(mc).GetWineMock.Return(w, nil),
				orderLogic:   mocks2.NewIOrderControllerMock(mc).UpdateMock.Return(nil).GetByUserInProcessMock.Return(o, nil),
			},
			args: args{
				ctx: myContext.UserToContext(context.Background(), &models.User{
					ID:       uuid.New(),
					Login:    "veronika",
					Password: "qwertyuiop",
					Fio:      "obrevskaya",
					Points:   10000,
					Email:    "1@mail.ru",
					Status:   models.Customer,
				}),
				elem: elem,
			},
			wantErr: true,
		},
		{
			name: "wine also in order",
			fields: fields{
				orderElemRep: mocks2.NewIOrderElementRepositoryMock(mc).InsertMock.Return(nil).DeleteMock.Return(nil).WineInOrderMock.Return(nil),
				orderRep:     mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, nil).InsertMock.Return(nil),
				wineRep:      mocks2.NewIWineRepositoryMock(mc).GetWineMock.Return(w, nil),
				orderLogic:   mocks2.NewIOrderControllerMock(mc).UpdateMock.Return(myErrors.ErrUpdate).GetByUserInProcessMock.Return(o, nil),
			},
			args: args{
				ctx: myContext.UserToContext(context.Background(), &models.User{
					ID:       uuid.New(),
					Login:    "veronika",
					Password: "qwertyuiop",
					Fio:      "obrevskaya",
					Points:   10000,
					Email:    "1@mail.ru",
					Status:   models.Customer,
				}),
				elem: elem,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			el := &OrderElemController{
				orderElemRep: tt.fields.orderElemRep,
				orderRep:     tt.fields.orderRep,
				wineRep:      tt.fields.wineRep,
				orderLogic:   tt.fields.orderLogic,
			}
			if err := el.Create(tt.args.ctx, tt.args.elem); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOrderElemController_Decrease(t *testing.T) {
	mc := minimock.NewController(t)
	o := &models.Order{
		ID:         uuid.New(),
		IDUser:     uuid.New(),
		TotalPrice: 1000,
		IsPoints:   false,
		Status:     models.ProcessOrder,
	}
	w := &models.Wine{
		ID:       uuid.New(),
		Name:     "q",
		Count:    10,
		Year:     1999,
		Strength: 3,
		Price:    100,
		Type:     "red",
	}
	w1 := &models.Wine{
		ID:       uuid.New(),
		Name:     "q",
		Count:    10,
		Year:     1999,
		Strength: 3,
		Price:    10011,
		Type:     "red",
	}
	elem := &models.OrderElement{
		ID:      uuid.New(),
		IDOrder: uuid.New(),
		IDWine:  uuid.New(),
		Count:   1,
	}

	elem2 := &models.OrderElement{
		ID:      uuid.New(),
		IDOrder: uuid.New(),
		IDWine:  uuid.New(),
		Count:   2,
	}
	type fields struct {
		orderElemRep interfaces2.IOrderElementRepository
		orderRep     interfaces2.IOrderRepository
		wineRep      interfaces2.IWineRepository
		orderLogic   interfaces2.IOrderController
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
			name: "successful decrease",
			fields: fields{
				orderElemRep: mocks2.NewIOrderElementRepositoryMock(mc).GetByIDMock.Return(elem, nil).
					DeleteMock.Return(nil).DecreaseMock.Return(nil).
					InsertMock.Return(nil).AddMock.Return(nil),
				orderRep:   mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, nil),
				wineRep:    mocks2.NewIWineRepositoryMock(mc).GetWineMock.Return(w, nil),
				orderLogic: mocks2.NewIOrderControllerMock(mc).UpdateMock.Return(nil),
			},
			args: args{
				ctx: myContext.UserToContext(context.Background(), &models.User{
					ID:       uuid.New(),
					Login:    "veronika",
					Password: "qwertyuiop",
					Fio:      "obrevskaya",
					Points:   10000,
					Email:    "1@mail.ru",
					Status:   models.Customer,
				}),
				ID: uuid.New(),
			},
			wantErr: false,
		},
		{
			name: "no context",
			fields: fields{
				orderElemRep: mocks2.NewIOrderElementRepositoryMock(mc).GetByIDMock.Return(elem, nil).
					DeleteMock.Return(nil).DecreaseMock.Return(nil).
					InsertMock.Return(nil).AddMock.Return(nil),
				orderRep:   mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, nil),
				wineRep:    mocks2.NewIWineRepositoryMock(mc).GetWineMock.Return(w, nil),
				orderLogic: mocks2.NewIOrderControllerMock(mc).UpdateMock.Return(nil),
			},
			args: args{
				ctx: context.Background(),
				ID:  uuid.New(),
			},
			wantErr: true,
		},
		{
			name: "error get elem",
			fields: fields{
				orderElemRep: mocks2.NewIOrderElementRepositoryMock(mc).GetByIDMock.Return(elem, myErrors.ErrGet).
					DeleteMock.Return(nil).DecreaseMock.Return(nil).
					InsertMock.Return(nil).AddMock.Return(nil),
				orderRep:   mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, nil),
				wineRep:    mocks2.NewIWineRepositoryMock(mc).GetWineMock.Return(w, nil),
				orderLogic: mocks2.NewIOrderControllerMock(mc).UpdateMock.Return(nil),
			},
			args: args{
				ctx: myContext.UserToContext(context.Background(), &models.User{
					ID:       uuid.New(),
					Login:    "veronika",
					Password: "qwertyuiop",
					Fio:      "obrevskaya",
					Points:   10000,
					Email:    "1@mail.ru",
					Status:   models.Customer,
				}),
				ID: uuid.New(),
			},
			wantErr: true,
		},
		{
			name: "error get wine",
			fields: fields{
				orderElemRep: mocks2.NewIOrderElementRepositoryMock(mc).GetByIDMock.Return(elem, nil).
					DeleteMock.Return(nil).DecreaseMock.Return(nil).
					InsertMock.Return(nil).AddMock.Return(nil),
				orderRep:   mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, nil),
				wineRep:    mocks2.NewIWineRepositoryMock(mc).GetWineMock.Return(w, myErrors.ErrGet),
				orderLogic: mocks2.NewIOrderControllerMock(mc).UpdateMock.Return(nil),
			},
			args: args{
				ctx: myContext.UserToContext(context.Background(), &models.User{
					ID:       uuid.New(),
					Login:    "veronika",
					Password: "qwertyuiop",
					Fio:      "obrevskaya",
					Points:   10000,
					Email:    "1@mail.ru",
					Status:   models.Customer,
				}),
				ID: uuid.New(),
			},
			wantErr: true,
		},
		{
			name: "error get order",
			fields: fields{
				orderElemRep: mocks2.NewIOrderElementRepositoryMock(mc).GetByIDMock.Return(elem, nil).
					DeleteMock.Return(nil).DecreaseMock.Return(nil).
					InsertMock.Return(nil).AddMock.Return(nil),
				orderRep:   mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, myErrors.ErrGet),
				wineRep:    mocks2.NewIWineRepositoryMock(mc).GetWineMock.Return(w, nil),
				orderLogic: mocks2.NewIOrderControllerMock(mc).UpdateMock.Return(nil),
			},
			args: args{
				ctx: myContext.UserToContext(context.Background(), &models.User{
					ID:       uuid.New(),
					Login:    "veronika",
					Password: "qwertyuiop",
					Fio:      "obrevskaya",
					Points:   10000,
					Email:    "1@mail.ru",
					Status:   models.Customer,
				}),
				ID: uuid.New(),
			},
			wantErr: true,
		},
		{
			name: "error price",
			fields: fields{
				orderElemRep: mocks2.NewIOrderElementRepositoryMock(mc).GetByIDMock.Return(elem, nil).
					DeleteMock.Return(nil).DecreaseMock.Return(nil).
					InsertMock.Return(nil).AddMock.Return(nil),
				orderRep:   mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, nil),
				wineRep:    mocks2.NewIWineRepositoryMock(mc).GetWineMock.Return(w1, nil),
				orderLogic: mocks2.NewIOrderControllerMock(mc).UpdateMock.Return(nil),
			},
			args: args{
				ctx: myContext.UserToContext(context.Background(), &models.User{
					ID:       uuid.New(),
					Login:    "veronika",
					Password: "qwertyuiop",
					Fio:      "obrevskaya",
					Points:   10000,
					Email:    "1@mail.ru",
					Status:   models.Customer,
				}),
				ID: uuid.New(),
			},
			wantErr: true,
		},
		{
			name: "error decrease",
			fields: fields{
				orderElemRep: mocks2.NewIOrderElementRepositoryMock(mc).GetByIDMock.Return(elem2, nil).
					DeleteMock.Return(nil).DecreaseMock.Return(myErrors.ErrDelete).
					InsertMock.Return(nil).AddMock.Return(nil),
				orderRep:   mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, nil),
				wineRep:    mocks2.NewIWineRepositoryMock(mc).GetWineMock.Return(w, nil),
				orderLogic: mocks2.NewIOrderControllerMock(mc).UpdateMock.Return(nil),
			},
			args: args{
				ctx: myContext.UserToContext(context.Background(), &models.User{
					ID:       uuid.New(),
					Login:    "veronika",
					Password: "qwertyuiop",
					Fio:      "obrevskaya",
					Points:   10000,
					Email:    "1@mail.ru",
					Status:   models.Customer,
				}),
				ID: uuid.New(),
			},
			wantErr: true,
		},
		{
			name: "error delete",
			fields: fields{
				orderElemRep: mocks2.NewIOrderElementRepositoryMock(mc).GetByIDMock.Return(elem, nil).
					DeleteMock.Return(myErrors.ErrDelete).DecreaseMock.Return(nil).
					InsertMock.Return(nil).AddMock.Return(nil),
				orderRep:   mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, nil),
				wineRep:    mocks2.NewIWineRepositoryMock(mc).GetWineMock.Return(w, nil),
				orderLogic: mocks2.NewIOrderControllerMock(mc).UpdateMock.Return(nil),
			},
			args: args{
				ctx: myContext.UserToContext(context.Background(), &models.User{
					ID:       uuid.New(),
					Login:    "veronika",
					Password: "qwertyuiop",
					Fio:      "obrevskaya",
					Points:   10000,
					Email:    "1@mail.ru",
					Status:   models.Customer,
				}),
				ID: uuid.New(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			el := &OrderElemController{
				orderElemRep: tt.fields.orderElemRep,
				orderRep:     tt.fields.orderRep,
				wineRep:      tt.fields.wineRep,
				orderLogic:   tt.fields.orderLogic,
			}
			if err := el.Decrease(tt.args.ctx, tt.args.ID); (err != nil) != tt.wantErr {
				t.Errorf("Decrease() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOrderElemController_Delete(t *testing.T) {
	mc := minimock.NewController(t)
	o := &models.Order{
		ID:         uuid.New(),
		IDUser:     uuid.New(),
		TotalPrice: 1000,
		IsPoints:   false,
		Status:     models.ProcessOrder,
	}
	w := &models.Wine{
		ID:       uuid.New(),
		Name:     "q",
		Count:    10,
		Year:     1999,
		Strength: 3,
		Price:    100,
		Type:     "red",
	}
	elem := &models.OrderElement{
		ID:      uuid.New(),
		IDOrder: uuid.New(),
		IDWine:  uuid.New(),
		Count:   1,
	}

	elem1 := &models.OrderElement{
		ID:      uuid.New(),
		IDOrder: uuid.New(),
		IDWine:  uuid.New(),
		Count:   20,
	}
	type fields struct {
		orderElemRep interfaces2.IOrderElementRepository
		orderRep     interfaces2.IOrderRepository
		wineRep      interfaces2.IWineRepository
		orderLogic   interfaces2.IOrderController
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
			name: "successful delete",
			fields: fields{
				orderElemRep: mocks2.NewIOrderElementRepositoryMock(mc).GetByIDMock.Return(elem, nil).
					DeleteMock.Return(nil).InsertMock.Return(nil),
				orderRep:   mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, nil),
				wineRep:    mocks2.NewIWineRepositoryMock(mc).GetWineMock.Return(w, nil),
				orderLogic: mocks2.NewIOrderControllerMock(mc).UpdateMock.Return(nil),
			},
			args: args{
				ctx: myContext.UserToContext(context.Background(), &models.User{
					ID:       uuid.New(),
					Login:    "veronika",
					Password: "qwertyuiop",
					Fio:      "obrevskaya",
					Points:   10000,
					Email:    "1@mail.ru",
					Status:   models.Customer,
				}),
				ID: uuid.New(),
			},
			wantErr: false,
		},
		{
			name: "no context",
			fields: fields{
				orderElemRep: mocks2.NewIOrderElementRepositoryMock(mc).GetByIDMock.Return(elem, nil).
					DeleteMock.Return(nil).InsertMock.Return(nil),
				orderRep:   mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, nil),
				wineRep:    mocks2.NewIWineRepositoryMock(mc).GetWineMock.Return(w, nil),
				orderLogic: mocks2.NewIOrderControllerMock(mc).UpdateMock.Return(nil),
			},
			args: args{
				ctx: context.Background(),
				ID:  uuid.New(),
			},
			wantErr: true,
		},
		{
			name: "error get elem",
			fields: fields{
				orderElemRep: mocks2.NewIOrderElementRepositoryMock(mc).GetByIDMock.Return(elem, myErrors.ErrGet).
					DeleteMock.Return(nil).InsertMock.Return(nil),
				orderRep:   mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, nil),
				wineRep:    mocks2.NewIWineRepositoryMock(mc).GetWineMock.Return(w, nil),
				orderLogic: mocks2.NewIOrderControllerMock(mc).UpdateMock.Return(nil),
			},
			args: args{
				ctx: myContext.UserToContext(context.Background(), &models.User{
					ID:       uuid.New(),
					Login:    "veronika",
					Password: "qwertyuiop",
					Fio:      "obrevskaya",
					Points:   10000,
					Email:    "1@mail.ru",
					Status:   models.Customer,
				}),
				ID: uuid.New(),
			},
			wantErr: true,
		},
		{
			name: "error get elem",
			fields: fields{
				orderElemRep: mocks2.NewIOrderElementRepositoryMock(mc).GetByIDMock.Return(elem, myErrors.ErrGet).
					DeleteMock.Return(nil).InsertMock.Return(nil),
				orderRep:   mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, nil),
				wineRep:    mocks2.NewIWineRepositoryMock(mc).GetWineMock.Return(w, nil),
				orderLogic: mocks2.NewIOrderControllerMock(mc).UpdateMock.Return(nil),
			},
			args: args{
				ctx: myContext.UserToContext(context.Background(), &models.User{
					ID:       uuid.New(),
					Login:    "veronika",
					Password: "qwertyuiop",
					Fio:      "obrevskaya",
					Points:   10000,
					Email:    "1@mail.ru",
					Status:   models.Customer,
				}),
				ID: uuid.New(),
			},
			wantErr: true,
		},
		{
			name: "error get wine",
			fields: fields{
				orderElemRep: mocks2.NewIOrderElementRepositoryMock(mc).GetByIDMock.Return(elem, nil).
					DeleteMock.Return(nil).InsertMock.Return(nil),
				orderRep:   mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, nil),
				wineRep:    mocks2.NewIWineRepositoryMock(mc).GetWineMock.Return(w, myErrors.ErrGet),
				orderLogic: mocks2.NewIOrderControllerMock(mc).UpdateMock.Return(nil),
			},
			args: args{
				ctx: myContext.UserToContext(context.Background(), &models.User{
					ID:       uuid.New(),
					Login:    "veronika",
					Password: "qwertyuiop",
					Fio:      "obrevskaya",
					Points:   10000,
					Email:    "1@mail.ru",
					Status:   models.Customer,
				}),
				ID: uuid.New(),
			},
			wantErr: true,
		},
		{
			name: "error get order",
			fields: fields{
				orderElemRep: mocks2.NewIOrderElementRepositoryMock(mc).GetByIDMock.Return(elem, nil).
					DeleteMock.Return(nil).InsertMock.Return(nil),
				orderRep:   mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, myErrors.ErrGet),
				wineRep:    mocks2.NewIWineRepositoryMock(mc).GetWineMock.Return(w, nil),
				orderLogic: mocks2.NewIOrderControllerMock(mc).UpdateMock.Return(nil),
			},
			args: args{
				ctx: myContext.UserToContext(context.Background(), &models.User{
					ID:       uuid.New(),
					Login:    "veronika",
					Password: "qwertyuiop",
					Fio:      "obrevskaya",
					Points:   10000,
					Email:    "1@mail.ru",
					Status:   models.Customer,
				}),
				ID: uuid.New(),
			},
			wantErr: true,
		},
		{
			name: "error price",
			fields: fields{
				orderElemRep: mocks2.NewIOrderElementRepositoryMock(mc).GetByIDMock.Return(elem1, nil).
					DeleteMock.Return(nil).InsertMock.Return(nil),
				orderRep:   mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, nil),
				wineRep:    mocks2.NewIWineRepositoryMock(mc).GetWineMock.Return(w, nil),
				orderLogic: mocks2.NewIOrderControllerMock(mc).UpdateMock.Return(nil),
			},
			args: args{
				ctx: myContext.UserToContext(context.Background(), &models.User{
					ID:       uuid.New(),
					Login:    "veronika",
					Password: "qwertyuiop",
					Fio:      "obrevskaya",
					Points:   10000,
					Email:    "1@mail.ru",
					Status:   models.Customer,
				}),
				ID: uuid.New(),
			},
			wantErr: true,
		},
		{
			name: "error delete",
			fields: fields{
				orderElemRep: mocks2.NewIOrderElementRepositoryMock(mc).GetByIDMock.Return(elem, nil).
					DeleteMock.Return(myErrors.ErrDelete).InsertMock.Return(nil),
				orderRep:   mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, nil),
				wineRep:    mocks2.NewIWineRepositoryMock(mc).GetWineMock.Return(w, nil),
				orderLogic: mocks2.NewIOrderControllerMock(mc).UpdateMock.Return(nil),
			},
			args: args{
				ctx: myContext.UserToContext(context.Background(), &models.User{
					ID:       uuid.New(),
					Login:    "veronika",
					Password: "qwertyuiop",
					Fio:      "obrevskaya",
					Points:   10000,
					Email:    "1@mail.ru",
					Status:   models.Customer,
				}),
				ID: uuid.New(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			el := &OrderElemController{
				orderElemRep: tt.fields.orderElemRep,
				orderRep:     tt.fields.orderRep,
				wineRep:      tt.fields.wineRep,
				orderLogic:   tt.fields.orderLogic,
			}
			if err := el.Delete(tt.args.ctx, tt.args.ID); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
