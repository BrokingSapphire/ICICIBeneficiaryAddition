package internal

import (
    "github.com/gin-gonic/gin"
)

// Security adds security headers - similar to helmet middleware
func Security() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Set security headers
        c.Header("X-Frame-Options", "DENY")
        c.Header("X-Content-Type-Options", "nosniff")
        c.Header("X-XSS-Protection", "1; mode=block")
        c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
        c.Header("Referrer-Policy", "strict-origin-when-cross-origin")
        c.Header("Content-Security-Policy", "default-src 'self'")
        
        c.Next()
    }
}