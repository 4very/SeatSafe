package models

import (
	"fmt"
	"github.com/revel/revel"
)

type reservation struct {
	UID          int
	privateID    string
	email        string
	name         string
}

