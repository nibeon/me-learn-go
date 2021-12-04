package main

import (

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	a := app.New()
	// заголовок окна программы
	w := a.NewWindow("Картинки!")
	w.Resize(fyne.NewSize(1700, 500))

	img1 := canvas.NewImageFromFile("1.jpg")
	img2 := canvas.NewImageFromFile("2.jpg")

	w.SetContent(container.NewGridWithColumns(2, img1, img2))

	w.ShowAndRun()

}
