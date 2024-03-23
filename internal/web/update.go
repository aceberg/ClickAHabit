package web

import (
	// "log"
	"net/http"
	"sync"

	// "github.com/aceberg/ClickAHabit/internal/db"
	"github.com/gin-gonic/gin"
)

var mu sync.Mutex

func updatePlan(c *gin.Context) {
	var resp int

	mu.Lock()

	date := c.Param("date")
	setChecksForDate(date) // today.go

	mu.Unlock()

	// log.Println("UPD DATE:", date)

	c.IndentedJSON(http.StatusOK, resp)
}
