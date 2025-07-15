package middleware

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "strings"
    "time"
    
    "github.com/brokingSapphire/SapphireICICI/internal/logger"
    "github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        
        var requestBody []byte
        if c.Request.Body != nil {
            requestBody, _ = io.ReadAll(c.Request.Body)
            c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
        }
        
        blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
        c.Writer = blw
        
        // Process request
        c.Next()
        
        // Calculate duration
        duration := time.Since(start)
    
        displayPath := c.Request.URL.Path
        if strings.HasPrefix(displayPath, "/api/v1") {
            displayPath = strings.TrimPrefix(displayPath, "/api/v1")
            if displayPath == "" {
                displayPath = "/"
            }
        }
        
        // Format request body for logging
        var requestBodyStr string
        if len(requestBody) > 0 && len(requestBody) < 1024 {
            var requestJSON interface{}
            if json.Unmarshal(requestBody, &requestJSON) == nil {
                if reqBytes, err := json.Marshal(requestJSON); err == nil {
                    requestBodyStr = string(reqBytes)
                }
            }
        }
        
        // Format response body for logging
        var responseBodyStr string
        if blw.body.Len() < 1024 && blw.body.Len() > 0 {
            var responseJSON interface{}
            if json.Unmarshal(blw.body.Bytes(), &responseJSON) == nil {
                if respBytes, err := json.Marshal(responseJSON); err == nil {
                    responseBodyStr = string(respBytes)
                }
            }
        }
        
        logMessage := fmt.Sprintf("%s %s %d %.3f ms - %d",
            c.Request.Method,
            displayPath,
            c.Writer.Status(),
            float64(duration.Nanoseconds())/1e6,
            c.Writer.Size(),
        )
        
        // Add request body if present
        if requestBodyStr != "" {
            logMessage += " " + requestBodyStr
        }
        
        // Add response body if present
        if responseBodyStr != "" {
            logMessage += " " + responseBodyStr
        }
      
        if c.Writer.Status() >= 400 {
            logger.Error(logMessage)
        } else {
            logger.Info(logMessage)
        }
    }
}

type bodyLogWriter struct {
    gin.ResponseWriter
    body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
    w.body.Write(b)
    return w.ResponseWriter.Write(b)
}