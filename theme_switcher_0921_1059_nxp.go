// 代码生成时间: 2025-09-21 10:59:27
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// ThemeSwitcherHandler 是处理主题切换的处理器
func ThemeSwitcherHandler(c *gin.Context) {
    theme := c.DefaultQuery("theme", "default") // 默认主题为 default
    // 尝试切换主题
    err := SwitchTheme(theme)
    if err != nil {
        // 如果切换主题失败，返回错误信息
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }
    // 成功切换主题
    c.JSON(http.StatusOK, gin.H{
        "message": "Theme switched successfully",
    })
}

// SwitchTheme 尝试切换主题，如果主题无效则返回错误
func SwitchTheme(theme string) error {
    supportedThemes := []string{"default", "dark", "light"}
    for _, t := range supportedThemes {
        if t == theme {
            return nil // 主题有效，返回 nil
        }
    }
    return fmt.Errorf("theme %s is not supported", theme)
}

func main() {
    r := gin.Default() // 创建一个默认的 Gin 路由器
    
    // 添加中间件
    r.Use(gin.Recovery()) // 恢复中间件，用于处理 panic 错误
    r.Use(gin.Logger())   // 日志中间件，用于记录请求日志
    
    // 注册主题切换处理函数
    r.GET("/theme", ThemeSwitcherHandler)
    
    // 启动服务
    r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
