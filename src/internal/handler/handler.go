package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	{
		router.POST("/sign-in")
		router.POST("/sign-up")
	}

	api := router.Group("/")
	{
		api.POST("/rector-appeal")
		news := api.Group("/news")
		{
			news.GET("/")
			news.GET("/:id")
			news.GET("/:category")
		}

		timetable := api.Group("/timetable")
		{
			timetable.GET("/")
			timetable.GET("/:group")
		}
	}

	return router
}
