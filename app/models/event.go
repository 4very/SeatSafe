package models

import (
	"fmt"
	"github.com/revel/revel"
)

type event struct {
	UID                int
	PublicID           string
	PrivateID          string
	eventName          string
	contactEmail       string
	publicallyListed   bool
	image              string
}

