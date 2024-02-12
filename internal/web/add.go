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

	ID := c.Param("id")

	for _, check := range allChecks {
		idStr = strconv.Itoa(check.ID)
		if ID == idStr {
			check.Count = check.Count + 1
			db.Update(appConfig.DBPath, check, check.ID)
		}
	}

	c.Redirect(http.StatusFound, "/")
}
