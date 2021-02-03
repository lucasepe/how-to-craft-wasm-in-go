package main

import (
	"hash/crc32"
	"strconv"
	"syscall/js"
)

func main() {
	fingerprintFunc := fingerPrint()
	defer fingerprintFunc.Release()
	exportFunction("fingerprint", fingerprintFunc)

	// Prevent the function from returning, which is required in a wasm module
	select {}
}

func fingerPrint() js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		// try to get the 'document' property of JS from the global scope
		doc := js.Global().Get("document")
		if !doc.Truthy() { // the JS way of testing for nil
			panic("unable to get 'document' object")
		}

		perf := js.Global().Get("performance")
		if !perf.Truthy() { // the JS way of testing for nil
			panic("performance should be available")
		}
		start := perf.Call("now")

		// Make the Canvas
		canvas := doc.Call("createElement", "canvas")

		canvas.Set("display", "none")
		ctx := canvas.Call("getContext", "2d")
		ctx.Call("beginPath")
		//ctx.Call("fillStyle")
		ctx.Call("fillText", "❤️🤪🎉👋", 50, 70)
		ctx.Call("stroke")

		data := canvas.Call("toDataURL", "image/png")
		crc := crc32.ChecksumIEEE([]byte(data.String()))
		end := perf.Call("now")

		return map[string]interface{}{
			"ms": end.Int() - start.Int(),
			"id": strconv.FormatInt(int64(crc), 16),
		}
	}

	return js.FuncOf(fn)
}

// exportFunction is an helper function to register and export Go function to JS
func exportFunction(functionName string, function js.Func) {
	js.Global().Set(functionName, function)
}
