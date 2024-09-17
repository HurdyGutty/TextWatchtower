package main

import (
	"context"
	"embed"
	"log"
	"os"

	"github.com/HurdyGutty/go_OCR/pkg/OCR"
	sse "github.com/HurdyGutty/go_OCR/pkg/SSE"
	"github.com/HurdyGutty/go_OCR/pkg/captureGroup"
	"github.com/HurdyGutty/go_OCR/pkg/instruct"
	"github.com/HurdyGutty/go_OCR/pkg/telegram"
	"github.com/joho/godotenv"
	"github.com/lxn/win"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed Tesseract/tessdata/*
var tesseract embed.FS

func main() {
	captureGroup.TrainingData = tesseract
	errorLoad := godotenv.Load()
	if errorLoad != nil {
		log.Fatal("Error loading .env file")
	}
	telegram.Token = os.Getenv("TELEGRAM_BOT_TOKEN")
	telegram.Chat_ID = os.Getenv("GROUP_ID")

	hDC := win.GetDC(0)
	defer win.ReleaseDC(0, hDC)
	width := int(win.GetDeviceCaps(hDC, win.HORZRES))
	height := int(win.GetDeviceCaps(hDC, win.VERTRES))

	// Create an instance of the app structure
	app := NewApp()
	app.Width = width
	app.Height = height

	instructBoard := instruct.NewInstructBoard()
	app.instructBoard = instructBoard
	captureGroup.InstructionBoard = instructBoard
	OCR.InstructionBoard = instructBoard
	sse.InstructBoard = instructBoard

	go func() {
		sse.Serve()
	}()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Text Watchtower",
		Width:  width,
		Height: height,
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
		WindowStartState: options.Fullscreen,
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
