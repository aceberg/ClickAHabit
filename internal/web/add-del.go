package web

import (
	"net/http"
	"strconv"

	"github.com/aceberg/ClickAHabit/internal/db"
	"github.com/gin-gonic/gin"
)

func addHandler(c *gin.Context) {
	var idStr string
	var resp int

	idStr = c.Param("id")
	ID, err := strconv.Atoi(idStr)
	allChecks = db.Select(appConfig.DBPath, "checks")

	if err == nil {
		for _, check := range allChecks {

			if ID == check.ID {
				check.Count = check.Count + 1
				resp = check.Count
				db.Update(appConfig.DBPath, "checks", check, check.ID)
				break
			}
		}
	}

	c.IndentedJSON(http.StatusOK, resp)
}

func delHandler(c *gin.Context) {

	IDstr := c.Param("id")
	ID, err := strconv.Atoi(IDstr)

	if err == nil {
		db.Delete(appConfig.DBPath, "checks", ID)
	}

	c.IndentedJSON(http.StatusOK, "")
}
