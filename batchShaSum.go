package main

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"log"
	"os"
	"strings"
)

func batchShaSum(fileName string, partSize int64) ([]string, error) {
	hashes := make([]string, 0)
	pwd, _ := os.Getwd()
	filePath := strings.Join([]string{pwd, fileName}, "/")
	f, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New("error opening file")
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal("Error closing file")
		}
	}(f)

	buf := make([]byte, partSize)

	var i int64 = 0
	// (0,100)(101,200)(201,300)
	for i = 0; ; i++ {
		if i == 1 {
			buf = make([]byte, partSize-1)
		}
		if i > 0 {
			f.Seek(1, 1)
		}
		_, err := f.Read(buf)
		h := sha256.New()
		h.Write(buf)
		hashes = append(hashes, hex.EncodeToString(h.Sum(nil)))
		if err == io.EOF {
			break
		}
	}
	return hashes, nil
}
