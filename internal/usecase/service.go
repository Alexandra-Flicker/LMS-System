package usecase

import (
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
