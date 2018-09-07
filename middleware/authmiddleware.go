package middleware

import (
	"../utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	_ "reflect"
)

func AuthMiddleWare(roles_acc map[string]bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		seed := utils.GetEnv("GLOBAL_SEED", utils.GLOBAL_SEED)
		key := []byte(seed)
		jwt_token := c.Request.Header.Get("Xauth")
		if res, err := utils.ParseJWTString(jwt_token, key); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		} else {
			//fmt.Println(res["Xrole"])
			//FIXME, need to move to db
			if !roles_acc[res["Xrole"]] {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				c.Abort()
				return
			} else {
				c.Header("Xuser", res["Xuser"])
				c.Header("Xrole", res["Xrole"])
				c.Next()
				return
			}
		}
	}
}
