package models

import (
	"github.com/revel/revel"
)

type Spot struct {
	UID            int
	SpotGroupUID   int
	ReservationUID int
}

func (s Spot) Validate(v *revel.Validation) {

}
