// +build js,wasm

package main

import (
	"log"
	"syscall/js"
)

func main() {
	log.Println("Hello Gophers!")
	window := js.Global()
	doc := window.Get("document")
	body := doc.Get("body")
	div := doc.Call("createElement", "div")
	div.Set("textContent", "hello!!")
	body.Call("appendChild", div)
	body.Set("onclick",
		js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			div := doc.Call("createElement", "div")
			div.Set("textContent", "click!!")
			body.Call("appendChild", div)
			return nil
		}))
	<-make(chan struct{})
}
