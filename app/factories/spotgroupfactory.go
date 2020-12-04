package factories

import (
	"SeatSafe/app/database"
	"SeatSafe/app/models"
)

type SpotGroupConfig struct {
	name       string
	numOfSpots int
}

func CreateSpotGroups(spotGroups []SpotGroupConfig, eventId int64) {
	for _, spotGroupConfig := range spotGroups {
		spotGroup := models.SpotGroup{
			EventId: eventId,
			Name:    spotGroupConfig.name}
		spotGroup.SpotGroupId = database.InsertSpotGroup(spotGroup)
		CreateSpots(spotGroup.SpotGroupId, spotGroupConfig.numOfSpots)
	}
}
