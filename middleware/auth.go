package middleware

import (
	"main/controllers"
	"main/db"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Middleware untuk mengecek token JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		// Validasi format token "Bearer <token>"
		if len(tokenString) < 7 || strings.ToUpper(tokenString[:7]) != "BEARER " {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
			c.Abort()
			return
		}

		tokenString = tokenString[7:]

		// Parsing token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return controllers.JwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
			c.Abort()
			return
		}

		// Ambil user dari token
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			email := claims["sub"].(string)

			var user db.User
			if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "User tidak ditemukan"})
				c.Abort()
				return
			}

			c.Set("user", user)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Klaim token tidak valid"})
			c.Abort()
		}

		c.Next()
	}
}
