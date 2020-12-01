package models

import "github.com/revel/revel"

type Event struct {
	UID              int
	PublicID         string
	PrivateID        string
	EventName        string
	ContactEmail     string
	PublicallyListed bool
	Image            string
}

func (e Event) Validate(v *revel.Validation) {

}
