package OCR

import (
	"context"
	"embed"
	"io/fs"
	"log"
	"os"

	"github.com/danlock/gogosseract"
)

func getAllFilenames(efs *embed.FS) (files []string, err error) {
	if err := fs.WalkDir(efs, ".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		files = append(files, path)

		return nil
	}); err != nil {
		return nil, err
	}

	return files, nil
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func ProcessOCR(trainingData embed.FS, image string) string {
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
