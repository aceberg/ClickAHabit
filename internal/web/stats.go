package web

import (
	// "log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/ClickAHabit/internal/db"
	"github.com/aceberg/ClickAHabit/internal/models"
)

func statsHandler(c *gin.Context) {
	var guiData models.GuiData
	var key string

	allChecks = db.Select(appConfig.DBPath)
	guiData.Config = appConfig

	statsMap := make(map[string]models.Stat)

	for _, check := range allChecks {
		if check.Count != 0 {
			key = check.Group + ": " + check.Name
			stat, exists := statsMap[key]

			if exists {
				stat.DTotal = stat.DTotal + 1
				stat.CTotal = stat.CTotal + check.Count
			} else {
				stat.Name = check.Name
				stat.Group = check.Group
				stat.DTotal = 1
				stat.CTotal = check.Count
			}

			statsMap[key] = stat
		}
	}

	guiData.Stats = statsMap

	c.HTML(http.StatusOK, "header.html", guiData)
	c.HTML(http.StatusOK, "stats.html", guiData)
}
