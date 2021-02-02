package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/heriparid/oauth-api/src/domain/access_token"
)

// AccessTokenHandler interface
type AccessTokenHandler interface {
	GetByID(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

// NewHandler create new instance
func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (h *accessTokenHandler) GetByID(c *gin.Context) {
	accessTokenID := c.Param("access_token_id")

	ac, err := h.service.GetByID(accessTokenID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, ac)
}
