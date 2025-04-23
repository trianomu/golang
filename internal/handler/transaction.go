package handler

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "net/http"
    "sync"
    "kreditplus/internal/domain"
)

var txMutex sync.Mutex

func TransactionHandler(r *gin.Engine, db *gorm.DB) {
    r.POST("/transactions", func(c *gin.Context) {
        txMutex.Lock()
        defer txMutex.Unlock()

        var trx domain.Transaction
        if err := c.ShouldBindJSON(&trx); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        if err := db.Create(&trx).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusCreated, trx)
    })
}
