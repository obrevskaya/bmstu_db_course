package controller

import (
	myErrors "console/errors"
	"console/handlers"
	"console/internal/consts"
	"console/utils"
	"console/view"
	"fmt"
	"net/http"
	"strconv"

	"github.com/fatih/color"
)

func (m *Menu) adminMenu(client *http.Client) error {

	err := m.authorizeAdmin(client)
	if err != nil {
		return err
	}
	fmt.Printf("Admin authorize!\n\n")

	return m.adminLoop(client)
}

func (m *Menu) adminLoop(client *http.Client) error {
	var num int = 1
	var err error

	for num != 0 {
		view.PrintAdminMenu()

		fmt.Scanf("%d", &num)

		if num == 0 {
			return nil
		}

		switch num {
		case 1:
			err = m.registerAdmin(client)
			if err != nil {
				fmt.Println(err)
			} else {
				c := color.New(color.FgGreen)
				c.Printf("Successfully admin register.\n")
			}
		case 2:
			err = m.getWines(client)
			if err != nil {
				fmt.Println(err)
			} else {
				c := color.New(color.FgGreen)
				c.Printf("Successfully get wines.\n")
			}
		case 3:
			err = m.confirmPay(client)
			if err != nil {
				fmt.Println(err)
			} else {
				c := color.New(color.FgGreen)
				c.Printf("Successfully confirm pay.\n")
			}
		case 4:
			err = m.addWine(client)
			if err != nil {
				fmt.Println(err)
			} else {
				c := color.New(color.FgGreen)
				c.Printf("Successfully add wine.\n")
			}
		case 5:
			err = m.deleteWine(client)
			if err != nil {
				fmt.Println(err)
			} else {
				c := color.New(color.FgGreen)
				c.Printf("Successfully delete wine.\n")
			}
		case 6:
			err = m.updateWine(client)
			if err != nil {
				fmt.Println(err)
			} else {
				c := color.New(color.FgGreen)
				c.Printf("Successfully update wine.\n")
			}
		default:
			return myErrors.ErrorCase
		}
	}
	return nil
}

func (m *Menu) authorizeAdmin(client *http.Client) error {
	authRequest, err := view.InputAuth()
	if err != nil {
		return err
	}

	response, err := handlers.AuthorizeClient(client, authRequest)
	if err != nil {
		return err
	}

	user, err := utils.ParseUserBody(response)
	if err != nil {
		return err
	}
	if user.Status != strconv.Itoa(consts.Admin) {
		return myErrors.ErrorAccess
	}

	m.login = authRequest.Login
	m.password = authRequest.Password
	return nil
}

func (m *Menu) registerAdmin(client *http.Client) error {
	registerRequest, err := view.InputUser()
	if err != nil {
		return err
	}

	registerRequest.Status = strconv.Itoa(consts.Admin)

	_, err = handlers.CreateClient(client, registerRequest)
	if err != nil {
		return err
	}

	m.login = registerRequest.Login
	m.password = registerRequest.Password
	return nil
}

func (m *Menu) confirmPay(client *http.Client) error {
	id, err := view.InputID()
	if err != nil {
		return err
	}

	_, err = handlers.PayBill(client, id, m.login, m.password)
	if err != nil {
		return err
	}

	return nil
}

func (m *Menu) addWine(client *http.Client) error {
	wine, err := view.InputWine()
	if err != nil {
		return err
	}

	_, err = handlers.AddWine(client, wine, m.login, m.password)
	if err != nil {
		return err
	}

	return nil
}

func (m *Menu) deleteWine(client *http.Client) error {
	id, err := view.InputID()
	if err != nil {
		return err
	}

	_, err = handlers.DeleteWine(client, id, m.login, m.password)
	if err != nil {
		return err
	}

	return nil
}

func (m *Menu) updateWine(client *http.Client) error {
	id, err := view.InputID()
	if err != nil {
		return err
	}

	response, err := handlers.GetWine(client, id, m.login, m.password)
	if err != nil {
		return err
	}
	oldWine, err := utils.ParseWineBody(response)
	if err != nil {
		return err
	}

	wine, err := view.UpdateWine(oldWine)
	if err != nil {
		return err
	}

	_, err = handlers.UpdateWine(client, wine, m.login, m.password)
	if err != nil {
		return err
	}

	return nil
}
