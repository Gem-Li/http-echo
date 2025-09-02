// Author: f210 (508699929@qq.com)
// Date: 2025/09/02
// Description router.go:

package main

import "github.com/gin-gonic/gin"

func RegistRoute(r *gin.Engine) {
	r.GET("/health", Health)
	r.GET("/", Echo)
}
