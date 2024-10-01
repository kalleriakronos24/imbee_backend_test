package router

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kalleriakronos24/imbee-backend/config"
	v1 "github.com/kalleriakronos24/imbee-backend/controllers/v1"
	"github.com/kalleriakronos24/imbee-backend/utils"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRouter() (router *gin.Engine) {
	router = gin.Default()
	str := []string{"http://localhost:4321"}

	if config.AppConfig.Environment == "PRODUCTION" {
		str = []string{"https://wadahgo.com"}
	}

	configCors := cors.DefaultConfig()
	configCors.AddAllowHeaders("Authorization")
	configCors.AllowOrigins = str
	router.Use(cors.New(configCors))

	commonRoute := router.Group("/")
	v1route := router.Group("/api/v1")

	v1route.Use()
	{

		if config.AppConfig.Environment == "DEVELOPMENT" {
			v1route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}

		fcm := v1route.Group("/fcm")
		{
			fcm.POST("/send", v1.POSTInsertFCM)
		}

		misc := v1route.Group("/misc")
		{
			misc.GET("/ping", v1.Pong)
			misc.POST("/upload", v1.UploadFileSingle)
			misc.POST("/upload-multiple", v1.UploadFileMultiple)
			if config.AppConfig.Environment == "PRODUCTION" {
				misc.POST("/restore/:fileName", utils.AuthOnly, v1.RestoreDatabase)
			}
		}
	}
	fileServingGroupRoute := config.AppConfig.APPUrlStaticFileGroupRoute
	fileServingMainRoute := config.AppConfig.AppUrlStaticFileMainRoute
	fileServing := commonRoute.Group(fileServingGroupRoute)
	{
		//todo improve static file serving security
		workdir, _ := os.Getwd()
		path := filepath.Join(workdir, "../files-uploaded")
		fileServing.StaticFS(fileServingMainRoute, http.Dir(path))
	}
	return router
}
