package entity

import (
	"time"
)

type Payment struct {
	id                string
	method            string
	installments      int
	tax               int
	totalValueInCents int
	status            string
	createdAt         time.Time
}
