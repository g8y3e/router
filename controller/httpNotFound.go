package controller

import (
	"github.com/g8y3e/router/entity"
)

type HttpNotFound struct {
}

func (h *HttpNotFound) Process(ctx *entity.Context) error {
	return nil
}
