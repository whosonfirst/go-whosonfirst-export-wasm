package static

import (
	"embed"
)

//go:embed javascript/* wasm/*
var FS embed.FS
