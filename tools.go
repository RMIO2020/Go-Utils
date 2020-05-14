package utils

import (
	"crypto/md5"
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

// MD5 使用md5对数据进行加密
func MD5(str string) string {
	md5str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5str
}

// SwitchTimeStampToStr 将传入的时间戳转为时间
func SwitchTimeStampToStr(timeStamp int64) string {
	t := time.Unix(timeStamp, 0)
	return t.Format("2006-01-02 15:04:05")
}

// Uid 生成uuid
func Uid() string {
	return uuid.Must(uuid.NewRandom()).String()
}

var defaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// RandomString 随机字符串
func RandomString(n int, allowedChars ...[]rune) string {
	var letters []rune
	if len(allowedChars) == 0 {
		letters = defaultLetters
	} else {
		letters = allowedChars[0]
	}

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// 生成订单id
func CreateOrderSn() string {
	return time.Now().Format("20060102150405") + RandomString(6)
}
