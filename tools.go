package utils

import (
	"crypto/md5"
	"fmt"
	"github.com/google/uuid"
	"time"
)

const (
	SECRET = "Rockx"
)

// MD5 使用md5对数据进行加密
func MD5(str string) string {
	md5str := fmt.Sprintf("%x", md5.Sum(append([]byte(str), []byte(SECRET)...)))
	return md5str
}

// SwitchTimeStampToStr 将传入的时间戳转为时间
func SwitchTimeStampToStr(timeStamp int64) string {
	t := time.Unix(timeStamp, 0)
	return t.Format("2006-01-02 15:04:05")
}

// 生成uuid
func Uid() string {
	return uuid.Must(uuid.NewRandom()).String()
}
