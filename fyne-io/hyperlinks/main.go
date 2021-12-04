package main

import (
	//"fmt"
	//"strconv"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	// заголовок окна программы
	w := a.NewWindow("Ссылка на веб-страницу")
	w.Resize(fyne.NewSize(300, 400))

	url, _ := url.Parse("https://fyne.io")
	link := widget.NewHyperlink("Оффициальный сайт Fyne", url)

	w.SetContent(container.NewVBox(
		link,
	))

	w.ShowAndRun()

}
