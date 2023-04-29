package model

import (
	"time"
)

type Model struct {
	ID int 
	Name string
	CreatedAt time.Time
	UpdatedAt time.Time
}
