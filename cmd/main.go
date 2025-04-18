package main

import (
	"github.com/dmedinaor11/sqlite-golang/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		panic(err)
	}
}
