package main

import (
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	// Create views
	greeting := 
		tview.NewTextView().
		SetText("Hello Mateo\nWelcome to your dashboard!").
		SetWordWrap(true).
		SetDynamicColors(true).
		SetBorder(true).
		SetTitle("Greeting, quote + ASCII")
		
	countdowns := tview.NewTextView().SetBorder(true).SetTitle("Countdowns")
	weather := tview.NewTextView().SetBorder(true).SetTitle("Weather")
	fitness := tview.NewTextView().SetBorder(true).SetTitle("Fitness")
	stocks := tview.NewTextView().SetBorder(true).SetTitle("Stocks")
	fun_fact := tview.NewTextView().SetBorder(true).SetTitle("Fun fact")

	flex := 
		tview.NewFlex().
		AddItem(tview.NewFlex().
			AddItem(greeting, 0, 2, false).
			AddItem(countdowns, 0, 1, false), 0, 1, false).
		AddItem(tview.NewFlex().
			AddItem(weather, 0, 1, false).
			AddItem(fitness, 0, 3, false).
			AddItem(stocks, 0, 1, false), 0, 2, false).
		AddItem(fun_fact, 0, 1, false).SetDirection(tview.FlexRow)
		
	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}
}