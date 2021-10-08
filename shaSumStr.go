package main

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"os"
	"strings"
)

func shaSumStr(fileName string, startByte int64, endByte int64) string {
	h := sha256.New()
	pwd, _ := os.Getwd()
	filePath := strings.Join([]string{pwd, fileName}, "/")
	f, err := os.Open(filePath)
	stat, _ := f.Stat()
	if err != nil {
		return "file not found"
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal("Error closing file")
		}
	}(f)

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
