package app

import (
	"github.com/gin-gonic/gin"
	"github.com/heriparid/oauth-api/src/domain/access_token"
	"github.com/heriparid/oauth-api/src/http"
	"github.com/heriparid/oauth-api/src/repository/db"
)

var (
	router = gin.Default()
)

// StartApplication to initialize dependencies
func StartApplication() {
	dbRepo := db.NewRepository()
	atService := access_token.NewService(dbRepo)
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)
	router.Run(":8080")
}
