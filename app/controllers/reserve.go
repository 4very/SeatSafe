package controllers

import (
	"github.com/revel/revel"
)

type Reserve struct {
	*revel.Controller
}

func (c Reserve) Main(id string) revel.Result {
	return c.Render()
}
