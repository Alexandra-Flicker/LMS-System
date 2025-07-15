package postgres

import (
	"lms_system/internal/domain"
	"lms_system/internal/infrastructure/repository/postgres/chapter"
	"lms_system/internal/infrastructure/repository/postgres/course"
	"lms_system/internal/infrastructure/repository/postgres/lesson"
)

type MainRepository struct {
}

func NewMainRepository() *MainRepository {
	return &MainRepository{}
}

func (r *MainRepository) Course() domain.CourseRepositoryInterface {
	return course.NewRepository()
}

func (r *MainRepository) Chapter() domain.ChapterRepositoryInterface {
	return chapter.NewRepository()
}

func (r *MainRepository) Lesson() domain.LessonRepositoryInterface {
	return lesson.NewRepository()
}
