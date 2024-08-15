package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

func main() {
	str := "12345"
	fmt.Printf("MD5(%s): %s\n", str, MD5(str))

	fmt.Printf("get now : %s\n", getTimeStr())

	fmt.Printf("get timestamp %d\n", getTimeInt())
}

func MD5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

func getTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func getTimeInt() int64 {
	return time.Now().Unix()
}
