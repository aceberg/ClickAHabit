package web

import (
	"net/http"
	"strconv"

	"github.com/aceberg/CheckList/internal/db"
	"github.com/gin-gonic/gin"
)

func addHandler(c *gin.Context) {
	var idStr string
	var resp int

	idStr = c.Param("id")
	ID, err := strconv.Atoi(idStr)
	allChecks = db.Select(appConfig.DBPath)

	if err == nil {
		for _, check := range allChecks {

			if ID == check.ID {
				check.Count = check.Count + 1
				resp = check.Count
				db.Update(appConfig.DBPath, check, check.ID)
				break
			}
		}
	}

	c.IndentedJSON(http.StatusOK, resp)
}
