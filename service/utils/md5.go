package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// md5 加密
func MD5(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}

// 字符串翻转
func ReverseString(str string) string {
	strArr := []rune(str)
	for i := 0; i < len(strArr)/2; i++ {
		strArr[len(strArr)-1-i], strArr[i] = strArr[i], strArr[len(strArr)-1-i]
	}
	return string(strArr)
}
