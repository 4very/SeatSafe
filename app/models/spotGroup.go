package models

import (
	"github.com/revel/revel"
)

type SpotGroup struct {
	SpotGroupId int
	EventId     int
	Name        string
}

func (sg SpotGroup) Validate(v *revel.Validation) {

}
