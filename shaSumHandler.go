package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func shaSumHandler(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	filename := queryParams["filename"][0]

	startByte, err1 := strconv.Atoi(queryParams["startbyte"][0])
	endByte, err2 := strconv.Atoi(queryParams["endbyte"][0])
	if err1 != nil || err2 != nil || filename == "" {
		c.IndentedJSON(http.StatusExpectationFailed, "Check parameters, something is missing or byte offsets are not integers")
	}
	endByte64 := int64(endByte)
	startByte64 := int64(startByte)
	sha := shaSumStr(filename, startByte64, endByte64)
	c.Data(200, "text; charset=utf-8", []byte(sha))
}
