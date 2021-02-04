package main

import "syscall/js"

func main() {
	// try to get the 'document' property of JS from the global scope
	doc := js.Global().Get("document")
	if !doc.Truthy() { // the JS way of testing for nil
		panic("unable to get 'document' object")
	}

	// try to get the 'body' property of JS from the global scope
	body := doc.Get("body")
	if !body.Truthy() {
		panic("unable to get 'body' object")
	}
	// create an 'H1' tag...
	h1 := doc.Call("createElement", "h1")
	// set some inner HTML
	h1.Set("innerHTML", "Hello from WebAssembly! (made with 💖 Go)")

	// attach the H1 element to the 'body'
	body.Call("appendChild", h1)
}
