package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthenticateJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authToken, err := ctx.Cookie("Authorization")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Auth token not found",
			})
			ctx.Abort()
		}

		token := strings.Fields(authToken)
		tokenString := token[1]
		verified, _ := utils.Validate(tokenString)
		if !verified {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Invalid auth token",
			})
			ctx.Abort()
			return
		}

		claims, err := utils.GetClaimsFromToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Invalid auth token",
			})
			ctx.Abort()
			return
		}

		ctx.Set("user_details", claims)
	}
}

func OnlyAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userDetails utils.SignedDetails = c.MustGet("user_details").(utils.SignedDetails)
		if !userDetails.IsAdmin {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Only admin can perform this action"})
			c.Abort()
			return
		} else {
			c.Next()
		}
	}
}
