package configs

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
)

var DB *gorm.DB

func ConnectDB() {
    dsn := "root:@tcp(127.0.0.1:3306)/test?parseTime=true"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }

    DB = db
    log.Println("Database connected successfully:", db)
}


// dsn := "root:password@tcp(localhost:3306)/kreditplus?charset=utf8mb4&parseTime=True&loc=Local"