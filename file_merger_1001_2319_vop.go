// 代码生成时间: 2025-10-01 23:19:51
package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
    "path/filepath"

    "github.com/gin-gonic/gin"
)

// FileMergerHandler 处理文件分割合并的请求
func FileMergerHandler(c *gin.Context) {
    src := c.Query("src") // 源目录或文件
    dest := c.Query("dest") // 目标文件
    
    // 参数检查
    if src == "" || dest == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "源目录和目标文件不能为空"
        })
        return
    }

    // 检查源路径是否存在
    if _, err := os.Stat(src); os.IsNotExist(err) {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "源文件不存在"
        })
        return
    }

    // 打开目标文件准备写入
    file, err := os.Create(dest)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "无法创建目标文件"
        })
        return
    }
    defer file.Close()

    // 读取源目录中的文件并合并
    files, err := os.ReadDir(src)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "读取源目录失败"
        })
        return
    }
    for _, f := range files {
        if !f.IsDir() {
            srcFile, err := os.Open(filepath.Join(src, f.Name()))
            if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{
                    "error": "打开源文件失败"
                })
                return
            }
            defer srcFile.Close()
            
            // 将文件内容写入目标文件
            if _, err := io.Copy(file, srcFile); err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{
                    "error": "合并文件失败"
                })
                return
            }
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "文件合并成功"
    })
}

func main() {
    router := gin.Default()

    // 注册文件合并处理器
    router.GET("/merge", FileMergerHandler)

    // 启动服务
    log.Fatal(router.Run(":8080"))
}
