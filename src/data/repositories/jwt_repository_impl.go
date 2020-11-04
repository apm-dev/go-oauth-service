package repositories

import (
	"github.com/apm-dev/go-oauth-service/src/core/errors"
	"github.com/apm-dev/go-oauth-service/src/domain/entities"
)

type JWTRepositoryImpl struct {
}

func (r *JWTRepositoryImpl) GetByAccessToken(at string) (*entities.JWT, *errors.RestError) {
	return nil, errors.InternalServerError("repository database not implemented")
}
