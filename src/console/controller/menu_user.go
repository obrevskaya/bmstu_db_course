package controller

import (
	myErrors "console/errors"
	"console/handlers"
	openapi "console/internal/client"
	"console/utils"
	"console/view"
	"errors"
	"fmt"
	"net/http"

	"github.com/fatih/color"
)

var (
	errExit = errors.New("exit")
)

func (m *Menu) userMenu(client *http.Client) error {
	var num int = 1
	var err error

	for num != 0 {

		view.PrintGuestMenu()
		fmt.Scanf("%d", &num)

		switch num {
		case 0:
			return nil
		case 1:
			err = m.register(client)
			if err != nil {
				c := color.New(color.FgRed)
				c.Println(err)
				continue
			} else {
				fmt.Printf("\nПользователь зарегестрирован!\n\n")
			}
		case 2:
			err = m.authorize(client)
			if err != nil {
				c := color.New(color.FgRed)
				c.Println(err)
				continue
			} else {
				fmt.Printf("Пользователь авторизирован!\n\n")
			}
		case 3:
			err = m.getWines(client)
			if err != nil {
				c := color.New(color.FgRed)
				c.Println(err)
			} else {
				c := color.New(color.FgGreen)
				c.Printf("Успешно получены вина.\n")
			}
			continue
		default:
			c := color.New(color.FgRed)
			c.Println(myErrors.ErrorCase)
			continue
		}
		return m.userLoop(client)
	}
	return nil
}

func (m *Menu) userLoop(client *http.Client) error {
	var num int = 1
	var err error

	for num != 0 {
		view.PrintUserMenu()

		fmt.Scanf("%d", &num)

		if num == 0 {
			return nil
		}

		switch num {
		case 1:
			err := m.logout()
			if err != nil {
				fmt.Println(err)
			} else {
				c := color.New(color.FgGreen)
				c.Printf("Успешно разлогирован.\n")
				return nil
			}
		case 2:
			err = m.getWines(client)
			if err != nil {
				fmt.Println(err)
			} else {
				c := color.New(color.FgGreen)
				c.Printf("Успешно получены вина.\n")
			}
		case 3:
			err = m.createElem(client)
			if err != nil {
				fmt.Println(err)
			} else {
				c := color.New(color.FgGreen)
				c.Printf("Успешно создан элемент заказа.\n")
			}
		case 4:
			err = m.addElem(client)
			if err != nil {
				fmt.Println(err)
			} else {
				c := color.New(color.FgGreen)
				c.Printf("Успешно увеличено количество вина в заказе.\n")
			}
		case 5:
			err = m.decreaseElem(client)
			if err != nil {
				fmt.Println(err)
			} else {
				c := color.New(color.FgGreen)
				c.Printf("Усппешно уменьшено количество вина в заказе.\n")
			}
		case 6:
			err = m.deleteElem(client)
			if err != nil {
				fmt.Println(err)
			} else {
				c := color.New(color.FgGreen)
				c.Printf("Успешно удален элемент заказа.\n")
			}
		case 7:
			err = m.getOrder(client)
			if err != nil {
				fmt.Println(err)
			} else {
				c := color.New(color.FgGreen)
				c.Printf("Успешно получен заказ.\n")
			}
		case 8:
			err = m.placeOrder(client)
			if err != nil {
				fmt.Println(err)
			} else {
				c := color.New(color.FgGreen)
				c.Printf("Успешно оформлен заказ.\n")
			}
		case 9:
			err = m.createUserWine(client)
			if err != nil {
				fmt.Println(err)
			} else {
				c := color.New(color.FgGreen)
				c.Printf("Успешно добавлено в любимое.\n")
			}
		case 10:
			err = m.getUserWines(client)
			if err != nil {
				fmt.Println(err)
			} else {
				c := color.New(color.FgGreen)
				c.Printf("Успешно получены любимые вина.\n")
			}
		case 11:
			err = m.deleteUserWine(client)
			if err != nil {
				fmt.Println(err)
			} else {
				c := color.New(color.FgGreen)
				c.Printf("Успешно удалено из любимого.\n")
			}
		default:
			return myErrors.ErrorCase
		}
	}
	return nil
}

func (m *Menu) authorize(client *http.Client) error {
	authRequest, err := view.InputAuth()
	if err != nil {
		return err
	}

	_, err = handlers.AuthorizeClient(client, authRequest)
	if err != nil {
		return err
	}

	m.login = authRequest.Login
	m.password = authRequest.Password
	return nil
}

func (m *Menu) register(client *http.Client) error {
	registerRequest, err := view.InputUser()
	if err != nil {
		return err
	}

	response, err := handlers.CreateClient(client, registerRequest)
	if err != nil {
		return err
	}

	fmt.Println(response.Body)

	m.login = registerRequest.Login
	m.password = registerRequest.Password
	return nil
}

