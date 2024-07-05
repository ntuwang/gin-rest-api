package middleware

import (
	"gin-rest-api/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func AccessLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		path := ctx.Request.URL.Path
		raw := ctx.Request.URL.RawQuery
		ctx.Next()
		latency := time.Now().Sub(start)
		clientIP := ctx.ClientIP()
		method := ctx.Request.Method
		statusCode := ctx.Writer.Status()
		bodySize := ctx.Writer.Size()
		if raw != "" {
			path = path + "?" + raw
		}
		log.WithCtx(ctx).Info("access",
			zap.String("path", path),
			zap.String("method", method),
			zap.String("remote_addr", ctx.Request.RemoteAddr),
			zap.Int("status_code", statusCode),
			zap.Int("body_size", bodySize),
			zap.String("client_ip", clientIP),
			zap.Duration("latency", latency),
		)
	}
}
