package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello Bro")

	w.SetContext(widget.NewLabel("Hello Bro!"))
	w.ShowAndRun()
}
