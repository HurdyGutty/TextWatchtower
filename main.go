package main

import (
	"embed"

	"github.com/HurdyGutty/go_OCR/pkg/instruct"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
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
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			instructBoard,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
