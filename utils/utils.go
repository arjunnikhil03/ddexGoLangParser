package utils

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

//Message is used to create success or failure messages
func Message(status bool, message string, statusCode int) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message, "statusCode": statusCode}
}

//Respond is used to reply to HTTP request
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")

	if _, ok := data["statusCode"]; ok {
		w.WriteHeader(data["statusCode"].(int))
	}
	json.NewEncoder(w).Encode(data)
}

//RandomString Returns random string with specified length
func RandomString(length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

//Strip ..
func Strip(s string) string {
	var result strings.Builder
	for i := 0; i < len(s); i++ {
		b := s[i]
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			('0' <= b && b <= '9') {
			result.WriteByte(b)
		}
	}
	return result.String()
}

//StrReplace
func StrReplace(search, replace, subject string) string {
	res := strings.Replace(
		subject,
		search,
		replace,
		-1,
	)

	return res
}

//EmptyString ...
func EmptyString(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

//MicroTime ...
func MicroTime() float64 {
	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc)
	micSeconds := float64(now.Nanosecond()) / 1000000000
	return float64(now.Unix()) + micSeconds
}

//InArray ...
func InArray(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}
