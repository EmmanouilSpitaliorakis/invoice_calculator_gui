package main

import (
	_ "embed"

	"github.com/wailsapp/wails"
)

// func basic() string {
// 	return "World!"
// }

func go_string() string {
	return "Bacon!"
}

//go:embed frontend/build/static/js/main.js
var js string

//go:embed frontend/build/static/css/main.css
var css string

func main() {

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 600,
		Title:  "Invoice Calculator",
		JS:     js,
		CSS:    css,
		Colour: "#ffffff",
	})
	app.Bind(go_string)
	app.Run()
}
