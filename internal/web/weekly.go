package web

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/ClickAHabit/internal/models"
)

func weeklyHandler(c *gin.Context) {
	var guiData models.GuiData

	guiData.Config = appConfig

	c.HTML(http.StatusOK, "header.html", guiData)
	c.HTML(http.StatusOK, "weekly.html", guiData)
}
