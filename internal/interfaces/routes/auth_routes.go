package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mehedicode-lab/go-clean-architecture/internal/interfaces/http"
)

func AuthRoutes(r *gin.RouterGroup, handler *http.AuthHandler) {
	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)
	r.POST("/refresh", handler.RefreshToken)

	protected := r.Group("/")
	protected.Use(http.AuthMiddleware())
	protected.GET("/profile", handler.Profile)
}
