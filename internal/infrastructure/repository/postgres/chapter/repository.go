package chapter

import (
	"context"
	"lms_system/internal/domain/entity"
)

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) CreateChapter(ctx context.Context, courseId uint, entity *entity.Chapter) (id uint, err error) {
	return 0, nil
}

func (r *Repository) UpdateChapterById(ctx context.Context, chapter *entity.Chapter) error {
	return nil
}

func (r *Repository) DeleteChapterById(ctx context.Context, chapterId uint) error {
	return nil
}

func (r *Repository) GetChapterById(ctx context.Context, id uint) (*entity.Chapter, error) {
	return nil, nil
}

func (r *Repository) GetChaptersByCourseId(ctx context.Context, id uint) ([]entity.Chapter, error) {
	return nil, nil
}
