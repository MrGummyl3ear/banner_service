package handler

import (
	"banner_service/pkg/service"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/singleflight"
)

type Handler struct {
	services     *service.Service
	singleflight singleflight.Group
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	h.singleflight = singleflight.Group{}
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	banner := router.Group("/banner", h.adminIdentity)
	{
		banner.GET("", h.getAllBanners)
		banner.POST("", h.createBanner)
		banner.PATCH("/:id", h.modifyBanner)
		banner.DELETE("/:id", h.deleteBanner)
	}

	router.GET("/user_banner", h.userIdentity, h.getUserBanner)

	return router
}
