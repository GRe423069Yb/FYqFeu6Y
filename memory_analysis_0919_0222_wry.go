// 代码生成时间: 2025-09-19 02:22:41
This handler provides an endpoint to analyze the memory usage of the current
application within a specified interval and error handling.
*/

package main

import (
    "fmt"
    "net/http"
    "runtime"
    "time"

    "github.com/gin-gonic/gin"
)

// MemoryStats represents memory usage statistics.
type MemoryStats struct {
    // Alloc is bytes of allocated heap objects.
    Alloc uint64 `json:"alloc"`
    // TotalAlloc is bytes of allocated heap objects.
    TotalAlloc uint64 `json:"total_alloc"`
    // Sys is the total bytes of memory obtained from the OS.
    Sys uint64 `json:"sys"`
    // NumGC is the number of garbage collection cycles that have run.
    NumGC uint32 `json:"num_gc"`
}

// getMemoryUsage retrieves the current memory usage statistics.
func getMemoryUsage() MemoryStats {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    return MemoryStats{
        Alloc:      m.Alloc,
        TotalAlloc: m.TotalAlloc,
        Sys:        m.Sys,
        NumGC:      m.NumGC,
    }
}

// memoryAnalysisHandler handles requests to analyze memory usage.
func memoryAnalysisHandler(c *gin.Context) {
    // Analyze memory usage
    stats := getMemoryUsage()

    // Send the memory usage statistics in JSON format.
    c.JSON(http.StatusOK, stats)
}

// main function to setup Gin router and middleware.
func main() {
    r := gin.Default()

    // Register the memory analysis handler with a route.
    r.GET("/memory", memoryAnalysisHandler)

    // Start the server.
    r.Run() // listening and serving on 0.0.0.0:8080
}
