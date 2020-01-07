package main

import "github.com/gin-gonic/gin"

// RunEngine 開始
func RunEngine() {
	r := gin.Default()

	r.GET("/begin", fakeLogin)
	r.POST("/check", checkAns)

	r.Run(":8000")
}
