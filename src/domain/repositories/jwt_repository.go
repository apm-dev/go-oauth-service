package repositories

import (
	"github.com/apm-dev/go-oauth-service/src/core/errors"
	"github.com/apm-dev/go-oauth-service/src/data/repositories"
	"github.com/apm-dev/go-oauth-service/src/domain/entities"
)

type JWTRepository interface {
	GetByAccessToken(at string) (*entities.JWT, *errors.RestError)
}

func GetNewJWTRepository() JWTRepository {
	return &repositories.JWTRepositoryImpl{}
}
