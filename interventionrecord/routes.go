package interventionrecord

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes configures all the routes for intervention record endpoints
func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	records := router.Group("/api/intervention-records")
	{
		records.GET("", GetInterventionRecords(db))
		records.GET("/:id", GetInterventionRecord(db))
		records.POST("", CreateInterventionRecord(db))
		records.PUT("/:id", UpdateInterventionRecord(db))
		records.DELETE("/:id", DeleteInterventionRecord(db))
	}
}
