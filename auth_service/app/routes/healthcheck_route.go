package routes

import "github.com/gin-gonic/gin"

type HealthRoutes struct {
	healthController controllers.HealthController
}

func NewHealthRouter(hc controllers.HealthController) HealthRoutes {
	return HealthRoutes{
		healthController: hc,
	}
}

func (r HealthRoutes) HealthRoute(router *gin.Engine) {
	public := router.Group("")
	public.GET("/", r.healthController.HealthCheck())
	public.GET("/deep", r.healthController.DeepHealthCheck())
}
