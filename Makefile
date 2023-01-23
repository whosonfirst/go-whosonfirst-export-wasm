GOROOT=$(shell go env GOROOT)

rebuild:
	@make wasmjs
	@make wasm

wasm:
	GOOS=js GOARCH=wasm go build -mod vendor -ldflags="-s -w" -o static/wasm/export_feature.wasm cmd/export-feature/main.go

example:
	go run -mod vendor cmd/example/main.go

wasmjs:
	cp "$(GOROOT)/misc/wasm/wasm_exec.js" static/javascript/
