package usecase

import (
	"context"
	"errors"
	"lms_system/internal/domain"
)

type Service struct {
	mainRepo domain.MainRepositoryInterface
}

func NewService(mainRepo domain.MainRepositoryInterface) *Service {
	return &Service{
		mainRepo: mainRepo,
	}
}

func (s *Service) Update(ctx context.Context, courseID uint) error {
	// Validate
	if courseID == 0 {
		return errors.New("course id is required")
	}

	s.mainRepo.Course().CreateCourse(ctx, 33)
}
