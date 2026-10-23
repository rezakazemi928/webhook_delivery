package httpapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewRouter registers HTTP routes and shared middleware.
func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	// Enable trusted proxies explicitly when deploying behind a reverse proxy.
	_ = router.SetTrustedProxies(nil)

	// Liveness only: database availability is checked during startup.
	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Register application endpoints here as their handlers are implemented.
	return router
}
