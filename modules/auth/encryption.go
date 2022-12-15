package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
)

func EncryptStringWithRSA_OAEP(publickey rsa.PublicKey, toEncryptstring string) string {
	// to label the cipher, so that it can't be used for something else
	label := []byte("for Auth")
	// source of entropy for randomizing the encrypt function
	rng := rand.Reader
	// encrypt the message
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rng, &publickey, []byte(toEncryptstring), label)
	if err != nil {
		panic(err)
	}
	// []byte to string
	return base64.StdEncoding.EncodeToString(ciphertext)
}

func DecryptStringWithRSA_OAEP(privatekey rsa.PrivateKey, toDecryptstring string) string {
	// string to []byte
	cipherText, _ := base64.StdEncoding.DecodeString(toDecryptstring)
	// to label the cipher, so that it can't be used for something else
	label := []byte("for Auth")
	// source of entropy for randomizing the encrypt function
	rng := rand.Reader
	// decrypt the message
	Decryptedstring, err := rsa.DecryptOAEP(sha256.New(), rng, &privatekey, cipherText, label)
	if err != nil {
		panic(err)
	}
	return string(Decryptedstring)
}

// Create rsa Private and Public key pair in memory everytime the server starts
// TODO: Make Database to store persistently

// those keys arent the same across files
// why

// Keys
var ServerPrivateKey, _ = rsa.GenerateKey(rand.Reader, 2048)
var ServerPublicKey = ServerPrivateKey.PublicKey

/*func ServerKeypair() (*rsa.PrivateKey, rsa.PublicKey) {
	PrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	PublicKey := PrivateKey.PublicKey

	return PrivateKey, PublicKey
}*/
