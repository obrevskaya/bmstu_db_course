package main

import (
	"console/controller"
	"fmt"
	"net/http"
)

func main() {
	client := &http.Client{}

	menu := controller.NewMenu()
	err := menu.RunMenu(client)
	if err != nil {
		fmt.Println(err)
	}
}
