package database

import (
	"SeatSafe/app"
	"SeatSafe/app/models"

	"github.com/revel/revel"
)

func InsertReservation(reservation models.Reservation) int64 {
	sql := "INSERT INTO Reservation (PrivateId, Email, Name, EventId) " +
		"VALUES (?, ?, ?, ?);"
	res, err := app.DB.Exec(sql, reservation.PrivateId, reservation.Email, reservation.Name, reservation.EventId)

	if err != nil {
		revel.AppLog.Error("Error reserving spots in database", "error", err)
		return -1
	}

	lastId, err2 := res.LastInsertId()
	if err2 != nil {
		revel.AppLog.Error("Error reading auto_increment id from Event insert", "error", err2)
		return -1
	}
	return lastId
}

func ReserveSpotsInSpotGroup(reservationId int64, spotGroupId int64, numOfSpots int64) {
	sql := "UPDATE Spot SET ReservationId = ? " +
		"WHERE SpotGroupId=? LIMIT ?"
	_, err := app.DB.Exec(sql, reservationId, spotGroupId, numOfSpots)
	if err != nil {
		revel.AppLog.Error("Error reserving spots in database", "error", err)
	}
}

func DeleteReservation(res models.Reservation) {

	_, err := app.DB.Exec("DELETE FROM Reservation WHERE ReservationId = ?;", res.ReservationId)
	if err != nil {
		revel.AppLog.Error("Error deleting an Reservation in database", "error", err)
		return
	}
	_, err = app.DB.Exec("UPDATE Spot SET ReservationId = Null WHERE ReservationId = ?;", res.ReservationId)
	if err != nil {
		revel.AppLog.Error("Error deleting an Reservation in database", "error", err)
		return
	}
}
