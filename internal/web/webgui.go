package web

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/CheckList/internal/check"
	"github.com/aceberg/CheckList/internal/conf"
	"github.com/aceberg/CheckList/internal/db"
	"github.com/aceberg/CheckList/internal/models"
	"github.com/aceberg/CheckList/internal/yaml"
)

// Gui - start web server
func Gui(dirPath, nodePath string) {

	confPath := dirPath + "/config.yaml"
	check.Path(confPath)

	appConfig = conf.Get(confPath)

	appConfig.DirPath = dirPath
	appConfig.DBPath = dirPath + "/sqlite.db"
	check.Path(appConfig.DBPath)
	appConfig.ConfPath = confPath
	appConfig.NodePath = nodePath

	log.Println("INFO: starting web gui with config", appConfig.ConfPath)

	db.Create(appConfig.DBPath)
	allChecks = db.Select(appConfig.DBPath)
	allPlans = yaml.Read(dirPath + "/plan.yaml")
	plansToSlice() // webgui.go

	address := appConfig.Host + ":" + appConfig.Port

	log.Println("=================================== ")
	log.Printf("Web GUI at http://%s", address)
	log.Println("=================================== ")

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	templ := template.Must(template.New("").ParseFS(templFS, "templates/*"))
	router.SetHTMLTemplate(templ) // templates

	router.StaticFS("/fs/", http.FS(pubFS)) // public

	router.GET("/", indexHandler)         // index.go
	router.GET("/add/:id", addHandler)    // add.go
	router.GET("/config/", configHandler) // config.go
	router.GET("/histdel/:id", histDel)   // history.go
	router.GET("/history/", histHandler)  // history.go

	router.POST("/config/", saveConfigHandler) // config.go

	err := router.Run(address)
	check.IfError(err)
}

func plansToSlice() {
	var oneFlatPlan models.Check

	groupList = []string{}
	flatPlans = []models.Check{}

	for _, plan := range allPlans {
		groupList = append(groupList, plan.Group)

		for _, item := range plan.Items {
			oneFlatPlan = models.Check{}

			oneFlatPlan.Group = plan.Group
			oneFlatPlan.Name = item.Name
			oneFlatPlan.Color = item.Color

			flatPlans = append(flatPlans, oneFlatPlan)
		}
	}
}
