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
	ImageUrl:         "https://assets.change.org/photos/7/qj/gy/GtQJGyGFioacDMA-400x225-noPad.jpg?1524796577",
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
	{SpotId: 4, SpotGroupId: 2, ReservationId: 0},
}

var reservationExample = []*models.Reservation{
	{ReservationId: 1, PrivateId: "privateid1", Email: "email1@gmail.com", Name: "Name1"},
	{ReservationId: 2, PrivateId: "privateid2", Email: "email2@gmail.com", Name: "Name2"},
	{ReservationId: 3, PrivateId: "privateid3", Email: "email3@gmail.com", Name: "Name3"},
	{ReservationId: 4, PrivateId: "privateid4", Email: "email4@gmail.com", Name: "Name4"},
}

type TempResJoin struct {
	Res           models.Reservation
	SpotsRes      int64
	SpotGroupName string
}

type TempSGJoin struct {
	SG          models.SpotGroup
	NumSpotsRem int64
	NumSpots    int64
}

func (c Event) View(id string) revel.Result {

	// render with just an error
	if id == "" {
		err := "You need to choose an event to access!"
		return c.Render(err)
	}

	event := eventExample
	spotGroups := spotGroupExample
	spots := spotExample
	reservations := reservationExample

	//// TOADD: QUERYING OF THE DATABASE ////

	// joining spots and reservations //
	// prob temp //
	var SGJoin []*TempSGJoin
	var temp *TempSGJoin
	for _, SG := range spotGroups {
		temp = &TempSGJoin{SG: *SG, NumSpotsRem: 0, NumSpots: 0}
		for _, spot := range spots {
			if spot.SpotGroupId == SG.SpotGroupId {
				temp.NumSpots++
				if spot.ReservationId == 0 {
					temp.NumSpotsRem++
				}
			}
		}
		SGJoin = append(SGJoin, temp)
	}

	var ResJoin []*TempResJoin
	var temp2 *TempResJoin

	for _, res := range reservations {
		temp2 = &TempResJoin{Res: *res, SpotsRes: 0, SpotGroupName: ""}

		for _, spot := range spots {
			if spot.ReservationId == res.ReservationId {
				temp2.SpotsRes++
				for _, sg := range spotGroups {
					if sg.SpotGroupId == spot.SpotGroupId {
						temp2.SpotGroupName = sg.Name
					}
				}
			}
		}
		ResJoin = append(ResJoin, temp2)
	}

	notFound := false // temp this will be set automatically

	// if it's not found in the database
	if notFound {
		err := "Event not found, please try again"
		return c.Render(err)
	}

	// render either public or private page
	if id[0] == 'v' { // render private page
		isadmin := true
		return c.Render(event, spotGroups, SGJoin, ResJoin, isadmin)
	}
	if id[0] == 'b' { // render public page
		isadmin := false
		return c.Render(event, spotGroups, SGJoin, isadmin)
	}

	// it needs this here but we wont need it ¯\_(ツ)_/¯
	return c.Render(event, spotGroups)
}

func (c Event) Create(id string) revel.Result {
	return c.Render()
}

func (c Event) List() revel.Result {
	return c.Render()
}
