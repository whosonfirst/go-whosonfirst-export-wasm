wasm:
	GOOS=js GOARCH=wasm go build -mod vendor -o static/wasm/export_feature.wasm cmd/export-feature/main.go

example:
	go run -mod vendor cmd/example/main.go
