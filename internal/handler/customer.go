package handler

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "kreditplus/internal/domain"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "pong"})
    })

    r.POST("/customers", func(c *gin.Context) {
        var customer domain.Customer
        if err := c.ShouldBindJSON(&customer); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        db.Create(&customer)
        c.JSON(http.StatusCreated, customer)
    })

    r.GET("/customers/:id", func(c *gin.Context) {
        var customer domain.Customer
        if err := db.First(&customer, "id = ?", c.Param("id")).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "customer not found"})
            return
        }
        c.JSON(http.StatusOK, customer)
    })
}
