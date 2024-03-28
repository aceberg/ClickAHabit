package web

import (
	// "log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/ClickAHabit/internal/db"
	"github.com/aceberg/ClickAHabit/internal/models"
)

var statsMap map[string]models.Stat

func statsHandler(c *gin.Context) {
	var guiData models.GuiData
	var key string
	var ID int

	idStr := c.Param("id")
	if idStr != "total" {
		ID, _ = strconv.Atoi(idStr)
	}

	allChecks = db.Select(appConfig.DBPath, "checks")
	guiData.Config = appConfig

	statsMap = make(map[string]models.Stat)

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

			stat.Checks = append(stat.Checks, check)
			statsMap[key] = stat
		}

		if check.ID == ID {
			guiData.Version = check.Group + ": " + check.Name
		}
	}

	guiData.Stats = statsMap

	c.HTML(http.StatusOK, "header.html", guiData)
	c.HTML(http.StatusOK, "stats.html", guiData)
}
