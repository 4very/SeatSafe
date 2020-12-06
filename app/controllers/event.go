package controllers

import (
	"SeatSafe/app/database"

	"github.com/revel/revel"
)

type Event struct {
	*revel.Controller
}

func (c Event) View(id string) revel.Result {

	// render with just an error
	if id == "" {
		err := "You need to choose an event to access!"
		return c.Render(err)
	}

	// query for event information
	event, isfound := database.GetEvent(id)

	// if the event id was not found
	if !isfound {
		err := "Event not found, please try again"
		return c.Render(err)
	}

	SGJoin := database.GetSeatGroupData(event.EventId) // Get seat group information
	ResJoin := database.GetResData(event.EventId)      // Get Reservation information

	// render either public or private page
	if id[0] == 'v' { // render private page
		isadmin := true
		return c.Render(event, SGJoin, ResJoin, isadmin)
	}
	if id[0] == 'b' { // render public page
		isadmin := false
		return c.Render(event, SGJoin, isadmin)
	}

	// it needs this here but we wont need it ¯\_(ツ)_/¯
	return c.Render(event)
}

func (c Event) Create(id string) revel.Result {
	return c.Render()
}

func (c Event) List() revel.Result {
	eventList := database.GetPublicEvents()
	return c.Render(eventList)
}
