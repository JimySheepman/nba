package main

import "auth-service/app"

// @title           Auth Service API
// @version         1.0.0
// @description     Users can login and get a token and use it to access other APIs
// @termsOfService  http://swagger.io/terms/

// @contact.name merlins
// @contact.email merlins.nightowl@protonmail.com

// @host      localhost:3012
// @securityDefinitions.apiKey ApiKeyAuth
// @type apiKey
// @in cookie
// @name Authorization
func main() {
	app.Start()
}
