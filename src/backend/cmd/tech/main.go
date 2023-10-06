package main

import (
	"context"
	"fmt"
	"os"

	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/app/tech"
)

func main() {
	a := tech.New()

	err := a.Init()
	if err != nil {
		fmt.Fprintf(os.Stderr, "init: %s", err)
		return
	}

	a.Run(context.Background())
}
