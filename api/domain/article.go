package domain

import "time"

type Article struct {
	ID       int64
	Subject  string
	Body     string
	Tags     Tags
	Modified time.Time
	Created  time.Time
}
