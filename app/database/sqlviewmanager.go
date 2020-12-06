package database

import (
	"SeatSafe/app"
	"SeatSafe/app/models"
	"database/sql"
	"fmt"
)

func GetEvent(id string) (*models.Event, bool) {
	res := app.DB.QueryRow("SELECT * FROM Event WHERE PrivateId=? OR PublicId=?;", id, id)

	var event *models.Event = &models.Event{}
	sqlErr := res.Scan(&event.EventId, &event.PublicId, &event.PrivateId, &event.EventName, &event.ContactEmail, &event.PublicallyListed, &event.ImageUrl)

	var isfound bool = sqlErr != sql.ErrNoRows
	return event, isfound
}

type EventLineData struct {
	EventName     string
	EventEmail    string
	EventPublicId string
	TotSeats      int64
	AvalSeats     int64
}

func GetPublicEvents() []*EventLineData {
	var eventList []*EventLineData
	var eventTemp *EventLineData
	SGquery, err := app.DB.Query(`Select e.EventName, e.ContactEmail, e.PublicId, count(*), count(CASE WHEN s.ReservationId is NULL THEN 1 END)
									FROM SeatSafe.Event e, SeatSafe.SpotGroup sg, SeatSafe.Spot s
									WHERE e.EventId = sg.EventId AND sg.SpotGroupId = s.SpotGroupId AND e.PublicallyListed = 1
									GROUP BY e.EventName, e.ContactEmail, e.PublicId`)

	if err != nil {
		fmt.Println(err)
	}
	for SGquery.Next() {
		eventTemp = &EventLineData{}
		SGquery.Scan(&eventTemp.EventName, &eventTemp.EventEmail, &eventTemp.EventPublicId, &eventTemp.TotSeats, &eventTemp.AvalSeats)
		eventList = append(eventList, eventTemp)
	}

	return eventList

}

type SGLineData struct {
	SGName      string
	NumSpotsRem int64
	NumSpots    int64
}

func GetSeatGroupData(eventId int64) []*SGLineData {

	var SGData []*SGLineData
	var SGTemp *SGLineData
	SGquery, err := app.DB.Query(`SELECT sg.Name, count(*), count(CASE WHEN s.ReservationId is NULL THEN 1 END)
								  FROM SeatSafe.SpotGroup sg, SeatSafe.Spot s
							      WHERE sg.SpotGroupId=s.SpotGroupId AND sg.EventId=?
							      GROUP BY sg.Name`, eventId)

	if err != nil {
		fmt.Println(err)
	}
	for SGquery.Next() {
		SGTemp = &SGLineData{}
		SGquery.Scan(&SGTemp.SGName, &SGTemp.NumSpots, &SGTemp.NumSpotsRem)
		SGData = append(SGData, SGTemp)
	}
	return SGData
}

type ResLineData struct {
	ResName       string
	ResEmail      string
	SpotsRes      int64
	SpotGroupName string
}

func GetResData(eventId int64) []*ResLineData {

	var ResJoin []*ResLineData
	var ResTemp *ResLineData
	ResQuery, err := app.DB.Query(`SELECT r.Name, r.Email, sg.Name, count(s.SpotId)
								FROM SeatSafe.Reservation r, SeatSafe.SpotGroup sg, SeatSafe.Spot s
								WHERE r.ReservationId = s.ReservationId AND s.SpotGroupId = sg.SpotGroupId AND r.EventId=?
								GROUP BY r.Name, r.Email, sg.Name`, eventId)

	if err != nil {
		fmt.Println(err)
	}
	for ResQuery.Next() {
		ResTemp = &ResLineData{}
		ResQuery.Scan(&ResTemp.ResName, &ResTemp.ResEmail, &ResTemp.SpotGroupName, &ResTemp.SpotsRes)
		ResJoin = append(ResJoin, ResTemp)
	}

	return ResJoin
}
