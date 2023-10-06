package menu

import (
	"context"
	"fmt"

	mycontext "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/context"
	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/logic/errors"
	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/logic/models"
	"github.com/dixonwille/wmenu"
	"github.com/fatih/color"
	"github.com/google/uuid"
)

func printWines(wines []*models.Wine) {
	c := color.New(color.FgCyan)
	for i, wine := range wines {
		c.Printf("Wine №%d:\n", i+1)
		fmt.Printf("\tID: %s\n", wine.ID)
		fmt.Printf("\tName: %s\n", wine.Name)
		fmt.Printf("\tCount: %d\n", wine.Count)
		fmt.Printf("\tYear: %d\n", wine.Year)
		fmt.Printf("\tStrength: %d\n", wine.Strength)
		fmt.Printf("\tPrice: %d\n", wine.Price)
		fmt.Printf("\tType: %s\n", wine.Type)
	}
}

func printOrder(o *models.Order) {
	fmt.Printf("\tID: %s\n", o.ID)
	fmt.Printf("\tTotal Price: %d\n", o.TotalPrice)
	fmt.Printf("\tIs Points: %t\n", o.IsPoints)
	fmt.Printf("\tStatus: %s\n", o.Status)
}

func (m *Menu) registerHandler(opt wmenu.Opt) error {
	user := models.User{}

	fmt.Print("Enter login: ")
	if _, err := fmt.Scan(&user.Login); err != nil {
		return fmt.Errorf("input login: %w", err)
	}

	fmt.Print("Enter password: ")
	if _, err := fmt.Scan(&user.Password); err != nil {
		return fmt.Errorf("input password: %w", err)
	}

	fmt.Print("Enter name: ")
	if _, err := fmt.Scan(&user.Fio); err != nil {
		return fmt.Errorf("input name: %w", err)
	}

	fmt.Print("Enter email: ")
	if _, err := fmt.Scan(&user.Email); err != nil {
		return fmt.Errorf("input email: %w", err)
	}

	user.Status = models.Customer
	err := m.userController.Create(opt.Value.(context.Context), &user)
	if err != nil {
		return fmt.Errorf("register: %w", err)
	}

	c := color.New(color.FgGreen)
	c.Printf("Successful registration.")

	m.login = user.Login
	m.password = user.Password

	return nil
}

func (m *Menu) authorizeHandler(opt wmenu.Opt) error {
	var login, password string

	fmt.Print("Enter login: ")
	if _, err := fmt.Scan(&login); err != nil {
		return fmt.Errorf("input login: %w", err)
	}

	fmt.Print("Enter password: ")
	if _, err := fmt.Scan(&password); err != nil {
		return fmt.Errorf("input password: %w", err)
	}

	user, err := m.userController.Authorize(opt.Value.(context.Context), login, password)
	if err != nil {
		return fmt.Errorf("authorize: %w", err)
	}

	c := color.New(color.FgGreen)
	c.Printf("Successful authorize.")

	m.login = user.Login
	m.password = user.Password

	return nil
}

func (m *Menu) logoutHandler(_ wmenu.Opt) error {

	m.login = ""
	m.password = ""

	c := color.New(color.FgGreen)
	c.Printf("Successful logout\n")

	return nil
}

func (m *Menu) getOrderHandler(opt wmenu.Opt) error {
	user, err := m.userController.Authorize(opt.Value.(context.Context), m.login, m.password)
	if err != nil {
		return fmt.Errorf("authorize: %w", err)
	}
	ctx := mycontext.UserToContext(opt.Value.(context.Context), user)

	var idString string
	fmt.Print("Enter ID: ")
	if _, err := fmt.Scan(&idString); err != nil {
		return fmt.Errorf("input id order: %w", err)
	}

	id, err := uuid.Parse(idString)
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	order, err := m.orderController.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("get order: %w", err)
	}
	els, err := m.elementController.GetByOrder(ctx, id)
	if err != nil {
		return fmt.Errorf("get elements: %w", err)
	}

	c := color.New(color.FgCyan)
	for i, el := range els {
		c.Printf("order element №%d:\n", i+1)
		fmt.Printf("\tIDWine: %s\n", el.IDWine)
		wine, err := m.wineController.GetWine(ctx, el.IDWine)
		if err != nil {
			return fmt.Errorf("get wine in order: %w", err)
		}
		fmt.Printf("\tName Wine: %s\n", wine.Name)
		fmt.Printf("\tCount: %d\n", el.Count)
	}
	printOrder(order)

	c = color.New(color.FgGreen)
	c.Printf("Successful get order.")

	return nil
}

