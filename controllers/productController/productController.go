package productController

import (
	"net/http"

	"github.com/fauzanfl/belajar-go-restapi/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var config = models.Config{}

func Index(c *gin.Context) {
	var products []models.Product
	config.DB.Find(products)

	c.JSON(http.StatusOK, gin.H{"products": products})
}

func Create(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	config.DB.Create(&product)
	c.JSON(http.StatusCreated, gin.H{"product": product})
}

func Show(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := config.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}

func Update(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if config.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Update failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Update success"})
}

func Delete(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if config.DB.Delete(&product, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Delete failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Delete success"})
}
