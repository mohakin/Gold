package main

import (
	"strconv"
	"time"

	"goldwatcher/repository"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (app *Config) getToolBar() *widget.Toolbar {
	toolBar := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			app.addHoldingsDialog()
		}),
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {
			app.refreshPriceContent()
		}),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {
			w := app.showPreferences()
			w.Resize(fyne.NewSize(300, 200))
			w.Show()
		}),
	)

	return toolBar
}

func (app *Config) showPreferences() fyne.Window {
	win := app.App.NewWindow("preferences")

	lbl := widget.NewLabel("Preferred Currency")
	cur := widget.NewSelect([]string{"USD", "CAD", "GBP"}, func(value string) {
		currency = value
		app.App.Preferences().SetString("currency", value)
	})
	cur.Selected = currency

	btn := widget.NewButton("Save", func() {
		win.Close()
		app.refreshPriceContent()
	})
	btn.Importance = widget.HighImportance

	win.SetContent(container.NewVBox(lbl, cur, btn))

	return win
}

func (app *Config) addHoldingsDialog() dialog.Dialog {
	addAmountEntry := widget.NewEntry()
	purchaseDateEntry := widget.NewEntry()
	purchasePriceEntry := widget.NewEntry()

	app.AddHoldingsPurchaseAmountEntry = addAmountEntry
	app.AddHoldingsPurchaseDateEntry = purchaseDateEntry
	app.AddHoldingsPurchasePriceEntry = purchasePriceEntry

	dateValidator := func(s string) error {
		if _, err := time.Parse("2006-01-02", s); err != nil {
			return err
		}
		return nil
	}
	purchaseDateEntry.Validator = dateValidator

	isIntValidator := func(s string) error {
		_, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		return nil
	}
	addAmountEntry.Validator = isIntValidator

	isFLoatValidator := func(s string) error {
		_, err := strconv.ParseFloat(s, 32)
		if err != nil {
			return err
		}
		return nil
	}
	purchasePriceEntry.Validator = isFLoatValidator

	purchaseDateEntry.PlaceHolder = "YYYY-MM-DD"

	// NOTE: create dialog
	addForm := dialog.NewForm(
		"Add Gold",
		"Add",
		"Cancel",
		[]*widget.FormItem{
			{Text: "Amount in toz", Widget: addAmountEntry},
			{Text: "Purchase Price", Widget: purchasePriceEntry},
			{Text: "Purchase Date", Widget: purchaseDateEntry},
		},
		func(valid bool) {
			if valid {
				amount, _ := strconv.Atoi(addAmountEntry.Text)
				purcahseDate, _ := time.Parse("2006-01-02", purchaseDateEntry.Text)
				purchasePrice, _ := strconv.ParseFloat(purchasePriceEntry.Text, 32)
				purchasePrice = purchasePrice * 100.0

				_, err := app.DB.InsertHolding(repository.Holdings{
					Amount:        amount,
					PurchaseDate:  purcahseDate,
					PurchasePrice: int(purchasePrice),
				})
				if err != nil {
					app.ErrorLog.Println(err)
				}
				app.refreshHoldingsTable()
			}
		},
		app.MainWindow)

	// NOTE: size and show the dialog
	addForm.Resize(fyne.Size{Width: 400})
	addForm.Show()

	return addForm
}
