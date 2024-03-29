package web

import (
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	// "github.com/aceberg/ClickAHabit/internal/models"
)

func dateHandler(c *gin.Context) {

	tab := c.Param("tab")
	date := c.Param("date")
	today := setToday()

	if today != lastToday {
		setTodayChecks()
	}

	checks := selectChecksByDate(tab, date)

	// Sort by Name
	sort.Slice(checks, func(i, j int) bool {
		return checks[i].Name < checks[j].Name
	})

	// Sort by Place
	sort.Slice(checks, func(i, j int) bool {
		return checks[i].Place < checks[j].Place
	})

	c.IndentedJSON(http.StatusOK, checks)
}
