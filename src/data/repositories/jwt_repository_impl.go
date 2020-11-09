package repositories

import (
	"github.com/apm-dev/go-oauth-service/src/core/errors"
	"github.com/apm-dev/go-oauth-service/src/data/datasources"
	"github.com/apm-dev/go-oauth-service/src/domain/entities"
)

type JWTRepositoryImpl struct {
	Cassandra datasources.CassandraDB
}

func (r *JWTRepositoryImpl) GetByAccessToken(at string) (*entities.JWT, *errors.RestError) {
	session, err := r.Cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	return nil, errors.InternalServerError("repository database not implemented")
}
