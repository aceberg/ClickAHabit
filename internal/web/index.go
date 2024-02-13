package web

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/CheckList/internal/db"
	"github.com/aceberg/CheckList/internal/models"
)

func indexHandler(c *gin.Context) {
	var guiData models.GuiData

	guiData.Config = appConfig
	allChecks = db.Select(appConfig.DBPath)

	guiData.Themes = groupList
	// guiData.Checks = setTodayChecks()
	setTodayChecks()
	guiData.Checks = allChecks

	c.HTML(http.StatusOK, "header.html", guiData)
	c.HTML(http.StatusOK, "index.html", guiData)
}

func setTodayChecks() (todayChecks []models.Check) {
	var changedDB bool
	date := time.Now().Format("2006-01-02")

	todayChecks = selectChecksByDate(date)

	for _, plan := range flatPlans {
		if !inSlice(plan, todayChecks) {
			plan.Date = date
			todayChecks = append(todayChecks, plan)
			db.Insert(appConfig.DBPath, plan)
			changedDB = true
		}
	}

	if changedDB {
		allChecks = db.Select(appConfig.DBPath)
		todayChecks = selectChecksByDate(date)
	}

	return todayChecks
}

func selectChecksByDate(date string) (returnChecks []models.Check) {
	for _, check := range allChecks {
		if check.Date == date {
			returnChecks = append(returnChecks, check)
		}
	}
	return returnChecks
}

func inSlice(plan models.Check, todayChecks []models.Check) bool {

	for _, check := range todayChecks {
		if (plan.Group == check.Group) && (plan.Name == check.Name) {
			return true
		}
	}

	return false
}
