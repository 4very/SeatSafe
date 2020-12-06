package controllers

import (
	"SeatSafe/app/database"
	"SeatSafe/app/factories"
	"SeatSafe/app/models"
	"strconv"

	"github.com/revel/revel"
	"github.com/twinj/uuid"
)

type EventManager struct {
	*revel.Controller
}

func (e EventManager) CreateEvent() revel.Result {
	event := e.createEventFromFormData()
	event.EventId = database.InsertEvent(event)
	e.createSpotGroups(event.EventId)
	return e.Redirect("/event?id=" + event.PrivateId)
}

func (e EventManager) createEventFromFormData() models.Event {
	formData := e.Params.Form
	var publicallyListed bool
	if formData.Get("eventPrivacy") == "12" {
		publicallyListed = true
	} else {
		publicallyListed = false
	}
	event := models.Event{
		PublicId:         "b" + uuid.NewV4().String(),
		PrivateId:        "v" + uuid.NewV4().String(),
		EventName:        formData.Get("eventName"),
		ContactEmail:     formData.Get("contactEmail"),
		PublicallyListed: publicallyListed,
		ImageUrl:         formData.Get("imageUrl"),
	}
	revel.AppLog.Error("Debugging event info", "event", event)
	return event
}

func (e EventManager) createSpotGroups(eventId int64) {
	formData := e.Params.Form
	var spotGroups []factories.SpotGroupConfig
	for i := 0; ; i++ {
		groupName := formData.Get("groupName" + strconv.Itoa(i))
		if groupName == "Seat Group Name" {
			continue
		}
		if groupName == "" {
			break
		}
		numOfSeats, _ := strconv.ParseInt(formData.Get("groupSeatCount"+strconv.Itoa(i)), 10, 64)
		spotGroups = append(spotGroups, factories.SpotGroupConfig{
			Name:       groupName,
			NumOfSpots: int(numOfSeats)})
	}
	factories.CreateSpotGroups(spotGroups, eventId)
}
