package http

import (
	"github.com/apm-dev/go-oauth-service/src/domain/usecases"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type JWTHandler interface {
	GetByAccessToken(*gin.Context)
}

func GetNewJWTHandler(s usecases.JWTService) JWTHandler {
	return &jwtHandler{svc: s}
}

type jwtHandler struct {
	svc usecases.JWTService
}

func (h *jwtHandler) GetByAccessToken(c *gin.Context) {
	jwt, err := h.svc.GetByAccessToken(strings.TrimSpace(c.Param("access_token")))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, jwt)
}
