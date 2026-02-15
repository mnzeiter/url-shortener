package db

import (
    "fmt"
    "log"
    "os"
    "time"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
    user := getEnv("DB_USER", "root")
    pass := getEnv("DB_PASSWORD", "")
    host := getEnv("DB_HOST", "localhost")
    port := getEnv("DB_PORT", "3306")
    name := getEnv("DB_NAME", "url_shortener")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&loc=Local",
        user, pass, host, port, name,
    )

    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }

    sqlDB, err := DB.DB()
    if err != nil {
        log.Fatalf("failed to get sql.DB: %v", err)
    }

    sqlDB.SetMaxIdleConns(10)
    sqlDB.SetMaxOpenConns(100)
    sqlDB.SetConnMaxLifetime(time.Hour)

    log.Println("Connected to MySQL successfully")
}

func getEnv(key, def string) string {
    if v := os.Getenv(key); v != "" {
        return v
    }
    return def
}
