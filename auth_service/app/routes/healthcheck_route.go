package routes

import (
	"auth-service/app/controllers"

	"github.com/gin-gonic/gin"
)

type HealthRoutes struct {
	healthController controllers.HealthController
}

func NewHealthRouter(hc controllers.HealthController) HealthRoutes {
	return HealthRoutes{healthController: hc}
}

func (hr HealthRoutes) HealthRoute(router *gin.Engine) {
	public := router.Group("")
	public.GET("/", hr.healthController.HealthCheck())
	public.GET("/deep", hr.healthController.DeepHealthCheck())
}
