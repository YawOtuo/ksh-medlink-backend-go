package interventionrecord

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetInterventionRecords returns all intervention records
func GetInterventionRecords(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var records []InterventionRecord
		result := db.Find(&records)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(http.StatusOK, records)
	}
}

// GetInterventionRecord returns a specific intervention record by ID
func GetInterventionRecord(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var record InterventionRecord
		result := db.First(&record, id)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Intervention record not found"})
			return
		}
		c.JSON(http.StatusOK, record)
	}
}

// CreateInterventionRecord creates a new intervention record
func CreateInterventionRecord(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var record InterventionRecord
		if err := c.ShouldBindJSON(&record); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		result := db.Create(&record)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(http.StatusCreated, record)
	}
}

// UpdateInterventionRecord updates an existing intervention record
func UpdateInterventionRecord(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var record InterventionRecord
		if err := db.First(&record, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Intervention record not found"})
			return
		}
		if err := c.ShouldBindJSON(&record); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		result := db.Save(&record)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(http.StatusOK, record)
	}
}

// DeleteInterventionRecord deletes an intervention record
func DeleteInterventionRecord(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var record InterventionRecord
		if err := db.First(&record, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Intervention record not found"})
			return
		}
		result := db.Delete(&record)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Intervention record deleted successfully"})
	}
}