func (m *Menu) logout() error {
	m.login = ""
	m.password = ""
	return nil
}

func (m *Menu) getWines(client *http.Client) error {
	winesRequest, err := view.InputLimitSkip()
	if err != nil {
		return err
	}

	response, err := handlers.GetWines(client, winesRequest)
	if err != nil {
		return err
	}

	wines, err := utils.ParseWinesBody(response)
	if err != nil {
		return err
	}

	view.PrintWines(wines)

	return nil
}

func (m *Menu) createElem(client *http.Client) error {
	request, err := view.InputElem()
	if err != nil {
		return err
	}

	_, err = handlers.CreateElem(client, request, m.login, m.password)
	if err != nil {
		return err
	}

	return nil
}

func (m *Menu) addElem(client *http.Client) error {
	id, err := view.InputID()
	if err != nil {
		return err
	}

	request := &openapi.AddElemRequest{}
	request.Id = id

	_, err = handlers.AddElem(client, request, m.login, m.password)
	if err != nil {
		return err
	}

	return nil
}

func (m *Menu) decreaseElem(client *http.Client) error {
	id, err := view.InputID()
	if err != nil {
		return err
	}

	request := &openapi.DecreaseElemRequest{}
	request.Id = id

	_, err = handlers.DecreaseElem(client, request, m.login, m.password)
	if err != nil {
		return err
	}

	return nil
}

func (m *Menu) deleteElem(client *http.Client) error {
	id, err := view.InputID()
	if err != nil {
		return err
	}

	_, err = handlers.DeleteElem(client, id, m.login, m.password)
	if err != nil {
		return err
	}

	return nil
}

func (m *Menu) getOrder(client *http.Client) error {
	id, err := view.InputID()
	if err != nil {
		return err
	}

	response, err := handlers.GetOrder(client, id, m.login, m.password)
	if err != nil {
		return err
	}

	order, err := utils.ParseOrderBody(response)
	if err != nil {
		return err
	}

	view.PrintOrder(order)

	response, err = handlers.GetByOrder(client, id, m.login, m.password)

	elems, err := utils.ParseElemsBody(response)
	if err != nil {
		return err
	}
	for _, el := range elems {
		response, err = handlers.GetWine(client, el.IdWine, m.login, m.password)
		if err != nil {
			return err
		}

		wine, err := utils.ParseWineBody(response)
		if err != nil {
			return err
		}

		view.PrintWineInOrder(wine, &el)
	}
	return nil
}

func (m *Menu) placeOrder(client *http.Client) error {
	id, err := view.InputID()
	if err != nil {
		return err
	}

	response, err := handlers.GetOrder(client, id, m.login, m.password)
	if err != nil {
		return err
	}

	order, err := utils.ParseOrderBody(response)
	if err != nil {
		return err
	}

	order.IsPoints, err = view.InputIsPoints()
	if err != nil {
		return err
	}

	_, err = handlers.PlaceOrder(client, order, m.login, m.password)
	if err != nil {
		return err
	}

	return nil
}

func (m *Menu) createUserWine(client *http.Client) error {
	idWine, err := view.InputID()
	if err != nil {
		return err
	}

	authRequest := &openapi.AuthRequest{Login: m.login, Password: m.password}
	response, err := handlers.AuthorizeClient(client, authRequest)
	if err != nil {
		return err
	}

	user, err := utils.ParseUserBody(response)
	if err != nil {
		return err
	}

	_, err = handlers.CreateUserWine(client, user.Id, idWine, m.login, m.password)
	if err != nil {
		return err
	}

	return nil
}

func (m *Menu) deleteUserWine(client *http.Client) error {
	idWine, err := view.InputID()
	if err != nil {
		return err
	}

	authRequest := &openapi.AuthRequest{Login: m.login, Password: m.password}
	response, err := handlers.AuthorizeClient(client, authRequest)
	if err != nil {
		return err
	}

	user, err := utils.ParseUserBody(response)
	if err != nil {
		return err
	}

	_, err = handlers.DeleteUserWine(client, user.Id, idWine, m.login, m.password)
	if err != nil {
		return err
	}

	return nil
}

func (m *Menu) getUserWines(client *http.Client) error {
	authRequest := &openapi.AuthRequest{Login: m.login, Password: m.password}
	response, err := handlers.AuthorizeClient(client, authRequest)
	if err != nil {
		return err
	}

	user, err := utils.ParseUserBody(response)
	if err != nil {
		return err
	}

	response, err = handlers.GetUserWines(client, user.Id, m.login, m.password)
	if err != nil {
		return err
	}

	userWines, err := utils.ParseUserWinesBody(response)
	if err != nil {
		return err
	}

	view.PrintUserWines(userWines)

	return nil
}
