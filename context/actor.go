package context

import "time"

type Actor struct {
	IP         string
	Username   string
	UpdateTime time.Time
}
