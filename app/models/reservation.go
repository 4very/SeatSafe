package models

import (
	"github.com/revel/revel"
)

type Reservation struct {
	UID       int
	PrivateID string
	Email     string
	Name      string
}

func (r Reservation) Validate(v *revel.Validation) {

}
