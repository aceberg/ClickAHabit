package web

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	// "github.com/aceberg/CheckList/internal/db"
	"github.com/aceberg/CheckList/internal/models"
)

func indexHandler(c *gin.Context) {
	var guiData models.GuiData

	guiData.Config = appConfig

	generatePlan()
	guiData.Plans = plans
	guiData.Checks = checks

	c.HTML(http.StatusOK, "header.html", guiData)
	c.HTML(http.StatusOK, "index.html", guiData)
}

func generatePlan() {
	var plan models.Plan
	var item models.Item

	plans = []models.Plan{}

	plan.Group = "Home"

	item.Name = "Check 1"
	item.Color = ""
	item.Count = 1
	plan.Items = append(plan.Items, item)

	item.Name = "Check 2"
	item.Color = "#7C68E9"
	plan.Items = append(plan.Items, item)

	item.Name = "Long Check 3"
	item.Color = "#59BF40"
	plan.Items = append(plan.Items, item)

	plans = append(plans, plan)
	plan.Group = "Test"

	item.Name = "Check 2"
	item.Color = "#7C68E9"
	item.Count = 3
	plan.Items = append(plan.Items, item)

	item.Count = 0

	item.Name = "Check 2"
	item.Color = "#7C68E9"
	plan.Items = append(plan.Items, item)

	item.Name = "Check 2"
	item.Color = "#7C68E9"
	plan.Items = append(plan.Items, item)

	item.Name = "Check 2"
	item.Color = "#7C68E9"
	plan.Items = append(plan.Items, item)

	plans = append(plans, plan)
}
