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
	UID:              1,
	PublicID:         "publicid1",
	PrivateID:        "privateid1",
	EventName:        "event example",
	ContactEmail:     "test@gmail.com",
	PublicallyListed: false,
	Image:            "example/path",
}

// spotgroup example hardcode //
var spotGroupExample = []*models.SpotGroup{
	{UID: 1, EventUID: 1, Name: "group1"},
	{UID: 2, EventUID: 1, Name: "group2"},
}

var spotExample = []*models.Spot{
	{UID: 1, SpotGroupUID: 1, ReservationUID: 1},
	{UID: 2, SpotGroupUID: 1, ReservationUID: 4},
	{UID: 3, SpotGroupUID: 2, ReservationUID: 2},
	{UID: 4, SpotGroupUID: 2, ReservationUID: 3},
}

var reservationExample = []*models.Reservation{
	{UID: 1, PrivateID: "privateid1", Email: "email1@gmail.com", Name: "Name1"},
	{UID: 2, PrivateID: "privateid2", Email: "email2@gmail.com", Name: "Name2"},
	{UID: 3, PrivateID: "privateid3", Email: "email3@gmail.com", Name: "Name3"},
	{UID: 4, PrivateID: "privateid4", Email: "email4@gmail.com", Name: "Name4"},
}

func (c Event) View() revel.Result {

	event := eventExample
	spotGroups := spotGroupExample

	return c.Render(event, spotGroups)
}

type TempJoin struct {
	Res          models.Reservation
	SpotUID      int
	SpotGroupUID int
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
			if spot.ReservationUID == reservation.UID {
				joins = append(joins, &TempJoin{Res: *reservation, SpotUID: spot.UID, SpotGroupUID: spot.SpotGroupUID})
			}
		}
	}

	return c.Render(event, spotGroups, joins)

}
