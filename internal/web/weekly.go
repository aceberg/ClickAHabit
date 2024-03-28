package web

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/ClickAHabit/internal/db"
	"github.com/aceberg/ClickAHabit/internal/models"
)

func weeklyHandler(c *gin.Context) {
	var guiData models.GuiData

	guiData.Config = appConfig

	setTodayChecks()

	c.HTML(http.StatusOK, "header.html", guiData)
	c.HTML(http.StatusOK, "weekly.html", guiData)
}

func setWeeklyForDate(date string) (wChecks []models.Check) {
	var changedDB bool
	var check models.Check

	wChecks = selectChecksByDate("weeks", date)

	for _, plan := range allPlans {
		if !inSlice(plan, wChecks) && !plan.Pause && plan.Weekly {
			check.Date = date
			check.Group = plan.Group
			check.Name = plan.Name
			check.Color = plan.Color
			check.Icon = plan.Icon
			check.Place = plan.Place
			check.Link = plan.Link
			wChecks = append(wChecks, check)
			db.Insert(appConfig.DBPath, "weeks", check)
			changedDB = true
		}
	}

	if changedDB {
		wChecks = selectChecksByDate("weeks", date)
	}

	return wChecks
}
