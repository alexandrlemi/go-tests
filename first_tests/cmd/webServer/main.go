package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Калькулятор")

	input := widget.NewEntry()
	input.SetPlaceHolder("0")

	var result float64
	var operator string
	var isNewInput bool

	buttons := []string{
		"7", "8", "9", "/",
		"4", "5", "6", "*",
		"1", "2", "3", "-",
		"0", ".", "=", "+",
	}

	grid := container.NewGridWithColumns(4)

	for _, btn := range buttons {
		btnText := btn
		button := widget.NewButton(btnText, func() {
			if btnText >= "0" && btnText <= "9" || btnText == "." {
				if isNewInput {
					input.SetText("")
					isNewInput = false
				}
				input.SetText(input.Text + btnText)
			} else if btnText == "=" {
				secondOperand, _ := strconv.ParseFloat(input.Text, 64)
				switch operator {
				case "+":
					result += secondOperand
				case "-":
					result -= secondOperand
				case "*":
					result *= secondOperand
				case "/":
					if secondOperand != 0 {
						result /= secondOperand
					} else {
						input.SetText("Ошибка")
						return
					}
				}
				input.SetText(strconv.FormatFloat(result, 'f', -1, 64))
				isNewInput = true
			} else {
				result, _ = strconv.ParseFloat(input.Text, 64)
				operator = btnText
				isNewInput = true
			}
		})
		grid.Add(button)
	}

	myWindow.SetContent(container.NewVBox(
		input,
		grid,
	))

	myWindow.ShowAndRun()
}
