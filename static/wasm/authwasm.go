package main

import (
	"bytes"
	"github.com/gowebapi/webapi"
	"github.com/gowebapi/webapi/html"
	"github.com/gowebapi/webapi/html/htmlevent"
	"gopps/modules/auth"
	"gopps/modules/utils"
	"net/http"
)

// this is the code in encrypt.wasm please do not execute, this is just so you can see the code
// if you know what you're doing feel free to edit this file and compile it like you need

func main() {
	// get button and input
	submitbuttonelement := webapi.GetWindow().Document().GetElementById("submit")
	adminnameinputelement := webapi.GetWindow().Document().GetElementById("adminname")
	adminpassinputelement := webapi.GetWindow().Document().GetElementById("adminpass")
	// something button
	button := html.HTMLButtonElementFromWrapper(submitbuttonelement)
	// something input
	adminnameinput := html.HTMLInputElementFromWrapper(adminnameinputelement)
	adminpassinput := html.HTMLInputElementFromWrapper(adminpassinputelement)
	// get the input value
	adminnameinput.Value()
	adminpassinput.Value()

	// on button press post adminpass and adminname to adminboard
	buttoncallback := func(event *htmlevent.MouseEvent, currentTarget *html.HTMLElement) {
		encryptpass := auth.EncryptStringWithRSA_OAEP(auth.ServerPublicKey, adminpassinput.Value())
		encryptname := auth.EncryptStringWithRSA_OAEP(auth.ServerPublicKey, adminnameinput.Value())
		http.Post("http://"+utils.IPaddress()+"/adminboard", "application/x-www-form-urlencoded",
			bytes.NewBufferString("adminname="+encryptname+"&adminpass="+encryptpass))
	}
	jsFunction := button.SetOnClick(buttoncallback)
	_ = jsFunction
	// prevent the program to terminate
	c := make(chan struct{}, 0)
	<-c
}
