package factories

import (
	"SeatSafe/app/database"
	"SeatSafe/app/models"
)

func CreateSpots(spotGroupId int64, numOfSpots int) {
	for i := 0; i < numOfSpots; i++ {
		spot := models.Spot{
			SpotGroupId: spotGroupId}
		database.InsertSpot(spot)
	}
}
