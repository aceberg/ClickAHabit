package web

import (
	"net/http"

	"github.com/gin-gonic/gin"

	// "github.com/aceberg/CheckList/internal/db"
	"github.com/aceberg/CheckList/internal/models"
)

func indexHandler(c *gin.Context) {
	var guiData models.GuiData

	guiData.Config = appConfig

	c.HTML(http.StatusOK, "header.html", guiData)
	c.HTML(http.StatusOK, "index.html", guiData)
}
