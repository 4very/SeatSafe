package models

import (
	"github.com/revel/revel"
)

type Reservation struct {
	ReservationId int
	PrivateId     string
	Email         string
	Name          string
}

func (r Reservation) Validate(v *revel.Validation) {

}
