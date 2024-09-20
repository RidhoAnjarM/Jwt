package controllers

import (
	"main/db"
	"net/http"
	"github.com/gin-gonic/gin"
)

// CreateRole
func CreateRole(c *gin.Context) {
	var role db.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&role).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat peran"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Peran berhasil dibuat"})
}

// Get semua roles
func GetRoles(c *gin.Context) {
	var roles []db.Role
	if err := db.DB.Find(&roles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendapatkan peran"})
		return
	}

	c.JSON(http.StatusOK, roles)
}

// GetRole berdasarkan id
func GetRole(c *gin.Context) {
	id := c.Param("id")
	var role db.Role
	if err := db.DB.First(&role, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Peran tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, role)
}

// UpdateRole 
func UpdateRole(c *gin.Context) {
	id := c.Param("id")
	var role db.Role
	if err := db.DB.First(&role, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Peran tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Save(&role).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui peran"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Peran berhasil diperbarui"})
}

// DeleteRole handler
func DeleteRole(c *gin.Context) {
	id := c.Param("id")
	if err := db.DB.Delete(&db.Role{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus peran"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Peran berhasil dihapus"})
}
