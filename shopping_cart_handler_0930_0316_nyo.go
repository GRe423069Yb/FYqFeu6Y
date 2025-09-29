// 代码生成时间: 2025-09-30 03:16:24
package main

import (
    "crypto/md5"
    "encoding/hex"
    "encoding/json"
    "net/http"
    "strconv"
    "time"

    "github.com/gin-gonic/gin"
)

// ShoppingCart represents the shopping cart data structure
type ShoppingCart struct {
    ID        string    `json:"id"`
    Items     []*Item   `json:"items"`
    CreatedAt time.Time `json:"createdAt"`
}

// Item represents an item in the shopping cart
type Item struct {
    ProductID string  `json:"productID"`
    Quantity  int     `json:"quantity"`
}

func main() {
    r := gin.Default()

    // Initialize a new shopping cart
    cart := &ShoppingCart{
        ID:        generateID(),
        CreatedAt: time.Now(),
    }

    // Example middleware to handle CORS
    r.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Next()
    })

    // GET /cart to retrieve the current shopping cart
    r.GET("/cart", func(c *gin.Context) {
        c.JSON(http.StatusOK, cart)
    })

    // POST /cart to add an item to the shopping cart
    r.POST("/cart", func(c *gin.Context) {
        var item Item
        if err := c.ShouldBindJSON(&item); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": err.Error(),
            })
            return
        }
        cart.Items = append(cart.Items, &item)
        c.JSON(http.StatusOK, cart)
    })

    // POST /cart/:productID/quantity to update the quantity of a product in the shopping cart
    r.POST("/cart/:productID/quantity", func(c *gin.Context) {
        productID := c.Param("productID")
        quantity, err := strconv.Atoi(c.PostForm("quantity"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "Invalid quantity",
            })
            return
        }
        updateQuantity(cart, productID, quantity)
        c.JSON(http.StatusOK, cart)
    })

    // Start the server
    r.Run()
}

// generateID creates a unique ID for the shopping cart
func generateID() string {
    h := md5.New()
    now := time.Now()
    h.Write([]byte(now.String()))
    return hex.EncodeToString(h.Sum(nil))
}

// updateQuantity updates the quantity of a product in the shopping cart
func updateQuantity(cart *ShoppingCart, productID string, quantity int) {
    for i, item := range cart.Items {
        if item.ProductID == productID {
            cart.Items[i].Quantity = quantity
            break
        }
    }
}
