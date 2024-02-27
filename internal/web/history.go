package web

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/ClickAHabit/internal/db"
	"github.com/aceberg/ClickAHabit/internal/models"
)

func histHandler(c *gin.Context) {
	var guiData models.GuiData

	date := c.Param("date")

	guiData.Config = appConfig
	allChecks = db.Select(appConfig.DBPath)

	if date == "today" {
		date = setToday()
	}
	guiData.Version = date

	for _, check := range allChecks {
		if check.Date == date {
			guiData.Checks = append(guiData.Checks, check)
		}
	}

	c.HTML(http.StatusOK, "header.html", guiData)
	c.HTML(http.StatusOK, "history.html", guiData)
}

func histDel(c *gin.Context) {

	IDstr := c.Param("id")
	ID, err := strconv.Atoi(IDstr)

	if err == nil {
		db.Delete(appConfig.DBPath, ID)
	}

	c.Redirect(http.StatusFound, "/history/today")
}
