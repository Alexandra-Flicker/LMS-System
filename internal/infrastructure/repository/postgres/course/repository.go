package course

import (
	"context"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"lms_system/internal/domain/entity"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateCourse(ctx context.Context, course entity.Course) (uint, error) {
	if err := r.db.WithContext(ctx).Create(&course).Error; err != nil {
		return 0, err
	}
	return course.ID, nil
}

func (r *Repository) UpdateCourseById(ctx context.Context, course entity.Course) error {
	err := r.db.WithContext(ctx).
		Model(&entity.Course{}).
		Where("id = ?", course.ID).
		Updates(course).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteCourseById(ctx context.Context, courseId uint) error {
	err := r.db.WithContext(ctx).
		Where("id = ?", courseId).
		Delete(&entity.Course{}).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetCourseById(ctx context.Context, id uint) (*entity.CourseAggregate, error) {
	var course entity.Course

	err := r.db.WithContext(ctx).
		Preload("Chapters").
		Where("id = ?", id).
		First(&course).Error

	if err != nil {
		return nil, err
	}

	return &entity.CourseAggregate{
		Course: course,
	}, nil
}

func (r *Repository) GetAllCourses(ctx context.Context) ([]entity.Course, error) {
	var courses []entity.Course

	err := r.db.WithContext(ctx).
		Find(&courses).Error

	if err != nil {
		return nil, err
	}

	return courses, nil
}
