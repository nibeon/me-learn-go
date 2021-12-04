package main

import (
	//"fmt"
	//"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	// заголовок окна программы
	w := a.NewWindow("Заказ еды")
	w.Resize(fyne.NewSize(300, 500))

	title := widget.NewLabel("Оформление заказа")

	name_label := widget.NewLabel("Введите ваше имя:")
	name := widget.NewEntry()

	food_label := widget.NewLabel("Выерите еду для заказа:")
	food := widget.NewCheckGroup([]string{"Пицца","Торт","Пиво","Чипсы","Сосиски"}, func(s []string) {})

	sex_label := widget.NewLabel("Ваш пол:")
	sex := widget.NewRadioGroup([]string{"Мужской","Женский"}, func(s string) {})

	result := widget.NewLabel("")

	btn := widget.NewButton("Оформить заказ", func() {
		username := name.Text
		food_arr := food.Selected
		sex_value := sex.Selected

		result.SetText(username + "\n" + sex_value + "\n")

		for _, i := range food_arr {
			result.SetText(result.Text + i + "\n")
		}
	})

	w.SetContent(container.NewVBox(
		title,
		name_label,
		name,
		food_label,
		food,
		sex_label,
		sex,
		btn,
		result,
	))

	w.ShowAndRun()

}
