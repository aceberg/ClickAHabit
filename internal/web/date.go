package web

import (
	"log"
	// "net/http"

	"github.com/gin-gonic/gin"
	// "github.com/aceberg/CheckList/internal/models"
)

func dateHandler(c *gin.Context) {

	date := c.Param("date")

	checks := selectChecksByDate(date)

	log.Println("DATE =", date)
	log.Println("CHECKS =", checks)
}
