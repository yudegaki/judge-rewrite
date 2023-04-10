package middleware

import (
	"github.com/gin-gonic/gin"
	"yudegaki.github.com/rewrite-judge/internal/db"
)

func Transaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := db.DB.Begin()
		defer func() {
			if 400 <= c.Writer.Status() {
				db.Rollback()
				return
			}
			db.Commit()
		}()
		c.Set("db", db)
		c.Next()
	}
}
