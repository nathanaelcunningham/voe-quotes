package models

import "time"

type Quote struct {
	ID      int
	Quote   string
	Author  string
	Created time.Time
}
