package app

import (
	_ "auth-service/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var (
	userRepository repository.UserRepository
	authService    services.AuthService
	authController controllers.AuthController
	authRoutes     routes.AuthRoutes
	healthRoutes   routes.HealthRoutes
)

func Start() {
	userDB := confings.ConnectDB()
	userRepository = repository.NewUserRepository(userDB)
	authService = services.NewAuthServiceImpl(userRepository)
	authController = *controllers.NewAuthController(authService)
	authRoutes = routes.NewAuthRouter(authController)

	go auth.InitialiseAuthServer()

	router := InitiseRestServer()
	router.Run(configs.EnvAuthHost() + ":" + configs.EnvPORT())
}

func InitiseRestServer() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authRoutes.AuthRoute(router)
	healthRoutes.HealthRoute(router)
	return router
}
