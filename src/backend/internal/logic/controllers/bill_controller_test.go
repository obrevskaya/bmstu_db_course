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

func TestBill_Create(t *testing.T) {
	mc := minimock.NewController(t)
	o := &models.Order{
		ID:         uuid.New(),
		IDUser:     uuid.New(),
		TotalPrice: 100,
		IsPoints:   false,
		Status:     "placed",
	}
	type fields struct {
		billRep  interfaces.IBillRepository
		userRep  interfaces.IUserRepository
		orderRep interfaces.IOrderRepository
	}
	type args struct {
		ctx  context.Context
		bill *models.Bill
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "successful create bill",
			fields: fields{
				orderRep: mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, nil),
				billRep:  mocks2.NewIBillRepositoryMock(mc).InsertMock.Return(nil),
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
				bill: &models.Bill{
					IDOrder: uuid.New(),
					Price:   100,
					Status:  models.NotPaidBill,
				},
			},
			wantErr: false,
		},
		{
			name: "no context in create bill",
			fields: fields{
				orderRep: mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, nil),
				billRep:  mocks2.NewIBillRepositoryMock(mc).InsertMock.Return(nil),
				userRep:  nil,
			},
			args: args{
				ctx: context.Background(),
				bill: &models.Bill{
					IDOrder: uuid.New(),
					Price:   100,
					Status:  models.NotPaidBill,
				},
			},
			wantErr: true,
		},
		{
			name: "error get order",
			fields: fields{
				orderRep: mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, myErrors.ErrGetDB),
				billRep:  mocks2.NewIBillRepositoryMock(mc).InsertMock.Return(nil),
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
				bill: &models.Bill{
					IDOrder: uuid.New(),
					Price:   100,
					Status:  models.NotPaidBill,
				},
			},
			wantErr: true,
		},
		{
			name: "error other user",
			fields: fields{
				orderRep: mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, nil),
				billRep:  mocks2.NewIBillRepositoryMock(mc).InsertMock.Return(nil),
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
				bill: &models.Bill{
					IDOrder: uuid.New(),
					Price:   100,
					Status:  models.NotPaidBill,
				},
			},
			wantErr: true,
		},
		{
			name: "error insert bill",
			fields: fields{
				orderRep: mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, nil),
				billRep:  mocks2.NewIBillRepositoryMock(mc).InsertMock.Return(myErrors.ErrInsert),
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
				bill: &models.Bill{
					IDOrder: uuid.New(),
					Price:   100,
					Status:  models.NotPaidBill,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BillController{
				billRep:  tt.fields.billRep,
				userRep:  tt.fields.userRep,
				orderRep: tt.fields.orderRep,
			}
			if err := b.Create(tt.args.ctx, tt.args.bill); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBill_Get(t *testing.T) {
	mc := minimock.NewController(t)

	b := &models.Bill{
		ID:      uuid.New(),
		IDOrder: uuid.New(),
		Price:   100,
		Status:  models.NotPaidBill,
	}

	type fields struct {
		billRep  interfaces.IBillRepository
		userRep  interfaces.IUserRepository
		orderRep interfaces.IOrderRepository
	}
	type args struct {
		ctx context.Context
		ID  uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Bill
		wantErr bool
	}{
		{
			name: "successful get bill",
			fields: fields{
				orderRep: nil,
				billRep:  mocks2.NewIBillRepositoryMock(mc).GetMock.Return(b, nil),
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
			want:    b,
			wantErr: false,
		},
		{
			name: "no context in get bill",
			fields: fields{
				orderRep: nil,
				billRep:  mocks2.NewIBillRepositoryMock(mc).GetMock.Return(b, nil),
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
			name: "error get bill",
			fields: fields{
				orderRep: nil,
				billRep:  mocks2.NewIBillRepositoryMock(mc).GetMock.Return(b, myErrors.ErrGetDB),
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
			b := &BillController{
				billRep:  tt.fields.billRep,
				userRep:  tt.fields.userRep,
				orderRep: tt.fields.orderRep,
			}
			got, err := b.Get(tt.args.ctx, tt.args.ID)
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

func TestBill_Update(t *testing.T) {
	mc := minimock.NewController(t)
	b := &models.Bill{
		ID:      uuid.New(),
		IDOrder: uuid.New(),
		Price:   100,
		Status:  models.NotPaidBill,
	}
	bPaid := &models.Bill{
		ID:      uuid.New(),
		IDOrder: uuid.New(),
		Price:   100,
		Status:  models.PaidBill,
	}
	o := &models.Order{
		ID:         uuid.New(),
		IDUser:     uuid.New(),
		TotalPrice: 100,
		IsPoints:   false,
		Status:     "placed",
	}

	u := &models.User{
		ID:       uuid.New(),
		Login:    "ty",
		Password: "fdhhjgjghfh",
		Fio:      "Obrevskaya Veronika",
		Email:    "obreaya.vera@mail.ru",
		Points:   0,
		Status:   models.Customer,
	}

	type fields struct {
		billRep  interfaces.IBillRepository
		userRep  interfaces.IUserRepository
		orderRep interfaces.IOrderRepository
	}
	type args struct {
		ctx    context.Context
		ID     uuid.UUID
		status string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "successful update bill",
			fields: fields{
				orderRep: mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, nil),
				billRep:  mocks2.NewIBillRepositoryMock(mc).GetMock.Return(b, nil).UpdateBillStatusMock.Return(nil),
				userRep:  mocks2.NewIUserRepositoryMock(mc).GetMock.Return(u, nil).UpdateUserPointsMock.Return(nil),
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
				ID:     uuid.New(),
				status: models.PaidBill,
			},
			wantErr: false,
		},
		{
			name: "bill also paid",
			fields: fields{
				orderRep: mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, nil),
				billRep:  mocks2.NewIBillRepositoryMock(mc).GetMock.Return(bPaid, nil).UpdateBillStatusMock.Return(nil),
				userRep:  mocks2.NewIUserRepositoryMock(mc).GetMock.Return(u, nil).UpdateUserPointsMock.Return(nil),
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
				ID:     uuid.New(),
				status: models.PaidBill,
			},
			wantErr: true,
		},
		{
			name: "no context in update bill",
			fields: fields{
				orderRep: mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, nil),
				billRep:  mocks2.NewIBillRepositoryMock(mc).GetMock.Return(b, nil).UpdateBillStatusMock.Return(nil),
				userRep:  mocks2.NewIUserRepositoryMock(mc).GetMock.Return(u, nil).UpdateUserPointsMock.Return(nil),
			},
			args: args{
				ctx:    context.Background(),
				ID:     uuid.New(),
				status: models.PaidBill,
			},
			wantErr: true,
		},
		{
			name: "no access rights update bill",
			fields: fields{
				orderRep: mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, nil),
				billRep:  mocks2.NewIBillRepositoryMock(mc).GetMock.Return(b, nil).UpdateBillStatusMock.Return(nil),
				userRep:  mocks2.NewIUserRepositoryMock(mc).GetMock.Return(u, nil).UpdateUserPointsMock.Return(nil),
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
				ID:     uuid.New(),
				status: models.PaidBill,
			},
			wantErr: true,
		},
		{
			name: "error get bill",
			fields: fields{
				orderRep: mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, nil),
				billRep:  mocks2.NewIBillRepositoryMock(mc).GetMock.Return(b, myErrors.ErrGetDB).UpdateBillStatusMock.Return(nil),
				userRep:  mocks2.NewIUserRepositoryMock(mc).GetMock.Return(u, nil).UpdateUserPointsMock.Return(nil),
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
				ID:     uuid.New(),
				status: models.PaidBill,
			},
			wantErr: true,
		},
		{
			name: "error order get",
			fields: fields{
				orderRep: mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, myErrors.ErrGetDB),
				billRep:  mocks2.NewIBillRepositoryMock(mc).GetMock.Return(b, nil).UpdateBillStatusMock.Return(nil),
				userRep:  mocks2.NewIUserRepositoryMock(mc).GetMock.Return(u, nil).UpdateUserPointsMock.Return(nil),
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
				ID:     uuid.New(),
				status: models.PaidBill,
			},
			wantErr: true,
		},
		{
			name: "error user get",
			fields: fields{
				orderRep: mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, nil),
				billRep:  mocks2.NewIBillRepositoryMock(mc).GetMock.Return(b, nil).UpdateBillStatusMock.Return(nil),
				userRep:  mocks2.NewIUserRepositoryMock(mc).GetMock.Return(u, myErrors.ErrGetDB).UpdateUserPointsMock.Return(nil),
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
				ID:     uuid.New(),
				status: models.PaidBill,
			},
			wantErr: true,
		},
		{
			name: "error bill update",
			fields: fields{
				orderRep: mocks2.NewIOrderRepositoryMock(mc).GetMock.Return(o, nil),
				billRep:  mocks2.NewIBillRepositoryMock(mc).GetMock.Return(b, nil).UpdateBillStatusMock.Return(myErrors.ErrUpdate),
				userRep:  mocks2.NewIUserRepositoryMock(mc).GetMock.Return(u, nil).UpdateUserPointsMock.Return(nil),
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
				ID:     uuid.New(),
				status: models.PaidBill,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BillController{
				billRep:  tt.fields.billRep,
				userRep:  tt.fields.userRep,
				orderRep: tt.fields.orderRep,
			}
			if err := b.UpdateBillStatus(tt.args.ctx, tt.args.ID, tt.args.status); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
