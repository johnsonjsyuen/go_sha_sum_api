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

	bytesToRead := endByte - startByte
	if stat.Size() < (bytesToRead+startByte)-1 {
		bytesToRead = stat.Size()
		println("too big, adjusting")
	}

	_, err = f.Seek(startByte, 0)
	if err != nil {
		return "file cannot be seeked there"
	}
	buf := make([]byte, bytesToRead)
	_, err = f.Read(buf)
	h.Write(buf)

	return hex.EncodeToString(h.Sum(nil))
}
