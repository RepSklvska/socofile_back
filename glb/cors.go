package glb

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CORS Middleware

var (
	AllAllowCORS = gin.HandlerFunc(func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "*")
		ctx.Header("Access-Control-Allow-Headers", "*")
		ctx.Header("Content-Type", "application/json")
		
		if ctx.Request.Method != "OPTIONS" {
			ctx.Next()
		} else {
			ctx.AbortWithStatus(http.StatusOK)
		}
	})
	MyDefaultCORS = cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000",
			"http://192.168.2.120:3000",
			"http://192.168.1.6:3000"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"*", "content-type", "content-length"},
		AllowCredentials: true,
	})
)
