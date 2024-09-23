package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/davidbyttow/govips/v2/vips"
)

var filenames = []string{"input/IMG_2608.jpg"}

const output = "output"

func main() {
	wg := new(sync.WaitGroup)
	for i, name := range filenames {
		wg.Add(1)
		go func(index int, filename string) {
			defer wg.Done()
			ref, err := vips.NewImageFromFile(filename)
			if err != nil {
				log.Fatal(err)
			}
			defer ref.Close()
			for q := 10; q <= 100; q += 10 {
				params := vips.NewJpegExportParams()
				params.Quality = q
				out, _, _ := ref.ExportJpeg(params)
				path := filepath.Join(output, fmt.Sprintf("%d_%d.jpg", index, q))
				os.WriteFile(path, out, 0644)
			}
		}(i, name)
	}
	wg.Wait()
}
