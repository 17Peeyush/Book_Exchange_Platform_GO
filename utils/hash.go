package utils

import (
	"hash/fnv"
	"math"
	"strconv"
)

//const base62Digits = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func UniqueId(link string) string {
	// Create a new FNV-1a hash object
	hashObj := fnv.New32()

	// Write the bytes of the input string to the hash object
	hashObj.Write([]byte(link))
	hashNum := int(math.Abs((float64(hashObj.Sum32()))))
	return strconv.Itoa(hashNum)
}