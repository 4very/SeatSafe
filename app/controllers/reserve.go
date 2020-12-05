package controllers

import (
	"github.com/revel/revel"
)

type Reserve struct {
	*revel.Controller
}

func (c Reserve) Main(eventId string) revel.Result {
	return c.Render()
}

func (c Reserve) Cancel(id string) revel.Result {
	return c.Render()
}