func (m *Menu) getWinesHandler(opt wmenu.Opt) error {
	var limit, skip int
	fmt.Print("Enter limit: ")
	if _, err := fmt.Scan(&limit); err != nil {
		return fmt.Errorf("input limit wine: %w", err)
	}

	fmt.Print("Enter skip: ")
	if _, err := fmt.Scan(&skip); err != nil {
		return fmt.Errorf("input skip wine: %w", err)
	}

	wines, err := m.wineController.GetWines(opt.Value.(context.Context), limit, skip)
	if err != nil {
		return fmt.Errorf("get wines: %w", err)
	}

	printWines(wines)

	c := color.New(color.FgGreen)
	c.Printf("Successful get wines.")

	return nil
}

func (m *Menu) createElementHandler(opt wmenu.Opt) error {
	user, err := m.userController.Authorize(opt.Value.(context.Context), m.login, m.password)
	if err != nil {
		return fmt.Errorf("authorize: %w", err)
	}
	ctx := mycontext.UserToContext(opt.Value.(context.Context), user)

	elem := models.OrderElement{Count: 1}

	var idString string
	fmt.Print("Enter id wine: ")
	if _, err := fmt.Scan(&idString); err != nil {
		return fmt.Errorf("input wine uuid: %w", err)
	}

	elem.IDWine, err = uuid.Parse(idString)
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	err = m.elementController.Create(ctx, &elem)
	if err != nil {
		return fmt.Errorf("create element: %w", err)
	}

	c := color.New(color.FgGreen)
	c.Printf("Successful create element.")
	return nil
}

func (m *Menu) addElementHandler(opt wmenu.Opt) error {
	user, err := m.userController.Authorize(opt.Value.(context.Context), m.login, m.password)
	if err != nil {
		return fmt.Errorf("authorize: %w", err)
	}
	ctx := mycontext.UserToContext(opt.Value.(context.Context), user)

	var idString string
	fmt.Print("Enter id elem: ")
	if _, err := fmt.Scan(&idString); err != nil {
		return fmt.Errorf("input elem uuid: %w", err)
	}

	id, err := uuid.Parse(idString)
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	err = m.elementController.Add(ctx, id)
	if err != nil {
		return fmt.Errorf("add element: %w", err)
	}

	c := color.New(color.FgGreen)
	c.Printf("Successful add element.")
	return nil
}

func (m *Menu) decreaseElementHandler(opt wmenu.Opt) error {
	user, err := m.userController.Authorize(opt.Value.(context.Context), m.login, m.password)
	if err != nil {
		return fmt.Errorf("authorize: %w", err)
	}
	ctx := mycontext.UserToContext(opt.Value.(context.Context), user)

	var idString string
	fmt.Print("Enter id elem: ")
	if _, err := fmt.Scan(&idString); err != nil {
		return fmt.Errorf("input elem uuid: %w", err)
	}

	id, err := uuid.Parse(idString)
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	err = m.elementController.Decrease(ctx, id)
	if err != nil {
		return fmt.Errorf("decrease element: %w", err)
	}

	c := color.New(color.FgGreen)
	c.Printf("Successful decrease element.")
	return nil
}

func (m *Menu) deleteElementHandler(opt wmenu.Opt) error {
	user, err := m.userController.Authorize(opt.Value.(context.Context), m.login, m.password)
	if err != nil {
		return fmt.Errorf("authorize: %w", err)
	}
	ctx := mycontext.UserToContext(opt.Value.(context.Context), user)

	var idString string
	fmt.Print("Enter id elem: ")
	if _, err := fmt.Scan(&idString); err != nil {
		return fmt.Errorf("input elem uuid: %w", err)
	}

	id, err := uuid.Parse(idString)
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	err = m.elementController.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("delete element: %w", err)
	}

	c := color.New(color.FgGreen)
	c.Printf("Successful delete element.")
	return nil
}

