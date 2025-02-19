package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/telexintegrations/grafana-loki-monitor/api"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	}))
	r.GET("/integration.json", getIntegrationJSON)
	r.POST("/tick", api.TickHandler)
	r.Run(":8080")
}
