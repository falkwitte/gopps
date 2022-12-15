package auth

import (
	"crypto/subtle"
)

func Authvalidator(formvalueusername, formvaluepassword, username, password string) bool {
	// use constanttimecompare to prevent timing attacks
	if subtle.ConstantTimeCompare([]byte(formvalueusername), []byte(username)) == 1 &&
		subtle.ConstantTimeCompare([]byte(formvaluepassword), []byte(password)) == 1 {
		return true
	}
	return false
}
