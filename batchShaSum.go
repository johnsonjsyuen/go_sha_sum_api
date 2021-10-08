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
	pwd, _ := os.Getwd()
	filePath := strings.Join([]string{pwd, fileName}, "/")
	f, err := os.Open(filePath)
	stat, _ := f.Stat()
	fileSize := stat.Size()

	iterations := fileSize / partSize
	if fileSize%partSize != 0 {
		iterations++
	}
	hashes := make([]string, 0, iterations)

	if err != nil {
		return nil, errors.New("error opening file")
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal("Error closing file")
		}
	}(f)

	var i int64 = 0

	var buf []byte

	// (0,100)(101,200)(201,300)
	for i = 0; i < iterations-1; i++ {
		// We have to manipulate the slice size to read exactly the bytes we want to get correct SHA sum

		if i == 0 {
			buf = make([]byte, partSize)
			//println(partSize)
		} else {
			buf = make([]byte, partSize-1)
			//println(partSize-1)
		}
		if i > 0 {
			f.Seek(1, 1)
		}

		_, _ = f.Read(buf)
		h := sha256.New()
		h.Write(buf)
		hashes = append(hashes, hex.EncodeToString(h.Sum(nil)))
		if err == io.EOF {
			break
		}
	}

	remainder := fileSize % partSize
	lastHashStart := (fileSize/partSize)*partSize + 1
	buf = make([]byte, remainder)
	f.Seek(lastHashStart, 0)
	_, _ = f.Read(buf)
	h := sha256.New()
	h.Write(buf)
	hashes = append(hashes, hex.EncodeToString(h.Sum(nil)))

	return hashes, nil
}
