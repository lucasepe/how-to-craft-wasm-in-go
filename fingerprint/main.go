package main

import (
	"hash/crc32"
	"strconv"
	"syscall/js"
)

func main() {
	// Transform our Go function to one to be used by JavaScript
	fingerprintFunc := js.FuncOf(computeFingerprint)
	// Don't forget to relase it
	defer fingerprintFunc.Release()

	// Register and export our Go function to JavaScript
	js.Global().Set("fingerprint", fingerprintFunc)

	// Prevent the function from returning,
	// which is required in a wasm module
	select {}
}

func computeFingerprint(this js.Value, args []js.Value) interface{} {
	// try to get the 'document' property of JS from the global scope
	doc := js.Global().Get("document")
	if !doc.Truthy() { // the JS way of testing for nil
		panic("unable to get 'document' object")
	}

	// to measure performance
	perf := js.Global().Get("performance")
	if !perf.Truthy() { // the JS way of testing for nil
		panic("performance should be available")
	}
	start := perf.Call("now")

	// Create a Canvas
	canvas := doc.Call("createElement", "canvas")
	// and don't diplay it
	canvas.Set("display", "none")

	// Grab the 2D context
	ctx := canvas.Call("getContext", "2d")
	// draw something
	ctx.Call("beginPath")
	ctx.Call("fillText", "❤️🤪🎉👋", 50, 70)
	ctx.Call("stroke")

	// Convert the canvas to Base64 image data
	data := canvas.Call("toDataURL", "image/png")

	// Compute the Crc32 checksum
	crc := crc32.ChecksumIEEE([]byte(data.String()))

	// Keep track of the computation end time
	end := perf.Call("now")

	// return a map with the elapsed time and the fingerprint
	return map[string]interface{}{
		"ms": end.Int() - start.Int(),
		"id": strconv.FormatInt(int64(crc), 16),
	}
}
