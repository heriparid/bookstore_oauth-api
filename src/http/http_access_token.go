package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/heriparid/oauth-api/src/domain/access_token"
	"github.com/heriparid/oauth-api/src/utils/errors"
)

// AccessTokenHandler interface
type AccessTokenHandler interface {
	GetByID(*gin.Context)
	Create(*gin.Context)
	UpdateExpiryTime(*gin.Context)
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

func (h *accessTokenHandler) Create(c *gin.Context) {
	var at access_token.AccessToken
	if err := c.ShouldBindJSON(&at); err != nil {
		restErr := errors.NewBadRequestError("invalid jsob body")
		c.JSON(restErr.Status, restErr)
		return
	}

	if err := h.service.Create(at); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, at)
}

func (h *accessTokenHandler) UpdateExpiryTime(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented yet")
}
