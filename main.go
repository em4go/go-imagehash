package main

import (
	"fmt"
	"image"
	"log"
	"os"

	_ "golang.org/x/image/webp" // ← registra WebP
)

func main() {
	path := "./testdata/greninja-event.webp"

	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("abriendo %s: %v", path, err)
	}
	defer f.Close()

	img, format, err := image.Decode(f)
	if err != nil {
		log.Fatalf("decodificando %s: %v", path, err)
	}
	b := img.Bounds()
	fmt.Printf("Formato detectado: %s — %dx%d px\n", format, b.Dx(), b.Dy())
}
