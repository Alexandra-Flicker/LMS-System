package entity

import "time"

type Lesson struct {
	ID            uint
	Name          string
	Description   string
	Content       string
	OrderPosition int
	ChapterID     uint
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
