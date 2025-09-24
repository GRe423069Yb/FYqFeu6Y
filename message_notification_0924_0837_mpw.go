// 代码生成时间: 2025-09-24 08:37:10
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "log"
)

// NotificationMessage is the structure that will be sent to the notification system
type NotificationMessage struct {
    Content string `json:"content"`
}

// NotificationHandler handles the incoming message notifications
func NotificationHandler(c *gin.Context) {
    var message NotificationMessage
    if err := c.ShouldBindJSON(&message); err != nil {
        // Bind JSON to NotificationMessage struct and handle errors
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON binding"})
        return
    }

    // Simulate sending the notification message
    log.Println("Sending notification: ", message.Content)

    // Return a success response
    c.JSON(http.StatusOK, gin.H{"message": "Notification sent successfully"})
}

func main() {
    // Create a new Gin router with default middleware: logger and recovery (catches any panic)
    router := gin.Default()

    // Define a route with POST method for sending notification messages
    router.POST("/notify", NotificationHandler)

    // Start the server on port 8080
    router.Run(":8080")
}
