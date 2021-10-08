package main

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func GetPathInt64(c *gin.Context, name string) (int64, error) {
	val := c.Params.ByName(name)
	if val == "" {
		return 0, errors.New(name + " path parameter value is empty or not specified")
	}
	i, err := strconv.Atoi(val)
	return int64(i), err
}

func shaSumStr(fileName string, startByte int64, endByte int64) string {
	h := sha256.New()
	pwd, _ := os.Getwd()
	filePath := strings.Join([]string{pwd, fileName}, "/")
	f, err := os.Open(filePath)
	stat, _ := f.Stat()
	if err != nil {
		return "file not found"
	}
	defer f.Close()

	bytesToIterate := endByte - startByte
	if stat.Size() < bytesToIterate {
		bytesToIterate = stat.Size()
	}

	_, err = f.Seek(startByte, 0)
	if err != nil {
		return "file cannot be seeked there"
	}
	var bufSize int64 = 1
	buf := make([]byte, bufSize)

	var i int64 = 0
	for i = 0; i < bytesToIterate; i++ {
		_, err := f.Read(buf)
		h.Write(buf)
		if err == io.EOF {
			break
		}
	}
	return hex.EncodeToString(h.Sum(nil))
}

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
	router.Run("0.0.0.0:8090")
}
