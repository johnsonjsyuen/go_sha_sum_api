package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func batchShaSumHandler(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	filename := queryParams["filename"][0]

	partSize, err1 := strconv.Atoi(queryParams["partsize"][0])
	if err1 != nil || filename == "" {
		c.IndentedJSON(http.StatusExpectationFailed, "Check parameters, something is missing or byte offsets are not integers")
	}
	partSize64 := int64(partSize)
	hashes, err := batchShaSum(filename, partSize64)
	if err != nil {
		c.IndentedJSON(http.StatusExpectationFailed, err)
	}
	c.IndentedJSON(http.StatusOK, hashes)
}
