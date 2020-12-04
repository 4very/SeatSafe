package models

import (
	"github.com/revel/revel"
)

type Spot struct {
	SpotId        int64
	SpotGroupId   int64
	ReservationId int64
}

func (s Spot) Validate(v *revel.Validation) {

}
