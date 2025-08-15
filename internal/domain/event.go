package domain

import "time"

type ClickEvent struct {
	ID        int64
	Alias     string
	Timestamp time.Time
	Source    string
}
