OUT_NAME := $(shell basename "$(PWD)")
OUT_DIR  := assets

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

.DEFAULT_GOAL := help

## wasm: Builds the WASM
wasm: clean
	GOOS=js GOARCH=wasm go build -o $(OUT_DIR)/$(OUT_NAME).wasm
	cp $(shell go env GOROOT)/misc/wasm/wasm_exec.js $(OUT_DIR)/

## serve: Starts the HTTP server
serve:
	ran -port=8888 -g -index=index.html

## dump: Dumps information about the WASM binary
dump: wasm
	wasm-objdump -h -j export -x $(OUT_DIR)/$(OUT_NAME).wasm

## clean: Clean the dist folder
clean:
	rm -f $(OUT_DIR)/*.wasm 

