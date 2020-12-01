package controllers

import (
	"github.com/revel/revel"
)

type Event struct {
	*revel.Controller
}

type spotGroup struct {
	Name string
	Num int
}

func (c Event) View() revel.Result {
	eventName := "Test"
	email := "test@gmail.com"
	var spotGroups  = []spotGroup {spotGroup{"group1",2}, spotGroup{"group2",6},}
	return c.Render(eventName, email, spotGroups)
}