package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

func main() {
	a := app.New()
	// заголовок окна программы
	w := a.NewWindow("Рисование фигур")
	w.Resize(fyne.NewSize(500, 500))

	// круги-овалы
	circle1 := canvas.NewCircle(color.NRGBA{255, 0, 0, 255})
	circle1.StrokeColor = color.NRGBA{0, 255, 0, 255}
	circle1.StrokeWidth = 3

	// квадраты/прямоугольники
	rectangle1 := canvas.NewRectangle(color.NRGBA{255, 255, 0, 255})
	rectangle1.StrokeColor = color.NRGBA{0, 0, 0, 255}
	rectangle1.StrokeWidth = 3

	// линии
	line1 := canvas.NewLine(color.NRGBA{195, 195, 195, 255})
	line1.StrokeWidth = 5

	// иконки
	icon1 := widget.NewIcon(theme.FileIcon())ga

	btn := widget.NewButton("Изменить!", func() {
		rectangle1.Hide()

		circle1.FillColor = color.NRGBA{45, 36, 36, 255}
		circle1.StrokeWidth = 8

		line1.StrokeColor = color.NRGBA{0, 0, 0, 255}
	})

	content1 := container.NewGridWithColumns(2, circle1, rectangle1)
	content2 := container.NewGridWithColumns(2, line1, icon1)

	content := container.NewGridWithRows(3, content1, content2, btn)

	w.SetContent(content)

	w.ShowAndRun()

}
