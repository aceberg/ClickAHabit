package web

import (
	// "log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/CheckList/internal/models"
	"github.com/aceberg/CheckList/internal/yaml"
)

func editHandler(c *gin.Context) {
	var guiData models.GuiData
	var id int

	guiData.Config = appConfig

	idStr := c.Param("id")

	if idStr != "new" {
		id, _ = strconv.Atoi(idStr)

		for _, plan := range allPlans {
			if plan.ID == id {
				guiData.OnePlan = plan
				break
			}
		}
	}

	c.HTML(http.StatusOK, "header.html", guiData)
	c.HTML(http.StatusOK, "plan-edit.html", guiData)
}

func savePlanHandler(c *gin.Context) {
	var plan models.Plan

	id := c.PostForm("id")
	plan.Group = c.PostForm("group")
	plan.Name = c.PostForm("name")
	plan.Color = c.PostForm("color")
	plan.Icon = c.PostForm("icon")
	place := c.PostForm("place")

	plan.ID, _ = strconv.Atoi(id)
	plan.Place, _ = strconv.Atoi(place)

	// log.Println("PLAN TO SAVE:", plan)

	if plan.ID > 0 {
		deletePlan(plan.ID)
	}
	allPlans = append(allPlans, plan)
	yaml.Write(appConfig.DirPath+"/plan.yaml", allPlans)

	c.Redirect(http.StatusFound, "/plan/")
}
