package routes

import "github.com/gin-gonic/gin"

type AuthRoutes struct {
	authController controllers.AuthController
}

func NewAuthRouter(authController controller.AuthController) AuthRoutes {
	return AuthRoutes{
		authController: authController,
	}
}

func (r AuthRoutes) AuthRoute(router *gin.Engine) {
	public := router.Group("/auth")
	public.POST("/login", r.authController.Login())
	public.POST("/logout", r.authController.Logout())
	public.POST("/verify-token", r.authController.VerifyToken())
}
