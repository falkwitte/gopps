package auth

import (
	"crypto/subtle"
	"fmt"
)

func Authvalidator(formvalueusername, formvaluepassword, username, password string) bool {

	// decrypt passwords before comparison
	fmt.Printf("Formvalueusername: ", formvalueusername)
	fmt.Println("\n")
	decryptformvalueusername := DecryptStringWithRSA_OAEP(*ServerPrivateKey, formvalueusername)
	decryptformvaluepassword := DecryptStringWithRSA_OAEP(*ServerPrivateKey, formvaluepassword)

	// use constanttimecompare to prevent timing attacks
	if subtle.ConstantTimeCompare([]byte(decryptformvalueusername), []byte(username)) == 1 &&
		subtle.ConstantTimeCompare([]byte(decryptformvaluepassword), []byte(password)) == 1 {
		return true
	}
	return false
}
