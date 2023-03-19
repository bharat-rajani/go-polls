package main

import (
	"context"
	"github.com/bharat-rajani/go-polls/internal/configurer"
)

func main() {
	err := configurer.StartAPIService(context.Background())
	if err != nil {
		panic(err)
	}
}
