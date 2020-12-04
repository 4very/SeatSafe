package models

import (
	"github.com/revel/revel"
)

type SpotGroup struct {
	SpotGroupId int64
	EventId     int64
	Name        string
}

func (sg SpotGroup) Validate(v *revel.Validation) {

}
