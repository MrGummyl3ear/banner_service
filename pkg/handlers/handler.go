package handler

import (
	"banner_service/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	banner := router.Group("/banner")
	{
		banner.GET("", h.adminIdentity, h.getAllBanners)
		banner.POST("", h.adminIdentity, h.createBanner)
		banner.PATCH("/:id", h.adminIdentity, h.modifyBanner)
		banner.DELETE("/:id", h.adminIdentity, h.deleteBanner)
	}

	router.GET("/user_banner", h.userIdentity, h.getUserBanner)

	return router
}
