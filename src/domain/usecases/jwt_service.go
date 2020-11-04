package usecases

import (
	"github.com/apm-dev/go-oauth-service/src/core/errors"
	"github.com/apm-dev/go-oauth-service/src/domain/entities"
	"github.com/apm-dev/go-oauth-service/src/domain/repositories"
)

type JWTService interface {
	GetByAccessToken(at string) (*entities.JWT, *errors.RestError)
}

func GetNewJWTService(r repositories.JWTRepository) JWTService {
	return &jwtService{repo: r}
}

type jwtService struct {
	repo repositories.JWTRepository
}

func (s *jwtService) GetByAccessToken(at string) (*entities.JWT, *errors.RestError) {
	return s.repo.GetByAccessToken(at)
}
