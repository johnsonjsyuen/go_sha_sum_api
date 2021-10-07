package main

import (
	"crypto/sha256"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func sha_sum(c *gin.Context) {
	h := sha256.New()
	h.Write([]byte("hello world\n"))
	c.IndentedJSON(http.StatusOK, h.Sum(nil))
}

func headers(c *gin.Context) {

	for name, headers := range c.Request.Header {
		for _, h := range headers {
			fmt.Fprintf(c.Writer, "%v: %v\n", name, h)
		}
	}
}

func main() {
	router := gin.Default()
	router.GET("/", sha_sum)
	router.GET("/headers", headers)
	router.Run("localhost:8090")
}
