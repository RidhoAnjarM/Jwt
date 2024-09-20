package middleware

import (
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

		if len(tokenString) < 7 || strings.ToUpper(tokenString[:7]) != "BEARER " {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
			c.Abort()
			return
		}

		tokenString = tokenString[7:]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("your_secret_key"), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			email := claims["sub"].(string)

			var user db.User
			if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "User tidak ditemukan"})
				c.Abort()
				return
			}

			// Simpan user ke context
			c.Set("user", user)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Klaim token tidak valid"})
			c.Abort()
		}

		c.Next()
	}
}

// Middleware untuk mengecek peran role
func RoleMiddleware(allowedRoles ...uint) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil data user dari context
		user, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "Tidak ada pengguna yang diautentikasi"})
			c.Abort()
			return
		}

		userRoleID := user.(db.User).RoleID

		// Cek apakah userRoleID ada di allowedRoles
		for _, roleID := range allowedRoles {
			if userRoleID == roleID {
				c.Next()
				return
			}
		}

		// Jika tidak cocok, tolak akses
		c.JSON(http.StatusForbidden, gin.H{"error": "Akses ditolak"})
		c.Abort()
	}
}