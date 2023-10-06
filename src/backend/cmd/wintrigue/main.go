package main

import (
	"context"
	"fmt"
	"os"

	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/app/http"
)

func main() {
	a := http.New()

	err := a.Init()
	if err != nil {
		fmt.Fprintf(os.Stderr, "init: %s", err)
		return
	}

	err = a.Run(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "run: %s", err)
		return
	}

}
