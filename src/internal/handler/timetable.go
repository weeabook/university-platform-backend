package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) getAllTimeTables(c *gin.Context) {
	timetables, err := h.services.TimetableService.GetAll()
	if err != nil {
		logrus.Error(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	SendJSONResponse(c, "timetables", timetables)
}

func (h *Handler) getByGroup(c *gin.Context) {
	group, err := strconv.Atoi(c.Param("group"))
	if err != nil {
		logrus.Error(err)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	timetable, err := h.services.TimetableService.GetByGroup(group)
	if err != nil {
		logrus.Error(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	http.ServeFile(c.Writer, c.Request, "../../uploads/"+strings.Split(timetable, "/")[1])
}
