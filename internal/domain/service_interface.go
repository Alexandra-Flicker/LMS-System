package domain

import (
	"context"
	"lms_system/internal/domain/dto"
	"lms_system/internal/domain/entity"
)

type ServiceInterface interface {
	UserServiceInterface
	AdminServiceInterface
}

type UserServiceInterface interface {
	BuyCourse(ctx context.Context, request dto.BuyCourseRequest) error
	GetAllCourses(ctx context.Context) ([]entity.Course, error)
	GetCourseById(ctx context.Context, id uint) (entity.CourseAggregate, error)
	GetChaptersInfoByCourseId(ctx context.Context, id uint) ([]entity.ChapterInfoAggregate, error)
	GetLessonById(ctx context.Context, id uint) (entity.Lesson, error)
}

type AdminServiceInterface interface {
	CreateCourse(ctx context.Context, entity entity.Course) (id uint, err error)
	CreateChapter(ctx context.Context, courseId uint, entity entity.Chapter) (id uint, err error)
	CreateLesson(ctx context.Context, chapterId uint, entity entity.Lesson) (id uint, err error)

	UpdateCourseById(ctx context.Context, course entity.Course) error
	UpdateChapterById(ctx context.Context, chapter entity.Chapter) error
	UpdateLessonById(ctx context.Context, lesson entity.Lesson) error

	DeleteCourseById(ctx context.Context, courseId uint) error
	DeleteChapterById(ctx context.Context, chapterId uint) error
	DeleteLessonById(ctx context.Context, lessonId uint) error
}
