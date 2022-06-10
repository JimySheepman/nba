package routes

import (
	"auth-service/app/controllers"

	"github.com/gin-gonic/gin"
)

type AuthRoutes struct {
	authController controllers.AuthController
}

func NewAuthRouter(authController controllers.AuthController) AuthRoutes {
	return AuthRoutes{authController: authController}
}

func (ar AuthRoutes) AuthRoute(router *gin.Engine) {
	public := router.Group("/auth")
	public.POST("/login", ar.authController.Login())
	public.POST("/logout", ar.authController.Logout())
	public.POST("/verify-token", ar.authController.VerifyToken())
}
