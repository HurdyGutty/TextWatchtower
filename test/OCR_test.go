package OCR_test

import (
	"os"
	"regexp"
	"testing"

	"github.com/HurdyGutty/go_OCR/pkg/OCR"
)

func TestOCROutput(t *testing.T) {
	img := "images/captureZone1.png"
	trainingData := os.DirFS("Tesseract/tessdata/eng.*")
	text := OCR.ProcessOCR(trainingData, img)

	want := regexp.MustCompile("Tesseract was originally developed at Hewlett-Packard Laboratories Bristol UK and at Hewlett-Packard Co, Greeley\r\nColorado USA between 1985 and 1994, with some more changes made in 1996 to port to Windows, and some C+\r\n+izing in 1998. In 2005 Tesseract was open sourced by HP. From 2006 until November 2018 it was developed by\r\nGoogle.")

	if !want.MatchString(text) {
		t.Fatalf(`Hello("Gladys") = %q, want match for %#q`, text, want)
	}
}
