package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"yudegaki.github.com/rewrite-judge/internal/config"
	"yudegaki.github.com/rewrite-judge/internal/db"
	middleware "yudegaki.github.com/rewrite-judge/internal/middlewares"
)

func main() {
	// Initialize the database
	db.InitDB()

	// Initialize the router
	r := gin.Default()
	r.Use(middleware.Transaction())
	middleware.Router(r)

	r.Run(fmt.Sprintf(":%d", config.APP_PORT))
}
