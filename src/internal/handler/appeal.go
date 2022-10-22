package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/weeabook/src/university-platform-backend/internal/model"
	"net/http"
)

func (h *Handler) newAppeal(c *gin.Context) {
	var input model.InputAppeal
	if err := c.BindJSON(&input); err != nil {
		logrus.Error(err)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userId, err := getUserId(c)
	if err != nil {
		logrus.Error(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	err = h.services.AppealService.Create(userId, input)
	if err != nil {
		logrus.Error(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	SendJSONResponse(c, "status", "success")

}
