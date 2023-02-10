package gkit

import (
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetServerAddrs(addrs string) []string {
	s := make([]string, 0)

	for _, v := range strings.Split(addrs, ",") {
		trimed := strings.Trim(v, " ")
		if len(trimed) > 0 {
			s = append(s, trimed)
		}
	}

	return s
}

func Join(strs ...string) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(str)
	}
	return sb.String()
}

func GetPageNumber(page string) int {
	var p int = 0

	if strings.Compare(page, "") != 0 {
		i, err := strconv.Atoi(page)
		if err == nil && i >= 1 {
			p = i - 1
		}
	}

	return p
}

func ConstructKey(prefix, key string) string {
	return prefix + ":" + key
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ResponseError(c *gin.Context, httpCode int, err error) {
	message := err.Error()
	c.JSON(httpCode, ErrResponse{
		Success: false,
		Error:   message,
	})
}
