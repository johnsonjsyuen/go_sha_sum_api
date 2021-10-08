package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

func headers(c *gin.Context) {

	for name, headers := range c.Request.Header {
		for _, h := range headers {
			_, _ = fmt.Fprintf(c.Writer, "%v: %v\n", name, h)
		}
	}
}

func main() {
	router := gin.Default()
	router.GET("/", shaSumHandler)
	router.GET("/headers", headers)
	router.GET("/batch", batchShaSumHandler)
	address := "0.0.0.0:8090"
	err := router.Run(address)
	if err != nil {
		log.Fatal(strings.Join([]string{"Problem starting server at ", "0.0.0.0:8090"}, ""))
	}
}
