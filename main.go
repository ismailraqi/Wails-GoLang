package main

import (
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func basic() string {
	return "World!"
}
func basic1() string {
	return "basic1!"
}
func main() {

	js := mewn.String("./frontend/dist/my-app/main.js")
	css := mewn.String("./frontend/dist/my-app/styles.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "test",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(basic1)
	app.Run()
}
