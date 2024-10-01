package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kalleriakronos24/imbee-backend/config"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var referer string

		if config.AppConfig.Environment == "DEVELOPMENT" {
			referer = "http://localhost:4321"
		}

		if config.AppConfig.Environment == "PRODUCTION" {
			referer = "http://195.35.20.241"
		}

		c.Writer.Header().Set("Access-Control-Allow-Origin", referer)
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			log.Print("does this running")
			c.Writer.Header().Set("Access-Control-Allow-Origin", referer)
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST,HEAD,PATCH,OPTIONS,GET,PUT,DELETE")
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}

	// var referer string

	// if config.AppConfig.Environment == "DEVELOPMENT" {
	// 	referer = "http://localhost:4321"
	// }

	// if config.AppConfig.Environment == "PRODUCTION" {
	// 	referer = "http://195.35.20.241"
	// }

	// c.Header("Access-Control-Allow-Origin", referer)
	// c.Header("Access-Control-Allow-Credentials", "true")
	// c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	// c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH,OPTIONS,GET,PUT,DELETE")

	// if c.Request.Method == "OPTIONS" {
	// 	c.Header("Access-Control-Allow-Origin", referer)
	// 	c.Header("Access-Control-Allow-Credentials", "true")
	// 	c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	// 	c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH,OPTIONS,GET,PUT,DELETE")
	// 	c.AbortWithStatus(204)
	// 	return
	// }
	// c.Next()
}
