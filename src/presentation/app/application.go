package app

import (
	"github.com/apm-dev/go-oauth-service/src/data/datasources"
	"github.com/apm-dev/go-oauth-service/src/domain/repositories"
	"github.com/apm-dev/go-oauth-service/src/domain/usecases"
	"github.com/apm-dev/go-oauth-service/src/presentation/http"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func StartApplication() {
	//	Build and test cassandra connection
	session, err := datasources.GetCassandraClient().GetSession()
	if err != nil {
		panic(err)
	}
	session.Close()

	//	Handle dependency injections
	handler :=
		http.GetNewJWTHandler(
			usecases.GetNewJWTService(
				repositories.GetNewJWTRepository(
					datasources.GetCassandraClient(),
				),
			),
		)

	//	Endpoints
	router.GET("/oauth/jwt/:access_token", handler.GetByAccessToken)

	//	Run application
	router.Run(":8080")
}
