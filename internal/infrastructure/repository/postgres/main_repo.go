package postgres

import (
	"gorm.io/gorm"
	"lms_system/internal/domain"
	"lms_system/internal/infrastructure/repository/postgres/chapter"
	"lms_system/internal/infrastructure/repository/postgres/course"
	"lms_system/internal/infrastructure/repository/postgres/lesson"
	"lms_system/internal/infrastructure/repository/postgres/user_course_access"
)

type MainRepository struct {
	db *gorm.DB
}

func NewMainRepository(db *gorm.DB) *MainRepository {
	return &MainRepository{db: db}
}

func (r *MainRepository) Course() domain.CourseRepositoryInterface {
	return course.NewRepository(r.db)
}

func (r *MainRepository) Chapter() domain.ChapterRepositoryInterface {
	return chapter.NewRepository(r.db)
}

func (r *MainRepository) Lesson() domain.LessonRepositoryInterface {
	return lesson.NewRepository(r.db)
}

func (r *MainRepository) UserCourseAccess() domain.UserCourseAccessInterface {
	return user_course_access.NewRepository(r.db)
}