func (m *Menu) createOrderHandler(opt wmenu.Opt) error {
	user, err := m.userController.Authorize(opt.Value.(context.Context), m.login, m.password)
	if err != nil {
		return fmt.Errorf("authorize: %w", err)
	}
	ctx := mycontext.UserToContext(opt.Value.(context.Context), user)

	order := &models.Order{
		IDUser:     user.ID,
		TotalPrice: 0,
		IsPoints:   false,
		Status:     models.ProcessOrder,
	}

	err = m.orderController.Create(ctx, order)
	if err != nil {
		return fmt.Errorf("create order: %w", err)
	}

	c := color.New(color.FgGreen)
	c.Printf("Successful create order.")
	return nil
}

func (m *Menu) placeOrderHandler(opt wmenu.Opt) error {
	user, err := m.userController.Authorize(opt.Value.(context.Context), m.login, m.password)
	if err != nil {
		return fmt.Errorf("authorize: %w", err)
	}
	ctx := mycontext.UserToContext(opt.Value.(context.Context), user)

	var idString string
	fmt.Print("Enter id order: ")
	if _, err := fmt.Scan(&idString); err != nil {
		return fmt.Errorf("input order uuid: %w", err)
	}

	id, err := uuid.Parse(idString)
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	order, err := m.orderController.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("get order: %w", err)
	}

	fmt.Print("Pay with points(yes or no): ")
	var isPoints string
	if _, err = fmt.Scan(&isPoints); err != nil {
		return fmt.Errorf("input ispoints: %w", err)
	}
	if isPoints == "yes" || isPoints == "y" {
		order.IsPoints = true
	} else if isPoints == "no" || isPoints == "n" {
		order.IsPoints = false
	} else {
		return fmt.Errorf("incorrect choice of variant")
	}

	order.Status = models.PlacedOrder
	err = m.orderController.Update(ctx, order)
	if err != nil {
		return fmt.Errorf("update order: %w", err)
	}

	c := color.New(color.FgGreen)
	c.Printf("Successful place order.")
	return nil
}

func (m *Menu) addWineHandler(opt wmenu.Opt) error {
	user, err := m.userController.Authorize(opt.Value.(context.Context), m.login, m.password)
	if err != nil {
		return fmt.Errorf("authorize: %w", err)
	}
	if user.Status != models.Admin {
		return errors.ErrAccess
	}
	ctx := mycontext.UserToContext(opt.Value.(context.Context), user)

	wine := models.Wine{}
	fmt.Print("Enter name: ")
	if _, err := fmt.Scan(&wine.Name); err != nil {
		return fmt.Errorf("input name: %w", err)
	}
	fmt.Print("Enter count: ")
	if _, err := fmt.Scan(&wine.Count); err != nil {
		return fmt.Errorf("input count: %w", err)
	}
	fmt.Print("Enter year: ")
	if _, err := fmt.Scan(&wine.Year); err != nil {
		return fmt.Errorf("input year: %w", err)
	}
	fmt.Print("Enter strength: ")
	if _, err := fmt.Scan(&wine.Strength); err != nil {
		return fmt.Errorf("input strength: %w", err)
	}
	fmt.Print("Enter price: ")
	if _, err := fmt.Scan(&wine.Price); err != nil {
		return fmt.Errorf("input price: %w", err)
	}
	fmt.Print("Enter type: ")
	if _, err := fmt.Scan(&wine.Type); err != nil {
		return fmt.Errorf("input type: %w", err)
	}
	err = m.wineController.Create(ctx, &wine)
	if err != nil {
		return fmt.Errorf("create wine: %w", err)
	}

	c := color.New(color.FgGreen)
	c.Printf("Successful add wine.")
	return nil
}

func (m *Menu) payBillHandler(opt wmenu.Opt) error {
	user, err := m.userController.Authorize(opt.Value.(context.Context), m.login, m.password)
	if err != nil {
		return fmt.Errorf("authorize: %w", err)
	}
	if user.Status != models.Admin {
		return errors.ErrAccess
	}
	ctx := mycontext.UserToContext(opt.Value.(context.Context), user)

	var idString string
	fmt.Print("Enter id bill: ")
	if _, err := fmt.Scan(&idString); err != nil {
		return fmt.Errorf("input bill uuid: %w", err)
	}

	id, err := uuid.Parse(idString)
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	err = m.billController.UpdateBillStatus(ctx, id, models.PaidBill)
	if err != nil {
		return fmt.Errorf("update bill status: %w", err)
	}

	c := color.New(color.FgGreen)
	c.Printf("Successful pay bill.")
	return nil
}

