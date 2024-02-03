package web

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/CheckList/internal/check"
	"github.com/aceberg/CheckList/internal/conf"
	// "github.com/aceberg/CheckList/internal/db"
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
	appConfig.Icon = icon

	log.Println("INFO: starting web gui with config", appConfig.ConfPath)

	// db.Create(appConfig.DBPath)

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
	router.GET("/config/", configHandler) // config.go

	router.POST("/config/", saveConfigHandler) // config.go

	err := router.Run(address)
	check.IfError(err)
}
