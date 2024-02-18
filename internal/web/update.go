package web

import (
	"log"
	"net/http"

	// "github.com/aceberg/CheckList/internal/db"
	"github.com/gin-gonic/gin"
)

func updatePlan(c *gin.Context) {
	var resp int

	date := c.Param("date")

	log.Println("UPD DATE:", date)

	c.IndentedJSON(http.StatusOK, resp)
}
