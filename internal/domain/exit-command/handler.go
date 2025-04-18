package exit

import (
	"context"
	"fmt"
	"os"
)

const Command = ".exit"

type Handler struct{}

func NewHandler() Handler {
	return Handler{}
}

func (h *Handler) Handle(c context.Context, statement string) error {
	fmt.Println("Goodbye!")
	os.Exit(0)
	return nil
}
