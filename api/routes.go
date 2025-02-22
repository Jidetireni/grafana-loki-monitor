package api

import (
	"github.com/gin-gonic/gin"
	"github.com/telexintegrations/grafana-loki-monitor/config"
)

func SetRoute(r *gin.Engine) {
	r.GET("/integration.json", config.GetIntegrationJSON)
	r.POST("/tick", TickHandler)
}
