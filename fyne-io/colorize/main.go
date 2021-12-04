package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	//"fyne.io/fyne/v2/container"
	//"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	a := app.New()
	// заголовок окна программы
	w := a.NewWindow("Цветастость приложения!")
	w.Resize(fyne.NewSize(300, 300))

	color_for := color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	label := canvas.NewText("ЦВЕТ", color_for)

	w.SetContent(label)

	w.ShowAndRun()

}
