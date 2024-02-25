package web

import (
	// "log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/ClickAHabit/internal/models"
)

// HeatMapData - data for HeatMap
type HeatMapData struct {
	X string
	Y string
	D string
	V int
}

func statsMore(c *gin.Context) {
	var heatMap []HeatMapData

	key := c.Param("key")
	stat, ok := statsMap[key]

	if ok {
		heatMap = generateHeatMap(stat.Checks)
	}

	c.IndentedJSON(http.StatusOK, heatMap)
}

func generateHeatMap(checks []models.Check) (heatMap []HeatMapData) {
	var heat HeatMapData

	w := 52 // weeks to show

	max := time.Now()
	min := max.AddDate(0, 0, -7*w)

	startDate := weekStartDate(min)
	countMap := countHeat(checks)

	dow := []string{"Mo", "Tu", "We", "Th", "Fr", "Sa", "Su"}

	for _, day := range dow {

		heat.Y = day

		for i := 0; i < w+1; i++ {

			heat.X = strconv.Itoa(i)
			heat.D = startDate.AddDate(0, 0, 7*i).Format("2006-01-02")
			heat.V = countMap[heat.D]
			heatMap = append(heatMap, heat)
		}

		startDate = startDate.AddDate(0, 0, 1)
	}

	return heatMap
}

func weekStartDate(date time.Time) time.Time {
	offset := (int(time.Monday) - int(date.Weekday()) - 7) % 7
	result := date.Add(time.Duration(offset*24) * time.Hour)
	return result
}

func countHeat(checks []models.Check) map[string]int {
	countMap := make(map[string]int)

	for _, check := range checks {
		count, exists := countMap[check.Date]
		if exists {
			countMap[check.Date] = count + check.Count
		} else {
			countMap[check.Date] = check.Count
		}
	}

	return countMap
}
