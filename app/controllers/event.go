package controllers

import (
	"SeatSafe/app/models"

	"github.com/revel/revel"
)

type Event struct {
	*revel.Controller
}

// event example hardcode //
var eventExample = &models.Event{
	EventId:          1,
	PublicId:         "publicid1",
	PrivateId:        "privateid1",
	EventName:        "event example",
	ContactEmail:     "test@gmail.com",
	PublicallyListed: false,
	ImageUrl:         "example/path",
}

// spotgroup example hardcode //
var spotGroupExample = []*models.SpotGroup{
	{SpotGroupId: 1, EventId: 1, Name: "group1"},
	{SpotGroupId: 2, EventId: 1, Name: "group2"},
}

var spotExample = []*models.Spot{
	{SpotId: 1, SpotGroupId: 1, ReservationId: 1},
	{SpotId: 2, SpotGroupId: 1, ReservationId: 4},
	{SpotId: 3, SpotGroupId: 2, ReservationId: 2},
	{SpotId: 4, SpotGroupId: 2, ReservationId: 3},
}

var reservationExample = []*models.Reservation{
	{ReservationId: 1, PrivateId: "privateid1", Email: "email1@gmail.com", Name: "Name1"},
	{ReservationId: 2, PrivateId: "privateid2", Email: "email2@gmail.com", Name: "Name2"},
	{ReservationId: 3, PrivateId: "privateid3", Email: "email3@gmail.com", Name: "Name3"},
	{ReservationId: 4, PrivateId: "privateid4", Email: "email4@gmail.com", Name: "Name4"},
}

func (c Event) View() revel.Result {

	event := eventExample
	spotGroups := spotGroupExample

	return c.Render(event, spotGroups)
}

type TempJoin struct {
	Res          models.Reservation
	SpotUID      int64
	SpotGroupUID int64
}

func (c Event) Admin() revel.Result {

	event := eventExample
	spotGroups := spotGroupExample
	spots := spotExample
	reservations := reservationExample

	// joining spots and reservations //
	// prob temp //
	var joins []*TempJoin
	for _, spot := range spots {
		for _, reservation := range reservations {
			if spot.ReservationId == reservation.ReservationId {
				joins = append(joins, &TempJoin{Res: *reservation, SpotUID: spot.SpotId, SpotGroupUID: spot.SpotGroupId})
			}
		}
	}

	return c.Render(event, spotGroups, joins)

}
