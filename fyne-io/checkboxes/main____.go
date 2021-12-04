package main

import (
	"fmt"
	//"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	// заголовок окна программы
	w := a.NewWindow("Подпишись")
	w.Resize(fyne.NewSize(300, 500))

	checks := widget.NewCheckGroup([]string{"Check 1","Check 2","Check 3"}, func(s []string) {

	})

	radio := widget.NewRadioGroup([]string{"Radio 1", "Radio 2"}, func(s string) {

	})

	btn := widget.NewButton("Нажми", func() {
		for _, i := range checks.Selected {
			fmt.Println(i)
		}
		fmt.Println(radio.Selected)
	})

	w.SetContent(container.NewVBox(
		checks,
		radio,
		btn,
	))

	w.ShowAndRun()

}
