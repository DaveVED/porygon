package main

import (
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
)

// HealthCheckEndpoints provide readiness and liveness endpoints/checks for the service.
func HealthCheckEndpoints(router *gin.Engine) {
    router.GET("/health/ready", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "UP"})
    })

    router.GET("/health/live", func(c *gin.Context) {
        c.Status(http.StatusNoContent)
    })
}

func run() {
    log.Println("running the service...")
    
    router := gin.Default()
    
    // HealthCheckEndpoints Initialization.
    HealthCheckEndpoints(router)

    if err := router.Run(); err != nil {
        log.Fatalf("failed: to run the service: %v", err)
    }

}

func main() {
    run()
}
