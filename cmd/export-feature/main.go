package main

import (
	"context"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-export/v2"
	"log"
	"syscall/js"
)

func ExportFunc(ex export.Exporter) js.Func {

	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		geojson_data := args[0].String()

		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			resolve := args[0]
			reject := args[1]

			go func() {

				ctx := context.Background()
				feature_data, err := ex.Export(ctx, []byte(geojson_data))

				if err != nil {
					reject.Invoke(fmt.Sprintf("Failed to export data, %v", err))
					return
				}

				resolve.Invoke(string(feature_data))
			}()

			return nil
		})

		promiseConstructor := js.Global().Get("Promise")
		return promiseConstructor.New(handler)
	})
}

func main() {

	ctx := context.Background()
	ex, err := export.NewExporter(ctx, "whosonfirst://")

	if err != nil {
		log.Fatalf("Failed to create new exporter, %v", err)
	}

	export_func := ExportFunc(ex)
	defer export_func.Release()

	js.Global().Set("export_feature", export_func)

	c := make(chan struct{}, 0)

	log.Println("WOF export feature WASM binary initialized")
	<-c
}
