// Author: f210 (508699929@qq.com)
// Date: 2025/09/02
// Description handler.go:

package main

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	c.String(200, "ok")
}

func Echo(c *gin.Context) {
	osname, err := os.Hostname()
	if err != nil {
		osname = "unknown"
	}

	var cookies []string

	for _, cookie := range c.Request.Cookies() {
		cookies = append(cookies, cookie.Name)
	}

	headers := make(map[string]string)
	for name, value := range c.Request.Header {
		headers[name] = strings.Join(value, "|")
	}

	parts := strings.Split(c.Request.RemoteAddr, ":")

	body := map[string]interface{}{
		"tcp": map[string]string{
			"ip":   strings.Join(parts[:(len(parts)-1)], ":"),
			"port": parts[len(parts)-1],
		},
		"os":       osname,
		"header":   headers,
		"protocol": c.Request.Proto,
		"host":     c.Request.Host,
		"method":   c.Request.Method,
		"path":     c.Request.URL.Path,
		"cookie":   cookies,
		"query":    c.Request.URL.RawQuery,
	}

	c.JSON(200, body)
}
