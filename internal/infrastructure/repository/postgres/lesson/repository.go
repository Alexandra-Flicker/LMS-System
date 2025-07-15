package lesson

import (
	"context"
	"lms_system/internal/domain/entity"
)

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) CreateLesson(ctx context.Context, chapterId uint, entity entity.Lesson) (id uint, err error) {
	return 0, nil
}

func (r *Repository) UpdateLessonById(ctx context.Context, lesson entity.Lesson) error {
	return nil
}

func (r *Repository) DeleteLessonById(ctx context.Context, lessonId uint) error {
	return nil
}

func (r *Repository) GetLessonById(ctx context.Context, id uint) (*entity.Lesson, error) {
	return nil, nil
}

func (r *Repository) GetAllLessonsByChapterId(ctx context.Context, chapterId uint) ([]entity.Lesson, error) {
	return nil, nil
}
