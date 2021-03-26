// test project main.go
package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	win := a.NewWindow("Hello World")
	win.SetContent(container.NewVBox(
		widget.NewLabel("Hello World!"),
		widget.NewButton("Quit", func() {
			a.Quit()
		}),
	))
	win.ShowAndRun()
}
