package controllers

import (
	"main/db"
	"net/http"
	"github.com/gin-gonic/gin"
)

// CreateAC
func CreateAC(c *gin.Context) {
	var ac db.AC
	if err := c.ShouldBindJSON(&ac); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&ac).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat AC"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "AC berhasil dibuat"})
}

// Get semua ac
func GetACs(c *gin.Context) {
	var acs []db.AC
	if err := db.DB.Find(&acs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendapatkan AC"})
		return
	}

	c.JSON(http.StatusOK, acs)
}

// GetAC berdasarkan id
func GetAC(c *gin.Context) {
	id := c.Param("id")
	var ac db.AC
	if err := db.DB.First(&ac, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "AC tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, ac)
}

// UpdateAC 
func UpdateAC(c *gin.Context) {
	id := c.Param("id")
	var ac db.AC
	if err := db.DB.First(&ac, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "AC tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&ac); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Save(&ac).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui AC"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "AC berhasil diperbarui"})
}

// DeleteAC berdasarkan id
func DeleteAC(c *gin.Context) {
	id := c.Param("id")
	if err := db.DB.Delete(&db.AC{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus AC"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "AC berhasil dihapus"})
}
