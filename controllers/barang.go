package controllers

import (
	"net/http"
	"rest/config"
	"rest/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BarangController struct {
	Db *gorm.DB
}

func Brg() *BarangController {
	db := config.InitDb()
	db.AutoMigrate(&models.Barang{})
	return &BarangController{Db: db}
}

func (repository *BarangController) CreateBarang(c *gin.Context) {
	var user models.Barang
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ada Data yang kosong"})
		return
	} else {
		models.CreateBarang(repository.Db, &user)
		c.JSON(http.StatusOK, gin.H{"data": user})
	}

}

func (repository *BarangController) GetAllBarang(c *gin.Context) {
	var user []models.Barang
	err := models.GetBarang(repository.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}
