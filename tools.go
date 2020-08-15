package utils

import (
	"crypto/md5"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/skip2/go-qrcode"
	"math/rand"
	"net/url"
	"strings"
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

// CreateOrderSn 生成订单id
func CreateOrderSn() string {
	return time.Now().Format("20060102150405") + RandomString(6)
}

// CreateQrCode 生成二维码base64
func CreateQrCodeBase64(username string, secret string, issuer string) (encoded string, err error) {
	username = url.QueryEscape(username)
	issuer = url.QueryEscape(issuer)
	q, err := qrcode.New(fmt.Sprintf("otpauth://totp/%v-%v?secret=%v&issuer=%v", username, issuer, secret, issuer), qrcode.Medium)
	if err != nil {
		return
	}
	// Optionally, disable the QR Code border.
	q.DisableBorder = true
	png, err := q.PNG(256)
	if err != nil {
		return
	}
	encoded = fmt.Sprintf("data:image/jpg;base64,%v", base64.StdEncoding.EncodeToString(png))
	return
}

// PlusDeal 加号处理
func PlusDeal(in string) string {
	in = strings.Trim(in, " ")
	return "+" + strings.TrimLeft(in, "+")
}

// generateSalt 获取盐值
func generateSalt(cost int) (salt string, err error) {
	if cost < 4 || cost > 31 {
		err = errors.New("cost must be between 4 and 31")
		return
	}
	// Get a 20-byte random string
	randValue := RandomString(20)
	// Form the prefix that specifies Blowfish (bcrypt) algorithm and cost parameter.
	salt = fmt.Sprintf("$2y$%02d$", cost)
	// Append the random salt data in the required base64 format.
	encodeRand := base64.StdEncoding.EncodeToString([]byte(randValue))
	salt = fmt.Sprintf("%v%v", salt, strings.Replace(encodeRand[:22], "+", ".", -1))
	return
}
