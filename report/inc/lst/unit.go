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
