package api

import (
	"../utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ProxyRedirect(proxy_key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Redirect(http.StatusUseProxy, utils.GetEnv(proxy_key, utils.API_PROXY_MAP[proxy_key])+c.Request.URL.String())
	}
}
