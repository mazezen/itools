package memerycache

import (
	"encoding/json"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	B  = 1 << (iota * 10) // 1
	KB                    // 1024
	MB                    // 1048576
	GB                    // 1073741824
	TB                    // 1099511627776
	PB                    // 1125899906842624
)

func CovertSize(size string, defaultMemorySize int64) (int64, string) {
	re, _ := regexp.Compile("[0-9]+")
	b := string(re.ReplaceAll([]byte(size), []byte("")))
	n, _ := strconv.ParseInt(strings.Replace(size, b, "", 1), 10, 64)
	b = strings.ToUpper(b)
	var byteN int64 = 0
	switch b {
	case "B":
		byteN = n
	case "KB":
		byteN = n * KB
	case "MB":
		byteN = n * MB
	case "GB":
		byteN = n * GB
	case "TB":
		byteN = n * TB
	case "PB":
		byteN = n * PB
	default:
		n = 0
	}
	if n == 0 {
		log.Println("size 仅支持 B、KB、MB、GB、TB、PB")
		n = defaultMemorySize
		byteN = n * MB
		b = "MB"
	}
	s := strconv.FormatInt(n, 10) + b
	return byteN, s
}

func CalculateSize(val interface{}) int64 {
	bytes, _ := json.Marshal(val)
	size := int64(len(bytes))
	return size
}
