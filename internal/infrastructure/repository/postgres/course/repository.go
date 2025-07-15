package course

import (
	"context"
	"lms_system/internal/domain/entity"
)

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) CreateCourse(ctx context.Context, entity entity.Course) (id uint, err error) {
	return 0, nil
}

func (r *Repository) UpdateCourseById(ctx context.Context, course entity.Course) error {
	return nil
}

func (r *Repository) DeleteCourseById(ctx context.Context, courseId uint) error {
	return nil
}

func (r *Repository) GetCourseById(ctx context.Context, id uint) (*entity.CourseAggregate, error) {
	return nil, nil
}

func (r *Repository) GetAllCourses(ctx context.Context) ([]entity.Course, error) {
	return nil, nil
}
