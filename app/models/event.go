package models

import "github.com/revel/revel"

type Event struct {
	EventId          int64
	PublicId         string
	PrivateId        string
	EventName        string
	ContactEmail     string
	PublicallyListed bool
	ImageUrl         string
}

func (e Event) Validate(v *revel.Validation) {

}
