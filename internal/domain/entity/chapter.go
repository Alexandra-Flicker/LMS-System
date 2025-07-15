package entity

import "time"

type ChapterAggregate struct {
	Chapter
	Lessons []Lesson
}
type Chapter struct {
	ID            uint
	Name          string
	Description   string
	OrderPosition int
	CourseID      uint
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
