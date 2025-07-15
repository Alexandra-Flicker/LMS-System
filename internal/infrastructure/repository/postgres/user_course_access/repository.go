package user_course_access

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

func (r *Repository) GetAllByUserId(ctx context.Context, userId uint) ([]entity.UserCourseAccess, error) {
	var accessList []entity.UserCourseAccess

	err := r.db.WithContext(ctx).
		Where("user_id = ?", userId).
		Find(&accessList).Error

	if err != nil {
		return nil, err
	}

	return accessList, nil
}

func (r *Repository) UpdateAccess(ctx context.Context, access *entity.UserCourseAccess) error {
	err := r.db.WithContext(ctx).
		Model(&entity.UserCourseAccess{}).
		Where("user_id = ? AND course_id = ?", access.UserID, access.CourseID).
		Updates(access).Error

	if err != nil {
		return err
	}

	return nil
}
