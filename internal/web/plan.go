package web

import (
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/ClickAHabit/internal/models"
	"github.com/aceberg/ClickAHabit/internal/yaml"
)

func planHandler(c *gin.Context) {
	var guiData models.GuiData

	guiData.Config = appConfig

	gr, ok := c.GetQuery("gr")

	for i := range allPlans {
		allPlans[i].ID = i + 1
		if ok && (allPlans[i].Group == gr) {
			guiData.Plans = append(guiData.Plans, allPlans[i])
		}
	}

	if !ok {
		guiData.Plans = allPlans
	}

	// Sort by Name
	sort.Slice(guiData.Plans, func(i, j int) bool {
		return guiData.Plans[i].Name < guiData.Plans[j].Name
	})

	// Sort by Group
	sort.Slice(guiData.Plans, func(i, j int) bool {
		return guiData.Plans[i].Group < guiData.Plans[j].Group
	})

	c.HTML(http.StatusOK, "header.html", guiData)
	c.HTML(http.StatusOK, "plan.html", guiData)
}

func planDel(c *gin.Context) {

	IDstr := c.Param("id")
	ID, err := strconv.Atoi(IDstr)

	if err == nil {
		deletePlan(ID)
		yaml.Write(appConfig.DirPath+"/plan.yaml", allPlans)
	}

	c.Redirect(http.StatusFound, "/plan/")
}

func deletePlan(ID int) {
	var newPlans []models.Plan

	for _, plan := range allPlans {
		if plan.ID != ID {
			newPlans = append(newPlans, plan)
		}
	}
	allPlans = newPlans
}
