package controllers

import (
	"SeatSafe/app/database"
	"SeatSafe/app/models"
	"strconv"

	"github.com/revel/revel"
	"github.com/twinj/uuid"
)

type ReservationManager struct {
	*revel.Controller
}

func (c ReservationManager) Main(id string) revel.Result {
	if id == "" {
		err := "Your need to choose an event to access!"
		return c.Render(err)
	}

	event, isFound := database.GetEvent(id)
	if !isFound {
		err := "Event not found, please try again"
		return c.Render(err)
	}

	SGJoin := database.GetSeatGroupData(event.EventId)

	return c.Render(event, SGJoin)
}

func (c ReservationManager) Cancel(id string) revel.Result {
	if id == "" {
		err := "Your need to choose a reservation to access!"
		return c.Render(err)
	}

	Res, isFound := database.GetResInfo(id)
	if !isFound {
		err := "Reservation not found, please try again"
		return c.Render(err)
	}

	ResData := database.GetResViewData(Res.ReservationId)

	return c.Render(Res, ResData)
}

func (c ReservationManager) Reserve(id string) revel.Result {
	formData := c.Params.Form
	event, isFound := database.GetEvent(formData.Get("eventId"))
	if !isFound {
		revel.AppLog.Error("Somehow tried to reserve an event that doesn't exist, stuff is broken")
		return c.Redirect("/", "Event not found, somehow.")
	}

	reservation := c.createReservationFromFormData(*event)
	reservation.ReservationId = database.InsertReservation(reservation)

	c.reserveSpotGroups(reservation.ReservationId)

	return c.Redirect("/")
}

func (c ReservationManager) createReservationFromFormData(event models.Event) models.Reservation {
	formData := c.Params.Form
	return models.Reservation{
		PrivateId: uuid.NewV4().String(),
		Email:     formData.Get("reserverEmail"),
		Name:      formData.Get("reserverName"),
		EventId:   event.EventId}
}

func (c ReservationManager) reserveSpotGroups(reservationId int64) {
	formData := c.Params.Form
	for i := 1; ; i++ {
		groupId, err := strconv.ParseInt(formData.Get("groupId"+strconv.Itoa(i)), 10, 64)
		if err != nil {
			break
		}
		numOfSeats, _ := strconv.ParseInt(formData.Get("seatsToReserveInGroup"+strconv.Itoa(i)), 10, 64)
		database.ReserveSpotsInSpotGroup(reservationId, groupId, numOfSeats)
	}
}

func (c ReservationManager) Delete(id string) revel.Result {
	res, _ := database.GetResInfo(id)
	database.DeleteReservation(*res)
	return c.Redirect("/")
}
