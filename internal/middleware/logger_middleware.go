package middleware

import (
	"time"

	"user-age-api/internal/logger"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)




// RequestLogger logs basic request details
func RequestLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {

		start := time.Now()
		err := c.Next()
		duration := time.Since(start)

		logger.Log.Info("http_request",
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Int("status", c.Response().StatusCode()),
			zap.Duration("duration", duration),
		)

		return err
	}
}
