package web

import (
	"net/http"
	"strconv"

	"github.com/aceberg/CheckList/internal/db"
	"github.com/gin-gonic/gin"
	// "github.com/aceberg/CheckList/internal/models"
)

func addHandler(c *gin.Context) {
	var idStr string
	var resp int

	ID := c.Param("id")
	allChecks = db.Select(appConfig.DBPath)

	for _, check := range allChecks {
		idStr = strconv.Itoa(check.ID)
		if ID == idStr {
			check.Count = check.Count + 1
			resp = check.Count
			db.Update(appConfig.DBPath, check, check.ID)
			break
		}
	}

	c.IndentedJSON(http.StatusOK, resp)
}
