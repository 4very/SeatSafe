package models

import (
	"github.com/revel/revel"
)

type SpotGroup struct {
	UID      int
	EventUID int
	Name     string
}

func (sg SpotGroup) Validate(v *revel.Validation) {

}
