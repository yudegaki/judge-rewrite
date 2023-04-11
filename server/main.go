package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"yudegaki.github.com/rewrite-judge/internal/config"
	"yudegaki.github.com/rewrite-judge/internal/external"
	"yudegaki.github.com/rewrite-judge/internal/middlewares"
)

func main() {
	// Initialize the database
	external.InitDB()

	// Initialize the router
	r := gin.Default()
	r.Use(middlewares.Transaction())
	middlewares.Router(r)

	r.Run(fmt.Sprintf(":%d", config.APP_PORT))
}
