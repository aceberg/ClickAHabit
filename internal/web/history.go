package web

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/CheckList/internal/db"
	"github.com/aceberg/CheckList/internal/models"
)

func histHandler(c *gin.Context) {
	var guiData models.GuiData

	guiData.Config = appConfig
	allChecks = db.Select(appConfig.DBPath)

	guiData.Checks = allChecks

	c.HTML(http.StatusOK, "header.html", guiData)
	c.HTML(http.StatusOK, "history.html", guiData)
}

func histDel(c *gin.Context) {

	IDstr := c.Param("id")
	ID, err := strconv.Atoi(IDstr)

	if err == nil {
		db.Delete(appConfig.DBPath, ID)
	}

	c.Redirect(http.StatusFound, "/history/")
}
