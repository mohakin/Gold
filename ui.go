package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func (app *Config) makeUI() {
	// NOTE: get the current price of gold - Finished
	openPrice, currentPrice, priceChange := app.getPriceText()

	// put the price information into a container

	priceContent := container.NewGridWithColumns(3,
		openPrice,
		currentPrice,
		priceChange,
	)

	app.PriceContainer = priceContent

	// NOTE: get tool bar - Finished
	toolBar := app.getToolBar()
	app.ToolBar = toolBar

	priceTabContent := app.pricesTab()
	holdingsTabContent := app.holdingsTab()

	// NOTE: get app tabs - Finished
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon(
			"Prices:",
			theme.HomeIcon(),
			priceTabContent,
		),
		container.NewTabItemWithIcon(
			"Holdings:",
			theme.InfoIcon(),
			holdingsTabContent,
		),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	// NOTE: add a container window - Finished
	finalContent := container.NewVBox(priceContent, toolBar, tabs)

	app.MainWindow.SetContent(finalContent)

	go func() {
		for range time.Tick(time.Second * 15) {
			app.refreshPriceContent()
		}
	}()
}

func (app *Config) refreshPriceContent() {
	app.InfoLog.Print("refreshing prices")
	open, current, change := app.getPriceText()
	app.PriceContainer.Objects = []fyne.CanvasObject{open, current, change}
	app.PriceContainer.Refresh()

	chart := app.getChart()
	app.PriceChartContainer.Objects = []fyne.CanvasObject{chart}
	app.PriceChartContainer.Refresh()
}

func (app *Config) refreshHoldingsTable() {
	app.Holdings = app.getHoldingSlice()
	app.HoldingsTable.Refresh()
}
