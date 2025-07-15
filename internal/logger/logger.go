package logger

import (
    "os"
    "path/filepath"
    
    "github.com/brokingSapphire/SapphireICICI/internal/env"
    "github.com/sirupsen/logrus"
    "gopkg.in/natefinch/lumberjack.v2"
)

var Log *logrus.Logger


func Init() {
    Log = logrus.New()
    
    // Set log level based on environment
    if env.Env.Env == "production" {
        Log.SetLevel(logrus.InfoLevel)
    } else {
        Log.SetLevel(logrus.DebugLevel)
    }
    
    // Create logs directory if it doesn't exist
    if err := os.MkdirAll("logs", 0755); err != nil {
        Log.Warn("Could not create logs directory:", err)
    }
    
    // Configure formatters and outputs based on environment
    if env.Env.Env == "production" {
        // Production: JSON format with file rotation
        Log.SetFormatter(&logrus.JSONFormatter{
            TimestampFormat: "2006-01-02 15:04:05",
        })
        
        Log.SetOutput(&lumberjack.Logger{
            Filename:   filepath.Join("logs", "combined.log"),
            MaxSize:    5,    // 5MB
            MaxBackups: 5,    // Keep 5 old files
            MaxAge:     28,   // Keep for 28 days
            Compress:   true, // Compress old files
        })
        
    } else {
        // Development: Colored text format to console
        Log.SetFormatter(&logrus.TextFormatter{
            FullTimestamp:   true,
            TimestampFormat: "2006-01-02 15:04:05",
            ForceColors:     true,
        })
        Log.SetOutput(os.Stdout)
    }
    // Add service name to all logs
    Log = Log.WithFields(logrus.Fields{
        "service": "icici-api-service",
    }).Logger
}

// Wrapper functions similar to your logger usage
func Info(args ...interface{}) {
    Log.Info(args...)
}

func Debug(args ...interface{}) {
    Log.Debug(args...)
}

func Error(args ...interface{}) {
    Log.Error(args...)
}

func Warn(args ...interface{}) {
    Log.Warn(args...)
}

func Fatal(args ...interface{}) {
    Log.Fatal(args...)
}

// Structured logging functions
func InfoWithFields(fields logrus.Fields, args ...interface{}) {
    Log.WithFields(fields).Info(args...)
}

func ErrorWithFields(fields logrus.Fields, args ...interface{}) {
    Log.WithFields(fields).Error(args...)
}