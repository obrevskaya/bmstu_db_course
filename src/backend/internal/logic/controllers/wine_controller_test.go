package controllers

import (
	"context"
	"reflect"
	"testing"

	myContext "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/context"
	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/interfaces"
	myErrors "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/logic/errors"
	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/logic/models"
	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/mocks"
	"github.com/gojuno/minimock/v3"
	"github.com/google/uuid"
)

func TestWine_Create(t *testing.T) {
	mc := minimock.NewController(t)

	type fields struct {
		wineRep interfaces.IWineRepository
	}
	type args struct {
		ctx  context.Context
		wine *models.Wine
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "successful create",
			fields: fields{
				wineRep: mocks.NewIWineRepositoryMock(mc).InsertMock.Return(nil),
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
				wine: &models.Wine{
					Name:     "bordo",
					Count:    100,
					Year:     2000,
					Strength: 9,
					Price:    540,
					Type:     "red",
				},
			},
			wantErr: false,
		},
		{
			name: "no context",
			fields: fields{
				wineRep: mocks.NewIWineRepositoryMock(mc).InsertMock.Return(nil),
			},
			args: args{
				ctx: context.Background(),
				wine: &models.Wine{
					Name:     "bordo",
					Count:    100,
					Year:     2000,
					Strength: 9,
					Price:    540,
					Type:     "red",
				},
			},
			wantErr: true,
		},
		{
			name: "no access right",
			fields: fields{
				wineRep: mocks.NewIWineRepositoryMock(mc).InsertMock.Return(nil),
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
				wine: &models.Wine{
					Name:     "bordo",
					Count:    100,
					Year:     2000,
					Strength: 9,
					Price:    540,
					Type:     "red",
				},
			},
			wantErr: true,
		},
		{
			name: "insert error",
			fields: fields{
				wineRep: mocks.NewIWineRepositoryMock(mc).InsertMock.Return(myErrors.ErrInsert),
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
				wine: &models.Wine{
					Name:     "bordo",
					Count:    100,
					Year:     2000,
					Strength: 9,
					Price:    540,
					Type:     "red",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WineController{
				wineRep: tt.fields.wineRep,
			}
			if err := w.Create(tt.args.ctx, tt.args.wine); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWine_Delete(t *testing.T) {
	mc := minimock.NewController(t)
	type fields struct {
		wineRep interfaces.IWineRepository
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
			name:   "successful delete",
			fields: fields{wineRep: mocks.NewIWineRepositoryMock(mc).DeleteMock.Return(nil)},
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
			name:   "no context",
			fields: fields{wineRep: mocks.NewIWineRepositoryMock(mc).DeleteMock.Return(nil)},
			args: args{
				ctx: context.Background(),
				ID:  uuid.New(),
			},
			wantErr: true,
		},
		{
			name:   "no access right",
			fields: fields{wineRep: mocks.NewIWineRepositoryMock(mc).DeleteMock.Return(nil)},
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
			name:   "delete error",
			fields: fields{wineRep: mocks.NewIWineRepositoryMock(mc).DeleteMock.Return(myErrors.ErrDelete)},
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
			w := &WineController{
				wineRep: tt.fields.wineRep,
			}
			if err := w.Delete(tt.args.ctx, tt.args.ID); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWine_GetWine(t *testing.T) {
	mc := minimock.NewController(t)

	w := &models.Wine{
		ID:       uuid.New(),
		Name:     "bordo",
		Count:    100,
		Year:     2000,
		Strength: 9,
		Price:    540,
		Type:     "red",
	}

	type fields struct {
		wineRep interfaces.IWineRepository
	}
	type args struct {
		ctx context.Context
		ID  uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Wine
		wantErr bool
	}{
		{
			name:   "successful get wine",
			fields: fields{wineRep: mocks.NewIWineRepositoryMock(mc).GetWineMock.Return(w, nil)},
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
			want:    w,
			wantErr: false,
		},
		{
			name:   "error get wine",
			fields: fields{wineRep: mocks.NewIWineRepositoryMock(mc).GetWineMock.Return(w, myErrors.ErrGetDB)},
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
			w := &WineController{
				wineRep: tt.fields.wineRep,
			}
			got, err := w.GetWine(tt.args.ctx, tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWine() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWine_Update(t *testing.T) {

	mc := minimock.NewController(t)

	w := &models.Wine{
		ID:    uuid.New(),
		Name:  "bordo",
		Count: 100,
		Price: 540,
	}

	type fields struct {
		wineRep interfaces.IWineRepository
	}
	type args struct {
		ctx  context.Context
		wine *models.Wine
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "successful update",
			fields: fields{wineRep: mocks.NewIWineRepositoryMock(mc).UpdateMock.Return(nil)},
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
				wine: w,
			},
			wantErr: false,
		},
		{
			name:   "no context in update wine",
			fields: fields{wineRep: mocks.NewIWineRepositoryMock(mc).UpdateMock.Return(nil)},
			args: args{
				ctx:  context.Background(),
				wine: w,
			},
			wantErr: true,
		},
		{
			name:   "error update",
			fields: fields{wineRep: mocks.NewIWineRepositoryMock(mc).UpdateMock.Return(myErrors.ErrUpdate)},
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
				wine: w,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WineController{
				wineRep: tt.fields.wineRep,
			}
			if err := w.Update(tt.args.ctx, tt.args.wine); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWineController_GetWines(t *testing.T) {
	mc := minimock.NewController(t)

	w := []*models.Wine{{
		ID:       uuid.New(),
		Name:     "beluga",
		Count:    100,
		Year:     2000,
		Strength: 9,
		Price:    540,
		Type:     "red"},
	}

	type fields struct {
		wineRep interfaces.IWineRepository
	}
	type args struct {
		ctx   context.Context
		limit int
		skip  int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*models.Wine
		wantErr bool
	}{
		{
			name: "successful test get wines",
			fields: fields{
				wineRep: mocks.NewIWineRepositoryMock(mc).GetWinesMock.Return(w, nil),
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
				limit: 10,
				skip:  0,
			},
			want:    w,
			wantErr: false,
		},
		{
			name: "error get data in get wines",
			fields: fields{
				wineRep: mocks.NewIWineRepositoryMock(mc).GetWinesMock.Return(w, myErrors.ErrGet),
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
				limit: 10,
				skip:  0,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WineController{
				wineRep: tt.fields.wineRep,
			}
			got, err := w.GetWines(tt.args.ctx, tt.args.limit, tt.args.skip)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWines() got = %v, want %v", got, tt.want)
			}
		})
	}
}
