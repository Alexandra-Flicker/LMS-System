package lesson

import (
	"context"
	"gorm.io/gorm"
	"lms_system/internal/domain/entity"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateLesson(ctx context.Context, chapterId uint, lesson entity.Lesson) (uint, error) {
	lesson.ChapterID = chapterId

	err := r.db.WithContext(ctx).Create(&lesson).Error
	if err != nil {
		return 0, err
	}

	return lesson.ID, nil
}

func (r *Repository) UpdateLessonById(ctx context.Context, lesson entity.Lesson) error {
	err := r.db.WithContext(ctx).
		Model(&entity.Lesson{}).
		Where("id = ?", lesson.ID).
		Updates(lesson).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteLessonById(ctx context.Context, lessonId uint) error {
	err := r.db.WithContext(ctx).
		Where("id = ?", lessonId).
		Delete(&entity.Lesson{}).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetLessonById(ctx context.Context, id uint) (*entity.Lesson, error) {
	var lesson entity.Lesson

	err := r.db.WithContext(ctx).
		Where("id = ?", id).
		First(&lesson).Error

	if err != nil {
		return nil, err
	}

	return &lesson, nil
}

func (r *Repository) GetAllLessonsByChapterId(ctx context.Context, chapterId uint) ([]entity.Lesson, error) {
	var lessons []entity.Lesson

	err := r.db.WithContext(ctx).
		Where("chapter_id = ?", chapterId).
		Find(&lessons).Error

	if err != nil {
		return nil, err
	}

	return lessons, nil
}
