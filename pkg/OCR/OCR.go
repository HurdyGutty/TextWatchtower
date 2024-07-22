package OCR

import (
	"context"
	"io/fs"
	"log"
	"os"

	"github.com/HurdyGutty/go_OCR/pkg/helper/logger"
	"github.com/danlock/gogosseract"
)

func handleErr(err error) {
	if err != nil {
		logger.Error(err)
	}
}

func ProcessOCR(trainingData fs.FS, image string) string {
	ctx := context.Background()
	trainingDataFile, err := trainingData.Open("Tesseract/tessdata/eng.traineddata")
	handleErr(err)
	cfg := gogosseract.Config{
		Language:     "eng",
		TrainingData: trainingDataFile,
	}

	f, err := os.OpenFile("tesseractLogfile.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	cfg.Stderr = f
	cfg.Stdout = f

	tess, err := gogosseract.New(ctx, cfg)
	handleErr(err)

	imageFile, err := os.Open(image)
	handleErr(err)

	err = tess.LoadImage(ctx, imageFile, gogosseract.LoadImageOptions{})
	handleErr(err)

	text, err := tess.GetText(ctx, func(progress int32) { log.Printf("Tesseract parsing is %d%% complete.", progress) })
	handleErr(err)
	handleErr(tess.Close(ctx))
	return text
}
