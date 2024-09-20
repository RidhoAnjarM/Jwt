package controllers

import (
	"main/db"
	"net/http"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser 
func CreateUser(c *gin.Context) {
	var user db.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// untuk memeriksa apakah email sudah terdaftar
	var existingUser db.User
	if err := db.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email sudah terdaftar"})
		return
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal meng-enkripsi password"})
		return
	}

	// Cari role berdasarkan RoleID
	var role db.Role
	if err := db.DB.First(&role, user.RoleID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Role tidak ditemukan"})
		return
	}

	// Buat user baru
	users := db.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: string(hashedPassword),
		Phone:    user.Phone,
		Gender:   user.Gender,
		Photo:    user.Photo,
		Address:  user.Address,
		RoleID:   role.ID,
	}

	// Simpan user ke database
	if err := db.DB.Create(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat pengguna"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pengguna berhasil dibuat"})
}

// Get semua users
func GetUsers(c *gin.Context) {
	var users []db.User
	if err := db.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendapatkan pengguna"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUser berdasarkan id
func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user db.User
	if err := db.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pengguna tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUser
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user db.User
	if err := db.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pengguna tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui pengguna"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pengguna berhasil diperbarui"})
}

// DeleteUser
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := db.DB.Delete(&db.User{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus pengguna"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pengguna berhasil dihapus"})
}
