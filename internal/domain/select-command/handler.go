package selectCommand

import (
	"context"
	"fmt"
)

const Command = "select"

type Handler struct{}

func NewHandler() Handler {
	return Handler{}
}

func (h *Handler) Handle(c context.Context, statement string) error {
	fmt.Println("Doing select...")
	return nil
}
