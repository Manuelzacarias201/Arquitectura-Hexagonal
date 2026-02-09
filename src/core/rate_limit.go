package core

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// AuthRateLimitMiddleware limita intentos por IP (p. ej. 10 por minuto) para login/register
func AuthRateLimitMiddleware(requestsPerMinute int) gin.HandlerFunc {
	type entry struct {
		limiter *rate.Limiter
		last    time.Time
	}
	var (
		mu      sync.Mutex
		clients = make(map[string]*entry)
	)
	return func(c *gin.Context) {
		ip := c.ClientIP()
		mu.Lock()
		if e, ok := clients[ip]; ok {
			e.last = time.Now()
			if !e.limiter.Allow() {
				mu.Unlock()
				c.JSON(http.StatusTooManyRequests, gin.H{
					"error": "Demasiados intentos. Espera un momento e inténtalo de nuevo.",
					"code":  "RATE_LIMIT_EXCEEDED",
				})
				c.Abort()
				return
			}
		} else {
			clients[ip] = &entry{
				limiter: rate.NewLimiter(rate.Every(time.Minute/time.Duration(requestsPerMinute)), requestsPerMinute),
				last:    time.Now(),
			}
			clients[ip].limiter.Allow()
		}
		mu.Unlock()
		c.Next()
	}
}
