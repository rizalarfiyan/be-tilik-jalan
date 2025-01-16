package service

import (
	"context"

	"github.com/rizalarfiyan/be-tilik-jalan/exception"
	"github.com/rizalarfiyan/be-tilik-jalan/internal/model"
	"github.com/rizalarfiyan/be-tilik-jalan/internal/repository"
)

type CCTVService interface {
	GetAll(ctx context.Context) model.CCTVs
}

type cctvService struct {
	exception exception.Exception
	repo      repository.CCTVRepository
}

func NewCCTVService(repo repository.CCTVRepository) CCTVService {
	return &cctvService{
		exception: exception.NewException(),
		repo:      repo,
	}
}

func (s *cctvService) GetAll(ctx context.Context) model.CCTVs {
	list, err := s.repo.GetAll(ctx)
	s.exception.ErrorSkipNotFound(err)
	return list
}
