package routes

import (
    "github.com/gin-gonic/gin"
    "url-shortener/internal/handlers"
)

func RegisterRoutes(r *gin.Engine) {
    r.POST("/shorten", handlers.ShortenURL)
    r.GET("/u/:code", handlers.Redirect)
    r.GET("/stats/:code", handlers.GetStats)
}
