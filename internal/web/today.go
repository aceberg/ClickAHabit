package web

import (
	"time"

	"github.com/aceberg/ClickAHabit/internal/db"
	"github.com/aceberg/ClickAHabit/internal/models"
)

func setToday() string {
	date := time.Now().Format("2006-01-02")

	return date
}

func setTodayChecks() {

	date := setToday()
	lastToday = date
	setChecksForDate(date)
	setWeeklyForDate(date) // weeks.go
}

func setChecksForDate(date string) (todayChecks []models.Check) {
	var changedDB bool
	var check models.Check

	todayChecks = selectChecksByDate("checks", date)

	for _, plan := range allPlans {
		if !inSlice(plan, todayChecks) && !plan.Pause && !plan.Weekly {
			check = copyPlan(plan)
			check.Date = date

			todayChecks = append(todayChecks, check)
			db.Insert(appConfig.DBPath, "checks", check)
			changedDB = true
		}
	}

	if changedDB {
		allChecks = db.Select(appConfig.DBPath, "checks")
		todayChecks = selectChecksByDate("checks", date)
	}

	return todayChecks
}

func selectChecksByDate(tabname string, date string) (returnChecks []models.Check) {

	checks := db.Select(appConfig.DBPath, tabname)

	for _, check := range checks {
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

func copyPlan(plan models.Plan) (check models.Check) {

	check.Group = plan.Group
	check.Name = plan.Name
	check.Color = plan.Color
	check.Icon = plan.Icon
	check.Place = plan.Place
	check.Link = plan.Link

	return check
}
