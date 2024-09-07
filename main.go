package main

import (
	"context"
	"embed"
	"io/fs"
	"log"
	"os"

	"github.com/HurdyGutty/go_OCR/pkg/captureGroup"
	"github.com/HurdyGutty/go_OCR/pkg/instruct"
	"github.com/HurdyGutty/go_OCR/pkg/telegram"
	"github.com/joho/godotenv"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed Tesseract/tessdata/*
var tesseract embed.FS

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

func main() {
	captureGroup.TrainingData = tesseract
	errorLoad := godotenv.Load()
	if errorLoad != nil {
		log.Fatal("Error loading .env file")
	}
	telegram.Token = os.Getenv("TELEGRAM_BOT_TOKEN")
	telegram.Chat_ID = os.Getenv("CHAT_ID")

	// Create an instance of the app structure
	app := NewApp()
	instructBoard := instruct.NewInstructBoard()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Text Watchtower",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Frameless:        true,
		BackgroundColour: &options.RGBA{R: 255, G: 0, B: 0, A: 0},
		Windows: &windows.Options{
			WebviewIsTransparent:              true,
			WindowIsTranslucent:               true,
			DisableFramelessWindowDecorations: true,
		},
		WindowStartState: options.Maximised,
		OnStartup: func(ctx context.Context) {
			app.startup(ctx)
			instructBoard.Startup(ctx)
		},
		Bind: []interface{}{
			app,
			instructBoard,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
