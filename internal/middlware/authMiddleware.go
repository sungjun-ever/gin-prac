package middlware

import "github.com/gin-gonic/gin"

func DummyAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		if token != "secret-token" {
			c.JSON(401, gin.H{"message": "인증 실패"})
			c.Abort()
			return
		}

		c.Next()
	}
}
