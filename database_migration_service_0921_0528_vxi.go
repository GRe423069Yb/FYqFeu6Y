// 代码生成时间: 2025-09-21 05:28:04
package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// DatabaseMigrationService 负责数据库迁移的工具
type DatabaseMigrationService struct {
    db *gorm.DB
}

// NewDatabaseMigrationService 初始化数据库迁移服务
func NewDatabaseMigrationService() *DatabaseMigrationService {
    // 连接数据库
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // 自动迁移模式
    db.AutoMigrate(&MigrationModel{})

    return &DatabaseMigrationService{db: db}
}

// MigrationModel 是一个示例迁移模型
type MigrationModel struct {
    ID   uint   `gorm:"primaryKey"`
    Name string
}

// Migrate 执行数据库迁移
func (s *DatabaseMigrationService) Migrate(c *gin.Context) {
    // 执行迁移
    if err := s.db.AutoMigrate(&MigrationModel{}); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to migrate database",
            "message": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Database migrated successfully",
    })
}

func main() {
    r := gin.Default()

    // 创建数据库迁移服务
    service := NewDatabaseMigrationService()

    // 注册迁移路由
    r.POST("/migrate", service.Migrate)

    // 启动服务
    if err := r.Run(":8080"); err != nil {
        log.Fatal("Failed to start server: ", err)
    }
}
