package controller

import (
	"console/errors"
	"console/view"
	"fmt"
	"net/http"

	"github.com/dixonwille/wmenu"
)

type Menu struct {
	mainMenu *wmenu.Menu

	login    string
	password string
}

func NewMenu() *Menu {
	return &Menu{}
}

func (m *Menu) RunMenu(client *http.Client) error {
	view.PrintRunMenu()

	var ch int
	fmt.Scanf("%d", &ch)

	switch ch {
	case 0:
		err := m.userMenu(client)
		if err != nil {
			return err
		}
	case 1:
		err := m.adminMenu(client)
		if err != nil {
			return err
		}
	default:
		return errors.ErrorInput
	}

	return nil
}
