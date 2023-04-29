package model

import (
	"time"
)

type Clinic struct {
	ID int
	Name string 
	CreatedAt time.Time
	UpdatedAt time.Time
}
