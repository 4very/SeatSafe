package database

import (
	"SeatSafe/app"
	"SeatSafe/app/models"

	"github.com/revel/revel"
)

type SqlEventManager struct {
}

func InsertEvent(ev models.Event) int64 {
	sql := "INSERT INTO Event (PublicId, PrivateId, EventName, ContactEmail, PublicallyListed, ImageUrl) " +
		"VALUES ($1, $2, $3, $4, $5, $6);"
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
	sql := "INSERT INTO Spot (SpotId, SpotGroupId, ReservationId) " +
		"VALUES ($1, $2, $3);"
	res, err := app.DB.Exec(sql, spot.SpotId, spot.SpotGroupId, spot.ReservationId)
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
		"VALUES ($1, $2);"
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
	sql := "UPDATE Spot SET ReservationId = $1 " +
		"SpotId = $2;"
	_, err := app.DB.Exec(sql, reservation.ReservationId, spot.SpotId)
	if err != nil {
		revel.AppLog.Error("Error reserving a Spot in database", "error", err)
		return
	}
}

func DeleteEvent(event models.Event) {
	sql := "DELETE FROM Event WHERE " +
		"EventId = $1;"
	_, err := app.DB.Exec(sql, event.EventId)
	if err != nil {
		revel.AppLog.Error("Error deleting an Event in database", "error", err)
		return
	}
}
