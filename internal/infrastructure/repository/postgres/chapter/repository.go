package chapter

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

func (r *Repository) CreateChapter(ctx context.Context, courseId uint, chapter *entity.Chapter) (uint, error) {
	chapter.CourseID = courseId

	err := r.db.WithContext(ctx).Create(chapter).Error
	if err != nil {
		return 0, err
	}

	return chapter.ID, nil
}

func (r *Repository) UpdateChapterById(ctx context.Context, chapter *entity.Chapter) error {
	err := r.db.WithContext(ctx).
		Model(&entity.Chapter{}).
		Where("id = ?", chapter.ID).
		Updates(chapter).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteChapterById(ctx context.Context, chapterId uint) error {
	err := r.db.WithContext(ctx).
		Where("id = ?", chapterId).
		Delete(&entity.Chapter{}).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetChapterById(ctx context.Context, id uint) (*entity.Chapter, error) {
	var chapter entity.Chapter

	err := r.db.WithContext(ctx).
		Where("id = ?", id).
		First(&chapter).Error

	if err != nil {
		return nil, err
	}

	return &chapter, nil
}

func (r *Repository) GetChaptersByCourseId(ctx context.Context, id uint) ([]entity.Chapter, error) {
	var chapters []entity.Chapter

	err := r.db.WithContext(ctx).
		Where("course_id = ?", id).
		Find(&chapters).Error

	if err != nil {
		return nil, err
	}

	return chapters, nil
}
