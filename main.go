package main

import (
	_ "embed"
	"fmt"

	"github.com/wailsapp/wails"
)

func go_string(applicantName string, tenancyLength int, feesAmount float64) string {
	return fmt.Sprintf("The Applicant Id is: %s, the Fees Amount is: %f, and the Tenancy Lenght is: %d", applicantName, feesAmount, tenancyLength)
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
