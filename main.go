package main

import (
	"embed"
	"fmt"

	"github.com/HurdyGutty/go_OCR/pkg/screenBox"
)

//go:embed Tesseract/tessdata/eng.*
var trainingData embed.FS

func main() {
	x1, y1, w, h := screenBox.DrawBox()

	// image := "images/captureZone1.png"

	// text := OCR.ProcessOCR(trainingData, image)

	// println(text)
	fmt.Println(x1, y1, w, h)
}
