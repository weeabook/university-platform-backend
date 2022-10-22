package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/weeabook/src/university-platform-backend/internal/model"
	"net/http"
)

func (h *Handler) Registration(c *gin.Context) {
	var input model.User

	if err := c.BindJSON(&input); err != nil {
		logrus.Error(err)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.AuthService.CreateUser(input)
	if err != nil {
		logrus.Error(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	SendJSONResponse(c, "id", id)
}

func (h *Handler) Login(c *gin.Context) {
	var input model.InputUser
	if err := c.BindJSON(&input); err != nil {
		logrus.Error(err)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.services.GenerateToken(input.Email, input.Password)
	if err != nil {
		logrus.Error(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	SendJSONResponse(c, "token", token)

}
