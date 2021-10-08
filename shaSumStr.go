package main

import (
	"crypto/sha256"
	"encoding/hex"
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

	if stat.Size() < (endByte + startByte) {
		endByte = stat.Size()
		println("too big, adjusting")
	}
	bytesToRead := endByte - startByte

	_, err = f.Seek(startByte, 0)
	if err != nil {
		return "file cannot be seeked there"
	}
	buf := make([]byte, bytesToRead)
	_, err = f.Read(buf)
	h.Write(buf)

	return hex.EncodeToString(h.Sum(nil))
}
