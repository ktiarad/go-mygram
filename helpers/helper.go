package helpers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetContentType(ctx *gin.Context) string {
	return ctx.Request.Header.Get("Content-Type")
}

func HashPass(p string) string {
	salt := 8

	password := []byte(p)

	hash, _ := bcrypt.GenerateFromPassword(password, salt)

	return string(hash)
}

func ComparePass(h, p []byte) bool {
	hash, pass := []byte(h), []byte(p)

	err := bcrypt.CompareHashAndPassword(hash, pass)

	return err == nil
}
