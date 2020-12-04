package controllers

import "github.com/revel/revel"

type EventManager struct {
	*revel.Controller
}

func CreateEvent() int {
	return 1
}
