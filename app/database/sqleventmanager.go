package database

import (
	"SeatSafe/app"
	"SeatSafe/app/models"
	"database/sql"
	"fmt"

	"github.com/revel/revel"
)

type SqlEventManager struct {
}

func InsertEvent(ev models.Event) int64 {
	sql := "INSERT INTO Event (PublicId, PrivateId, EventName, ContactEmail, PublicallyListed, ImageUrl) " +
		"VALUES (?, ?, ?, ?, ?, ?);"
	res, err := app.DB.Exec(sql, ev.PublicId, ev.PrivateId, ev.EventName, ev.ContactEmail, ev.PublicallyListed, ev.ImageUrl)
	if err != nil {
		revel.AppLog.Error("Error inserting an event into database", "error", err)
		return -1
	}

	lastId, err2 := res.LastInsertId()
	if err2 != nil {
		revel.AppLog.Error("Error reading auto_increment id from Event insert", "error", err2)
		return -1
	}
	return lastId
}

func InsertSpot(spot models.Spot) int64 {
	sql := "INSERT INTO Spot (SpotGroupId) " +
		"VALUES (?);"
	res, err := app.DB.Exec(sql, spot.SpotGroupId)
	if err != nil {
		revel.AppLog.Error("Error inserting a Spot into database", "error", err)
		return -1
	}

	lastId, err2 := res.LastInsertId()
	if err2 != nil {
		revel.AppLog.Error("Error reading auto_increment id from Spot insert", "error", err2)
	}
	return lastId
}

func InsertSpotGroup(spotGroup models.SpotGroup) int64 {
	sql := "INSERT INTO SpotGroup (EventId, Name) " +
		"VALUES (?, ?);"
	res, err := app.DB.Exec(sql, spotGroup.EventId, spotGroup.Name)
	if err != nil {
		revel.AppLog.Error("Error inserting a SpotGroup into database", "error", err)
		return -1
	}

	lastId, err2 := res.LastInsertId()
	if err2 != nil {
		revel.AppLog.Error("Error reading auto_increment id from SpotGroup insert", "error", err2)
	}
	return lastId
}

func ReserveSpot(reservation models.Reservation, spot models.Spot) {
	sql := "UPDATE Spot SET ReservationId = ? " +
		"SpotId = ?;"
	_, err := app.DB.Exec(sql, reservation.ReservationId, spot.SpotId)
	if err != nil {
		revel.AppLog.Error("Error reserving a Spot in database", "error", err)
		return
	}
}

func DeleteEvent(event models.Event) {
	sql := "DELETE FROM Event WHERE " +
		"EventId = ?;"
	_, err := app.DB.Exec(sql, event.EventId)
	if err != nil {
		revel.AppLog.Error("Error deleting an Event in database", "error", err)
		return
	}
}

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
									FROM seatsafe.event e, seatsafe.spotgroup sg, seatsafe.spot s
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
								  FROM seatsafe.spotgroup sg, seatsafe.spot s
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
								FROM seatsafe.reservation r, seatsafe.spotgroup sg, seatsafe.spot s
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
