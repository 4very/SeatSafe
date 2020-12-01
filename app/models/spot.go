package models

import (
	"fmt"
	"github.com/revel/revel"
)

type spot struct {
	UID            int
	spotGroupUID   int
	reservationUID int
}

