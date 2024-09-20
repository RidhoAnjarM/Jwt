package controllers

import (
	"github.com/gin-gonic/gin"
	"main/db"
	"net/http"
)

// CreateService
func CreateService(c *gin.Context) {
	var service db.Service
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&service).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat layanan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Layanan berhasil dibuat"})
}

// Get semua service
func GetServices(c *gin.Context) {
	var services []db.Service
	if err := db.DB.Find(&services).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendapatkan layanan"})
		return
	}

	c.JSON(http.StatusOK, services)
}

// GetService
func GetService(c *gin.Context) {
	id := c.Param("id")
	var service db.Service
	if err := db.DB.First(&service, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Layanan tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, service)
}

// UpdateService
func UpdateService(c *gin.Context) {
	id := c.Param("id")
	var service db.Service
	if err := db.DB.First(&service, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Layanan tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Save(&service).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui layanan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Layanan berhasil diperbarui"})
}

// DeleteService
func DeleteService(c *gin.Context) {
	id := c.Param("id")
	if err := db.DB.Delete(&db.Service{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus layanan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Layanan berhasil dihapus"})
}
