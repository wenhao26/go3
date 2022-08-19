package router

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"coinsky_go_project/app/middleware"
)

func InitRouter()  {
	r := gin.Default()
	gin.ForceConsoleColor()
	gin.SetMode(gin.DebugMode)

	r.NoRoute()
	r.Use(middleware.Cors(), middleware.RateLimiter(time.Second, 100, 100))

	// TODO
	r.GET("hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello")
	})

	// r.Run()
	server := &http.Server{
		Addr: ":8888",
		Handler: r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	_ = server.ListenAndServe()
}