func (m *Menu) deleteWineHandler(opt wmenu.Opt) error {
	user, err := m.userController.Authorize(opt.Value.(context.Context), m.login, m.password)
	if err != nil {
		return fmt.Errorf("authorize: %w", err)
	}
	if user.Status != models.Admin {
		return errors.ErrAccess
	}
	ctx := mycontext.UserToContext(opt.Value.(context.Context), user)

	var idString string
	fmt.Print("Enter id wine: ")
	if _, err := fmt.Scan(&idString); err != nil {
		return fmt.Errorf("input wine uuid: %w", err)
	}

	id, err := uuid.Parse(idString)
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	err = m.wineController.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("delete wine: %w", err)
	}

	c := color.New(color.FgGreen)
	c.Printf("Successful delete wine.")
	return nil
}

func (m *Menu) updateWineHandler(opt wmenu.Opt) error {
	user, err := m.userController.Authorize(opt.Value.(context.Context), m.login, m.password)
	if err != nil {
		return fmt.Errorf("authorize: %w", err)
	}
	if user.Status != models.Admin {
		return errors.ErrAccess
	}
	ctx := mycontext.UserToContext(opt.Value.(context.Context), user)

	var idString string
	fmt.Print("Enter id wine: ")
	if _, err := fmt.Scan(&idString); err != nil {
		return fmt.Errorf("input wine uuid: %w", err)
	}

	id, err := uuid.Parse(idString)
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	wine, err := m.wineController.GetWine(ctx, id)
	if err != nil {
		return fmt.Errorf("get wine: %w", err)
	}

	fmt.Print("Enter new name (old value ", wine.Name, "): ")
	if _, err := fmt.Scan(&wine.Name); err != nil {
		return fmt.Errorf("input name: %w", err)
	}
	fmt.Print("Enter new count (old value ", wine.Count, "): ")
	if _, err := fmt.Scan(&wine.Count); err != nil {
		return fmt.Errorf("input count: %w", err)
	}
	fmt.Print("Enter new year (old value ", wine.Year, "): ")
	if _, err := fmt.Scan(&wine.Year); err != nil {
		return fmt.Errorf("input year: %w", err)
	}
	fmt.Print("Enter new strength (old value ", wine.Strength, "): ")
	if _, err := fmt.Scan(&wine.Strength); err != nil {
		return fmt.Errorf("input strength: %w", err)
	}
	fmt.Print("Enter new price (old value ", wine.Price, "): ")
	if _, err := fmt.Scan(&wine.Price); err != nil {
		return fmt.Errorf("input price: %w", err)
	}
	fmt.Print("Enter new type (old value ", wine.Type, "): ")
	if _, err := fmt.Scan(&wine.Type); err != nil {
		return fmt.Errorf("input : %w", err)
	}

	err = m.wineController.Update(ctx, wine)
	if err != nil {
		return fmt.Errorf("update wine: %w", err)
	}
	c := color.New(color.FgGreen)
	c.Printf("Successful update wine.")
	return nil
}

func (m *Menu) updateUserPointsHandler(opt wmenu.Opt) error {
	user, err := m.userController.Authorize(opt.Value.(context.Context), m.login, m.password)
	if err != nil {
		return fmt.Errorf("authorize: %w", err)
	}
	if user.Status != models.Admin {
		return errors.ErrAccess
	}
	ctx := mycontext.UserToContext(opt.Value.(context.Context), user)

	var idString string
	fmt.Print("Enter id user: ")
	if _, err := fmt.Scan(&idString); err != nil {
		return fmt.Errorf("input user uuid: %w", err)
	}

	id, err := uuid.Parse(idString)
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	var points int
	fmt.Print("Enter points: ")
	if _, err := fmt.Scan(&points); err != nil {
		return fmt.Errorf("input points: %w", err)
	}

	err = m.userController.UpdateUserPoints(ctx, id, points)
	if err != nil {
		return fmt.Errorf("update user points: %w", err)
	}
	c := color.New(color.FgGreen)
	c.Printf("Successful update user points.")
	return nil
}
