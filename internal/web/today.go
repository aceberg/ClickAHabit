package web

import (
	"time"

	"github.com/aceberg/CheckList/internal/db"
	"github.com/aceberg/CheckList/internal/models"
)

func setTodayChecks() (todayChecks []models.Check) {
	var changedDB bool
	var check models.Check
	date := time.Now().Format("2006-01-02")

	todayChecks = selectChecksByDate(date)

	for _, plan := range allPlans {
		if !inSlice(plan, todayChecks) {
			check.Date = date
			check.Group = plan.Group
			check.Name = plan.Name
			check.Color = plan.Color
			check.Icon = plan.Icon
			check.Place = plan.Place
			todayChecks = append(todayChecks, check)
			db.Insert(appConfig.DBPath, check)
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

	allChecks = db.Select(appConfig.DBPath)

	for _, check := range allChecks {
		if check.Date == date {
			returnChecks = append(returnChecks, check)
		}
	}
	return returnChecks
}

func inSlice(plan models.Plan, todayChecks []models.Check) bool {

	for _, check := range todayChecks {
		if (plan.Group == check.Group) && (plan.Name == check.Name) {
			return true
		}
	}

	return false
}
