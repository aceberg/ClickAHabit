package web

import (
	// "log"
	"time"

	"github.com/aceberg/ClickAHabit/internal/db"
	"github.com/aceberg/ClickAHabit/internal/models"
)

func setWeeklyForDate(date string) (wChecks []models.Check) {
	var changedDB bool
	var check models.Check

	wChecks = selectChecksByDate("weeks", date)

	for _, plan := range allPlans {
		if !inSlice(plan, wChecks) && !plan.Pause && plan.Weekly {
			check = copyPlan(plan)
			check.Date = date
			check.Count = weekCount(check)
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

func weekCount(check models.Check) int {

	day, _ := time.Parse("2006-01-02", check.Date)
	yesterday := day.AddDate(0, 0, -1)

	wChecks := db.Select(appConfig.DBPath, "weeks")

	// log.Println("NAME", check.Name)

	for sameWeek(day, yesterday) {
		yStr := yesterday.Format("2006-01-02")
		// log.Println("SEARCHING DATE:", yStr)

		for _, c := range wChecks {
			if c.Date == yStr && c.Name == check.Name && c.Count > 0 {
				// log.Println("FOUND DATE:", yStr)
				return c.Count
			}
		}
		yesterday = yesterday.AddDate(0, 0, -1)
	}

	// log.Println("FOUND NOTHING!")
	return 0
}

func sameWeek(day, beforeDay time.Time) bool {

	weekStart := weekStartDate(day).AddDate(0, 0, -1)

	return weekStart.Before(beforeDay)
}
