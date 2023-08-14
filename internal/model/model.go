package model

import "time"

type List struct {
	ID       int
	Title    string
	ActiveAt time.Time
}
