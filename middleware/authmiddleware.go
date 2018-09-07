package middleware

import (
	"../utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		seed := utils.GetEnv("GLOBAL_SEED", utils.GLOBAL_SEED)
		key := []byte(seed)
		jwt_token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJYcm9sZSI6ImFkbWluIiwiWHVzZXIiOiJoZWxpb3MifQ.lNA0CQiMmdF40rmwEpKFBmzTUYfhtaIwQiNuPNdIKc0"
		if res, err := utils.ParseJWTString(jwt_token, key); err != nil {
			fmt.Println(res["Xrole"])
			c.Next()
			return
		} else {
			fmt.Println(err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
	}
}
