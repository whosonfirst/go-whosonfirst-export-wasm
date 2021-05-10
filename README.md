# go-whosonfirst-export-wasm

Go package for compiling the `Export` method of the go-whosonfirst-export package to a JavaScript-compatible WebAssembly (wasm) binary. It also provides a net/http middleware packages for appending the necessary static assets and HTML resources to use the wasm binary in web applications.

## Example

```
package main

import (
	"embed"
	"flag"
	"fmt"
	wasm "github.com/whosonfirst/go-whosonfirst-export-wasm/http"
	"log"
	"net/http"
)

//go:embed index.html example.*
var FS embed.FS

func main() {

	host := flag.String("host", "localhost", "The host name to listen for requests on")
	port := flag.Int("port", 8080, "The host port to listen for requests on")

	flag.Parse()

	mux := http.NewServeMux()

	wasm.AppendAssetHandlers(mux)

	http_fs := http.FS(FS)
	example_handler := http.FileServer(http_fs)

	wasm_opts := wasm.DefaultWASMOptions()
	wasm_opts.EnableWASMExec()

	example_handler = wasm.AppendResourcesHandler(example_handler, wasm_opts)

	mux.Handle("/", example_handler)

	addr := fmt.Sprintf("%s:%d", *host, *port)
	log.Printf("Listening for requests on %s\n", addr)

	http.ListenAndServe(addr, mux)
}
```

_Error handling omitted for brevity._

## See also

* https://github.com/whosonfirst/go-whosonfirst-export