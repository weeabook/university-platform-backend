package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"strconv"
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
	c.Writer.Header().Set("Content-Disposition", "attachment; filename="+timetable)
	c.Writer.Header().Set("Content-Type", c.Request.Header.Get("Content-Type"))
	file, err := os.Open("./" + timetable)
	if err != nil {
		logrus.Error(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	io.Copy(c.Writer, file)
}
