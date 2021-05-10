package wasm

import (
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-export-wasm/static"
	"github.com/aaronland/go-http-rewrite"
	"io/fs"
	_ "log"
	"net/http"
	"path/filepath"
	"strings"
)

// WASMOptions provides a list of JavaScript and CSS link to include with HTML output.
type WASMOptions struct {
	JS  []string
	CSS []string
}

// Append the Javascript and CSS URLs for the wasm_exec.js script.
func (opts *WASMOptions) EnableWASMExec() {
	opts.JS = append(opts.JS, "/javascript/wasm_exec.js")
}

// Return a *WASMOptions struct with default paths and URIs.
func DefaultWASMOptions() *WASMOptions {

	opts := &WASMOptions{
		CSS: []string{

		},
		JS: []string{
		},
	}

	return opts
}

// AppendResourcesHandler will rewrite any HTML produced by previous handler to include the necessary markup to load WASM JavaScript and CSS files and related assets.
func AppendResourcesHandler(next http.Handler, opts *WASMOptions) http.Handler {
	return AppendResourcesHandlerWithPrefix(next, opts, "")
}

// AppendResourcesHandlerWithPrefix will rewrite any HTML produced by previous handler to include the necessary markup to load WASM JavaScript files and related assets ensuring that all URIs are prepended with a prefix.
func AppendResourcesHandlerWithPrefix(next http.Handler, opts *WASMOptions, prefix string) http.Handler {

	js := opts.JS
	css := opts.CSS

	if prefix != "" {

		for i, path := range js {
			js[i] = appendPrefix(prefix, path)
		}

		for i, path := range css {
			css[i] = appendPrefix(prefix, path)
		}
	}

	ext_opts := &rewrite.AppendResourcesOptions{
		JavaScript:  js,
		Stylesheets: css,
	}

	return rewrite.AppendResourcesHandler(next, ext_opts)
}

// AssetsHandler returns a net/http FS instance containing the embedded WASM assets that are included with this package.
func AssetsHandler() (http.Handler, error) {

	http_fs := http.FS(static.FS)
	return http.FileServer(http_fs), nil
}

// AssetsHandler returns a net/http FS instance containing the embedded WASM assets that are included with this package ensuring that all URLs are stripped of prefix.
func AssetsHandlerWithPrefix(prefix string) (http.Handler, error) {

	fs_handler, err := AssetsHandler()

	if err != nil {
		return nil, err
	}

	prefix = strings.TrimRight(prefix, "/")

	if prefix == "" {
		return fs_handler, nil
	}

	rewrite_func := func(req *http.Request) (*http.Request, error) {
		req.URL.Path = strings.Replace(req.URL.Path, prefix, "", 1)
		return req, nil
	}

	rewrite_handler := rewrite.RewriteRequestHandler(fs_handler, rewrite_func)
	return rewrite_handler, nil
}

// Append all the files in the net/http FS instance containing the embedded WASM assets to an *http.ServeMux instance.
func AppendAssetHandlers(mux *http.ServeMux) error {
	return AppendAssetHandlersWithPrefix(mux, "")
}

// Append all the files in the net/http FS instance containing the embedded WASM assets to an *http.ServeMux instance ensuring that all URLs are prepended with prefix.
func AppendAssetHandlersWithPrefix(mux *http.ServeMux, prefix string) error {

	asset_handler, err := AssetsHandlerWithPrefix(prefix)

	if err != nil {
		return nil
	}

	walk_func := func(path string, info fs.DirEntry, err error) error {

		if path == "." {
			return nil
		}

		if info.IsDir() {
			return nil
		}

		if prefix != "" {
			path = appendPrefix(prefix, path)
		}

		if !strings.HasPrefix(path, "/") {
			path = fmt.Sprintf("/%s", path)
		}

		// log.Println("APPEND", path)

		mux.Handle(path, asset_handler)
		return nil
	}

	return fs.WalkDir(static.FS, ".", walk_func)
}

func appendPrefix(prefix string, path string) string {

	prefix = strings.TrimRight(prefix, "/")

	if prefix != "" {
		path = strings.TrimLeft(path, "/")
		path = filepath.Join(prefix, path)
	}

	return path
}
