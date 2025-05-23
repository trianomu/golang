package main

import (
    "github.com/gin-gonic/gin"
    "kreditplus/configs"
    "kreditplus/internal/domain"
    "kreditplus/internal/handler"
)

func main() {
    configs.ConnectDB() 

    err := configs.DB.AutoMigrate(
        &domain.Customer{},
        &domain.Transaction{},
        &domain.Limit{},
    )
    if err != nil {
        panic("Migration failed: " + err.Error())
    }

    r := gin.Default()
    handler.SetupRoutes(r, configs.DB)
    r.Run(":8080")
}
