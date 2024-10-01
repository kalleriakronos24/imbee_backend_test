package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/kalleriakronos24/imbee-backend/fb"
	"github.com/kalleriakronos24/imbee-backend/migrations"
	"github.com/kalleriakronos24/imbee-backend/rabbitmq"
	"github.com/kalleriakronos24/imbee-backend/utils"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/kalleriakronos24/imbee-backend/config"
	"github.com/kalleriakronos24/imbee-backend/docs"
	"github.com/kalleriakronos24/imbee-backend/pkg/sockets"
	"github.com/kalleriakronos24/imbee-backend/router"
	"github.com/kalleriakronos24/imbee-backend/services"
	"github.com/spf13/viper"
)

func init() {
	config.InitializeAppConfig()
	if !config.AppConfig.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
}

func runServer() {

	/*
	* set server timezone to Jakarta, Indonesia
	* so is any request from client will converted to our TimeZone
	 */
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Fatalln(err)
		return
	}
	time.Local = loc
	// ============================================================

	// firebase settings
	fb.FirebaseApp()
	// ===================================================

	// swagger configs
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Imbee server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:3009"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	// ====================================================================

	// rabbitmq receivers
	rabbitmq.RabbitMQApp()
	//rabbitmq.ReceiveQueue()
	// ======================================================

	//automatic database backup
	if config.AppConfig.Environment == "PRODUCTION" {
		go utils.DatabaseBackupCron()
	}

	// initialize db and migrations
	if err := services.InitializeServices(); err != nil {
		log.Fatalln(err)
	}
	migrations.Migrate()
	// ==================================================

	// serve all routes and routes configuration
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.AppConfig.Port),
		Handler:        router.InitializeRouter(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
	// ==========================================
}

func main() {
	if viper.GetBool("SOCKET_ENABLED") {
		server := socketio.NewServer(nil)
		var serv = sockets.RunSocketConnection(server)
		http.Handle("/socket.io/", serv)
	}
	runServer()
}
