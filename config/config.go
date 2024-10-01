package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port        int
	Environment string
	Debug       bool

	APPUrl                     string
	APPUrlStaticFileGroupRoute string
	AppUrlStaticFileMainRoute  string
	APPUrlClientSide           string
	APPUrlClientSideProd       string

	DBHost                 string
	DBPort                 int
	DBDatabase             string
	DBUsername             string
	DBPassword             string
	DBUrl                  string
	DBContainerName        string
	DBBackupTimerInMinutes uint64
	DBRestoreApiKey        string

	JWTSecret string

	SOCKETEnabled bool

	MAILHost     string
	MAILPort     int
	MAILUsername string
	MAILPassword string
	MAILFrom     string
}

var AppConfig Config

func InitializeAppConfig() Config {
	viper.SetConfigName(".env") // allow directly reading from .env file
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("/")
	viper.AllowEmptyEnv(true)
	viper.AutomaticEnv()
	_ = viper.ReadInConfig()

	AppConfig.Port = viper.GetInt("PORT")
	AppConfig.Environment = viper.GetString("ENVIRONMENT")
	AppConfig.Debug = viper.GetBool("DEBUG")

	AppConfig.APPUrl = viper.GetString("APP_URL")
	AppConfig.APPUrlStaticFileGroupRoute = viper.GetString("APP_STATIC_FILE_GROUP_ROUTE")
	AppConfig.AppUrlStaticFileMainRoute = viper.GetString("APP_STATIC_FILE_MAIN_ROUTE")
	AppConfig.APPUrlClientSide = viper.GetString("APP_URL_CLIENT_SIDE")
	AppConfig.APPUrlClientSide = viper.GetString("APP_URL_CLIENT_SIDE_PROD")

	AppConfig.DBHost = viper.GetString("DB_HOST")
	AppConfig.DBPort = viper.GetInt("DB_PORT")
	AppConfig.DBDatabase = viper.GetString("DB_DATABASE")
	AppConfig.DBUsername = viper.GetString("DB_USERNAME")
	AppConfig.DBPassword = viper.GetString("DB_PASSWORD")
	AppConfig.DBUrl = viper.GetString("DATABASE_URL")
	AppConfig.DBContainerName = viper.GetString("DB_CONTAINER_NAME")
	AppConfig.DBBackupTimerInMinutes = viper.GetUint64("DB_BACKUP_TIMER_IN_MINUTES")
	AppConfig.DBRestoreApiKey = viper.GetString("DB_RESTORE_API_KEY")

	AppConfig.JWTSecret = viper.GetString("JWT_SECRET")
	AppConfig.SOCKETEnabled = viper.GetBool("SOCKET_ENABLED")

	AppConfig.MAILHost = viper.GetString("MAIL_HOST")
	AppConfig.MAILPort = viper.GetInt("MAIL_PORT")
	AppConfig.MAILUsername = viper.GetString("MAIL_PASSWORD")
	AppConfig.MAILFrom = viper.GetString("MAIL_FROM")

	log.Printf("[INIT] configuration loaded")
	return AppConfig
}
