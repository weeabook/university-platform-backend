package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/weeabook/src/university-platform-backend/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	{
		router.POST("/sign-in", h.Login)
		router.POST("/sign-up", h.Registration)
	}

	api := router.Group("/")
	{
		appeal := api.Group("/rector-appeal", h.userIdentity)
		{
			appeal.POST("/rector-appeal", h.newAppeal)
		}

		news := api.Group("/news")
		{
			news.GET("/", h.getAllNews)
			news.GET("/:id", h.getNewsById)
			news.GET("/:category", h.getNewsByCategory)
		}

		timetable := api.Group("/timetable", h.userIdentity)
		{
			timetable.GET("/")
			timetable.GET("/:group")
		}
	}

	return router
}
