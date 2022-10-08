package domain

import "time"

type Tag struct {
	ID       int64
	Name     string
	Modified time.Time
	Created  time.Time
}

type Tags []*Tag
