package middlewares

import (
	"github.com/gin-gonic/gin"
	"yudegaki.github.com/rewrite-judge/internal/controllers"
)

func Router(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World!")
	})

	authenticated := r.Group("/api", isAuthenticated())
	{
		authenticated.GET("/users", controllers.GetAllUsers)
		// authenticated.GET("/users/:id", controllers.GetUser)
	}

	r.POST("/register", SignUp)
	r.POST("/login", SignIn)
	r.GET("/signout", SignOut)

}
