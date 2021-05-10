// This package provides tools for compiling the `Export` method of the go-whosonfirst-export package to a JavaScript-compatible WebAssembly (wasm) binary.
//
// It also provides a net/http middleware packages for appending the necessary static assets and HTML resources to use the wasm binary in web applications.
//
// To build the WebAssembly binary run the following command:
//
//	GOOS=js GOARCH=wasm go build -mod vendor -o export_feature.wasm cmd/export-feature/main.go
package wasm
