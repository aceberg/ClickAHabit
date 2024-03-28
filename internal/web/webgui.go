package web

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/ClickAHabit/internal/check"
	"github.com/aceberg/ClickAHabit/internal/conf"
	"github.com/aceberg/ClickAHabit/internal/db"
	"github.com/aceberg/ClickAHabit/internal/yaml"
)

// Gui - start web server
func Gui(dirPath, nodePath string) {

	confPath := dirPath + "/config.yaml"
	check.Path(confPath)

	appConfig = conf.Get(confPath)

	appConfig.DirPath = dirPath
	appConfig.DBPath = dirPath + "/sqlite1.db"
	check.Path(appConfig.DBPath)
	appConfig.ConfPath = confPath
	appConfig.NodePath = nodePath

	log.Println("INFO: starting web gui with config", appConfig.ConfPath)

	db.Create(appConfig.DBPath)
	allChecks = db.Select(appConfig.DBPath, "checks")
	allPlans = yaml.Read(appConfig.DirPath + "/plan.yaml")

	address := appConfig.Host + ":" + appConfig.Port

	log.Println("=================================== ")
	log.Printf("Web GUI at http://%s", address)
	log.Println("=================================== ")

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	templ := template.Must(template.New("").ParseFS(templFS, "templates/*"))
	router.SetHTMLTemplate(templ) // templates

	router.StaticFS("/fs/", http.FS(pubFS)) // public

	router.GET("/", indexHandler)            // index.go
	router.GET("/add/:id", addHandler)       // add-del.go
	router.GET("/config/", configHandler)    // config.go
	router.GET("/date/:date", dateHandler)   // date.go
	router.GET("/del/:id", delHandler)       // add-del.go
	router.GET("/plan/", planHandler)        // plan.go
	router.GET("/planedit/:id", editHandler) // plan-edit.go
	router.GET("/plandel/:id", planDel)      // plan.go
	router.GET("/stats/:id", statsHandler)   // stats.go
	router.GET("/smore/:key", statsMore)     // stats-more.go
	router.GET("/update/:date", updatePlan)  // update.go
	router.GET("/weekly/", weeklyHandler)    // weekly.go

	router.POST("/config/", saveConfigHandler) // config.go
	router.POST("/planedit/", savePlanHandler) // plan-edit.go

	err := router.Run(address)
	check.IfError(err)
}
