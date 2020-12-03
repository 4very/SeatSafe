package models

import (
	"github.com/revel/revel"
)

type Spot struct {
	SpotId        int
	SpotGroupId   int
	ReservationId int
}

func (s Spot) Validate(v *revel.Validation) {

}
