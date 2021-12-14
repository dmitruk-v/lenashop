package main

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/evanw/esbuild/pkg/api"
)

func esbuild() api.BuildResult {

	result := api.Build(api.BuildOptions{
		LogLevel: api.LogLevelDebug,
		Define: map[string]string{
			"process.env.NODE_ENV": "\"development\"",
		},
		// EntryPoints: []string{
		// 	"./src/home/home.jsx",
		// 	"./src/products/products.jsx",
		// },
		EntryPointsAdvanced: []api.EntryPoint{
			{InputPath: "./src/home/home.jsx", OutputPath: "home"},
			{InputPath: "./src/products/products.jsx", OutputPath: "products"},
		},
		Outdir:   "../static/js/",
		Bundle:   true,
		External: []string{"react", "react-dom"},
		// MinifyIdentifiers: true,
		// MinifyWhitespace:  true,
		// MinifySyntax:      true,
		Write:       true,
		Incremental: true,
		Plugins:     []api.Plugin{},
	})
	return result
}

func watch(root string, result api.BuildResult, pollDelay time.Duration) {
	log.Println("--- Watching for changes every", pollDelay, "---")
	var (
		// Maps modification timestamps to path names.
		modMap = map[string]time.Time{}

		// Channel that simply notifies for any recursive change.
		ch = make(chan struct{})
	)

	// Starts a goroutine that polls every 100ms.
	go func() {
		for range time.Tick(pollDelay) {
			// Walk directory `src`. This means we are polling recursively.
			if err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}

				// Get the current path’s modification time; if no such modification time
				// exists, simply create a first write.
				if prev, ok := modMap[path]; !ok {
					modMap[path] = info.ModTime()
				} else {
					// Path has been modified; therefore get the new modification time and
					// update the map. Finally, emit an event on our channel.
					if next := info.ModTime(); prev != next {
						modMap[path] = next
						ch <- struct{}{}
					}
				}
				return nil
			}); err != nil {
				panic(err)
			}
		}
	}()

	for range ch {
		// fmt.Println("something changed")
		res := result.Rebuild()
		if len(res.Errors) > 0 {
			log.Println(res.Errors)
		}
	}
}

func main() {
	watch("./src/", esbuild(), 5*time.Second)
}
