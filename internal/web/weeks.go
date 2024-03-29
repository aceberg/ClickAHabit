package web

import (
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
			check.Date = date
			check.Group = plan.Group
			check.Name = plan.Name
			check.Color = plan.Color
			check.Icon = plan.Icon
			check.Place = plan.Place
			check.Link = plan.Link
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

	yStr := yesterday.Format("2006-01-02")
	wChecks := db.Select(appConfig.DBPath, "weeks")

	for _, c := range wChecks {
		if c.Date == yStr && c.Name == check.Name {
			return c.Count
		}
	}

	return 0
}

// func sameWeek(day, yesterday time.Time) bool {

// }
