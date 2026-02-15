package handlers

import (
    "math/rand"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "url-shortener/internal/db"
    "url-shortener/internal/models"
)

type ShortenRequest struct {
    URL string `json:"url" binding:"required"`
}

func init() {
    rand.Seed(time.Now().UnixNano())
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateCode(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

func ShortenURL(c *gin.Context) {
    var req ShortenRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "url is required"})
        return
    }

    code := generateCode(6)

    url := models.URL{
        OriginalURL: req.URL,
        ShortCode:   code,
    }

    if err := db.DB.Create(&url).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save url"})
        return
    }

    shortURL := c.Request.Host + "/u/" + code

    c.JSON(http.StatusCreated, gin.H{
        "short_url":   "http://" + shortURL,
        "short_code":  code,
        "original_url": req.URL,
    })
}

func Redirect(c *gin.Context) {
    code := c.Param("code")

    var url models.URL
    if err := db.DB.Where("short_code = ?", code).First(&url).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "short url not found"})
        return
    }

    db.DB.Model(&url).UpdateColumn("clicks", url.Clicks+1)

    c.Redirect(http.StatusFound, url.OriginalURL)
}

func GetStats(c *gin.Context) {
    code := c.Param("code")

    var url models.URL
    if err := db.DB.Where("short_code = ?", code).First(&url).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "short url not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "original_url": url.OriginalURL,
        "short_code":   url.ShortCode,
        "clicks":       url.Clicks,
        "created_at":   url.CreatedAt,
    })
}
