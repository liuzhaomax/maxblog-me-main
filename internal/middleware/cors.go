package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"maxblog-me-main/internal/core"
	"net/http"
)

func Cors() gin.HandlerFunc {
	var corsWhiteList = []string{
		fmt.Sprintf("http://%s", core.GetUpstreamAddr()),
		fmt.Sprintf("http://%s", "127.0.0.1:3000"), // TODO for st env
	}
	return func(ctx *gin.Context) {
		if core.In(corsWhiteList, ctx.Request.Header.Get("Origin")) {
			ctx.Header("Access-Control-Allow-Origin", ctx.Request.Header.Get("Origin"))
		}
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, AccessToken, X-CSRF-Token, Authorization, Token, Set-Cookie, X-Requested-With, Access-Control-Allow-Origin, Content-Security-Policy")
		ctx.Header("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, DELETE, OPTIONS")
		ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		// let go all options request
		method := ctx.Request.Method
		if method == "OPTIONS" {
			ctx.Header("Access-Control-Max-Age", "86400") // one day
			ctx.AbortWithStatus(http.StatusNoContent)
		}
		ctx.Next()
	}
}
