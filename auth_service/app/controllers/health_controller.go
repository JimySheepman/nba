package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct {
	healthService services.healthService
}

// HealthCheck godoc
// @Summary HealthCheck
// @Description This request is used to check the health of the entire service at once
// @Tags Auth Service
// @Schemes
// @Accept json
// @Produce json
// @Success	200  {object} 	utils.HealthCheckResponse
// @Failure 500  {number} 	http.StatusInternalServerError
// @Router / [GET]
func (c HealthController) HealthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		response := utils.HealthCheck()
		c.JSON(http.StatusOK, response)
	}
}

// HealthCheck godoc
// @Summary Deep HealthCheck
// @Description This request is used to check the health of the every single service at once
// @Tags Auth Service
// @Schemes
// @Accept json
// @Produce json
// @Success	200  {object} 	utils.HealthCheckResponse
// @Failure 500  {number} 	http.StatusInternalServerError
// @Router /deep [GET]
func (c HealthController) DeepHealthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		response := utils.DeepHealthCheck()
		c.JSON(http.StatusOK, response)
	}
}
