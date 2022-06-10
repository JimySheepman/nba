package app

import (
	"auth-service/app/controllers"
	"auth-service/app/routes"
	"auth-service/configs"
	_ "auth-service/docs"
	"auth-service/domain/repository"
	"auth-service/domain/services"
	"auth-service/grpc/auth"

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
	userDB := configs.ConnectDB()
	userRepository = repository.NewUserRepository(userDB)
	authService = services.NewAuthServiceImpl(userRepository)
	authController = *controllers.NewAuthController(authService)
	authRoutes = routes.NewAuthRouter(authController)

	go auth.InitialiseAuthServer()

	router := InitialiseRestServer()
	router.Run(configs.EnvAuthHost() + ":" + configs.EnvPORT())
}

func InitialiseRestServer() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authRoutes.AuthRoute(router)
	healthRoutes.HealthRoute(router)
	return router
}
