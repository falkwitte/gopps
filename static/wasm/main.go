package main

import (
	"bytes"
	"fmt"
	"gopps/modules/auth"
	"net/http"
)

// JUST FOR TESTING
func main() {

	encryptpass := auth.EncryptStringWithRSA_OAEP(auth.ServerPublicKey, "admin")
	encryptname := auth.EncryptStringWithRSA_OAEP(auth.ServerPublicKey, "admin")

	fmt.Printf("adminname: ", encryptname)
	fmt.Println("\n")
	response, _ := http.Post("http://localhost:8080/adminboard", "application/x-www-form-urlencoded", bytes.NewBufferString("adminname="+encryptname+"&adminpass="+encryptpass))
	fmt.Printf("Response: ", response)
}
