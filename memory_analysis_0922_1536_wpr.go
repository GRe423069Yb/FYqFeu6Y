// 代码生成时间: 2025-09-22 15:36:38
package main

import (
    "fmt"
    "net/http"
    "runtime"
    "github.com/gin-gonic/gin"
)

// MemoryUsage provides a handler function that returns memory usage statistics.
// It includes error handling and uses Gin middleware for logging.
func MemoryUsage(c *gin.Context) {
    // Get the memory usage statistics
    var m runtime.MemStats
    runtime.ReadMemStats(&m)

    // Calculate the memory usage percentage
    memoryUsage := float64(m.Alloc) / float64(m.Sys) * 100

    // Prepare the response data
    data := struct {
        Alloc       uint64  \"json:\"alloc\"\"
        Sys         uint64  \"