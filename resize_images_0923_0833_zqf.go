// 代码生成时间: 2025-09-23 08:33:47
package main

import (
    "image"
    "image/jpeg"
    "image/png"
    "net/http"
    "os"
    "strconv"
    
    "github.com/gin-gonic/gin"
)

// ImageResizer is the handler function that resizes images based on the provided width and height.
# 增强安全性
// It expects a JSON payload with a slice of image paths and the desired dimensions.
func ImageResizer(c *gin.Context) {
    // Define error variable
    var err error
    defer func() {
        if r := recover(); r != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Internal Server Error",
            })
        }
    }()

    // Bind JSON payload to struct
    var payload struct {
        Images []string `json:"images"`
        Width  int      `json:"width"`
        Height int      `json:"height"`
    }
    if err = c.ShouldBindJSON(&payload); err != nil {
# NOTE: 重要实现细节
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    // Check if width and height are positive
    if payload.Width <= 0 || payload.Height <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Width and height must be greater than zero.",
        })
# 优化算法效率
        return
    }

    // Process each image
    for _, path := range payload.Images {
# TODO: 优化性能
        file, err := os.Open(path)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to open image file.",
            })
            return
        }
        defer file.Close()

        img, _, err := image.Decode(file)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
# TODO: 优化性能
                "error": "Failed to decode image.",
            })
            return
        }

        // Create a new image with the desired width and height
        imgResized := resizeImage(img, payload.Width, payload.Height)
# 添加错误处理

        // Save the resized image to a file
        outputPath := path + "_resized." + getExtension(path)
        out, err := os.Create(outputPath)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
# FIXME: 处理边界情况
                "error": "Failed to create resized image file.",
            })
            return
        }
# 优化算法效率
        defer out.Close()

        // Write the new image to the file
        switch getExtension(path) {
        case "jpg", "jpeg":
# 改进用户体验
            err = jpeg.Encode(out, imgResized, nil)
        case "png":
            err = png.Encode(out, imgResized)
        default:
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Unsupported image format.",
            })
            return
        }
# 扩展功能模块
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to save resized image.",
            })
            return
        }
# 增强安全性
    }

    // Return success response
    c.JSON(http.StatusOK, gin.H{
        "message": "Images resized successfully.",
    })
}

// resizeImage resizes the given image to the specified width and height.
// It maintains the aspect ratio of the original image.
func resizeImage(img image.Image, width, height int) image.Image {
    // Calculate the ratio of the new size to the original size
    ratio := float64(width) / float64(img.Bounds().Dx())
# 增强安全性
    if float64(height)/float64(img.Bounds().Dy()) < ratio {
# 优化算法效率
        ratio = float64(height) / float64(img.Bounds().Dy())
    }
    newWidth := int(float64(img.Bounds().Dx()) * ratio)
    newHeight := int(float64(img.Bounds().Dy()) * ratio)

    // Create a new image with the new size
    imgResized := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

    // Draw the original image onto the new image with the new size
    imgResized = resize.Thumbnail(newWidth, newHeight, img, resize.Lanczos3)

    return imgResized
}

// getExtension returns the file extension of the given path.
func getExtension(path string) string {
    i := strings.LastIndex(path, ".")
    if i == -1 {
        return ""
    }
    return path[i+1:]
}

func main() {
    r := gin.Default()
    r.POST("/resize", ImageResizer)
    r.Run() // listen and serve on 0.0.0.0:8080
# FIXME: 处理边界情况
}