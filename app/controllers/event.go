package controllers

import (
	"SeatSafe/app"
	"SeatSafe/app/models"
	"database/sql"
	"fmt"

	"github.com/revel/revel"
)

type Event struct {
	*revel.Controller
}

type ResJoinStruct struct {
	ResName       string
	ResEmail      string
	SpotsRes      int64
	SpotGroupName string
}

type SGJoinStruct struct {
	SGName      string
	NumSpotsRem int64
	NumSpots    int64
}

func (c Event) View(id string) revel.Result {

	// render with just an error
	if id == "" {
		err := "You need to choose an event to access!"
		return c.Render(err)
	}

	// query for event information
	res := app.DB.QueryRow("SELECT * FROM Event WHERE PrivateId=? OR PublicId=?;", id, id)

	var event *models.Event = &models.Event{}
	sqlErr := res.Scan(&event.EventId, &event.PublicId, &event.PrivateId, &event.EventName, &event.ContactEmail, &event.PublicallyListed, &event.ImageUrl)

	if sqlErr == sql.ErrNoRows { // if the event is not found in the database
		err := "Event not found, please try again"
		return c.Render(err)
	}

	// Get seat group information
	var SGJoin []*SGJoinStruct
	var SGTemp *SGJoinStruct
	SGquery, err := app.DB.Query(`SELECT sg.Name, count(*), count(CASE WHEN s.ReservationId is NULL THEN 1 END)
								  FROM seatsafe.spotgroup sg, seatsafe.spot s
							      WHERE sg.SpotGroupId=s.SpotGroupId AND sg.EventId=?
							      GROUP BY sg.Name`, event.EventId)

	if err != nil {
		fmt.Println(err)
	}
	for SGquery.Next() {
		SGTemp = &SGJoinStruct{}
		SGquery.Scan(&SGTemp.SGName, &SGTemp.NumSpots, &SGTemp.NumSpotsRem)
		SGJoin = append(SGJoin, SGTemp)
	}

	// Get Reservation information
	var ResJoin []*ResJoinStruct
	var ResTemp *ResJoinStruct
	ResQuery, err := app.DB.Query(`SELECT r.Name, r.Email, sg.Name, count(s.SpotId)
								FROM seatsafe.reservation r, seatsafe.spotgroup sg, seatsafe.spot s
								WHERE r.ReservationId = s.ReservationId AND s.SpotGroupId = sg.SpotGroupId AND r.EventId=?
								GROUP BY r.Name, r.Email, sg.Name`, event.EventId)

	if err != nil {
		fmt.Println(err)
	}
	for ResQuery.Next() {
		ResTemp = &ResJoinStruct{}
		ResQuery.Scan(&ResTemp.ResName, &ResTemp.ResEmail, &ResTemp.SpotGroupName, &ResTemp.SpotsRes)
		ResJoin = append(ResJoin, ResTemp)
	}

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

type EventList struct {
	EventName     string
	EventEmail    string
	EventPublicId string
	TotSeats      int64
	AvalSeats     int64
}

func (c Event) List() revel.Result {

	var eventList []*EventList
	var eventTemp *EventList
	SGquery, err := app.DB.Query(`Select e.EventName, e.ContactEmail, e.PublicId, count(*), count(CASE WHEN s.ReservationId is NULL THEN 1 END)
									FROM seatsafe.event e, seatsafe.spotgroup sg, seatsafe.spot s
									WHERE e.EventId = sg.EventId AND sg.SpotGroupId = s.SpotGroupId AND e.PublicallyListed = 1
									GROUP BY e.EventName, e.ContactEmail, e.PublicId`)

	if err != nil {
		fmt.Println(err)
	}
	for SGquery.Next() {
		eventTemp = &EventList{}
		SGquery.Scan(&eventTemp.EventName, &eventTemp.EventEmail, &eventTemp.EventPublicId, &eventTemp.TotSeats, &eventTemp.AvalSeats)
		eventList = append(eventList, eventTemp)
	}

	return c.Render(eventList)
}
