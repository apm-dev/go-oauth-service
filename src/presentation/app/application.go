package app

import (
	"github.com/apm-dev/go-oauth-service/src/domain/repositories"
	"github.com/apm-dev/go-oauth-service/src/domain/usecases"
	"github.com/apm-dev/go-oauth-service/src/presentation/http"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func StartApplication() {
	handler := http.GetNewJWTHandler(usecases.GetNewJWTService(repositories.GetNewJWTRepository()))

	router.GET("/oauth/jwt/:access_token", handler.GetByAccessToken)

	router.Run(":8080")
}
