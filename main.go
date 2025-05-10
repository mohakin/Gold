package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"goldwatcher/repository"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"

	_ "github.com/glebarez/go-sqlite"
)

// NOTE: Config is the type used to share data with various parts of our application.
// NOTE: It includes the parts of our GUI that are dynamic and will need to be updated,
// NOTE: such as the holdings table, gold price info, and the chart. In order to refresh
// NOTE: those things, we need a reference to them, and this is a convenient place to put
// NOTE: them, insead of package level variables.
type Config struct {
	App                            fyne.App
	InfoLog                        *log.Logger
	ErrorLog                       *log.Logger
	DB                             repository.Repository
	MainWindow                     fyne.Window
	PriceContainer                 *fyne.Container
	ToolBar                        *widget.Toolbar
	PriceChartContainer            *fyne.Container
	Holdings                       [][]interface{}
	HoldingsTable                  *widget.Table
	HTTPClient                     *http.Client
	AddHoldingsPurchaseAmountEntry *widget.Entry
	AddHoldingsPurchaseDateEntry   *widget.Entry
	AddHoldingsPurchasePriceEntry  *widget.Entry
}

func main() {
	var myApp Config

	// TODO: create a fyne app - Finished
	fyneApp := app.New()
	myApp.App = fyneApp
	myApp.HTTPClient = &http.Client{}

	// TODO: create our loggers - Finished
	myApp.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	myApp.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// TODO: open a connection to the database
	sqlDB, err := myApp.connectSQL()
	if err != nil {
		log.Panic(err)
	}

	// TODO: create a database repository
	myApp.setupDB(sqlDB)

	currency = fyneApp.Preferences().StringWithFallback("currency", "CAD")

	// TODO: create and size a fyne window - Finished
	myApp.MainWindow = fyneApp.NewWindow("Gold Watcher")
	myApp.MainWindow.Resize(fyne.NewSize(1024, 768))
	myApp.MainWindow.SetFixedSize(false)
	myApp.MainWindow.SetMaster()

	myApp.makeUI()

	// TODO: show and run the app - Finished

	myApp.MainWindow.ShowAndRun()
}

func (app *Config) connectSQL() (*sql.DB, error) {
	path := ""

	if os.Getenv("DB_PATH") != "" {
		path = os.Getenv("DB_PATH")
	} else {
		path = app.App.Storage().RootURI().Path() + "/sql.db"
		app.InfoLog.Println()
	}

	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (app *Config) setupDB(sqlDB *sql.DB) {
	app.DB = repository.NewSQLiteRepository(sqlDB)

	err := app.DB.Migrate()
	if err != nil {
		app.ErrorLog.Println(err)
		log.Panic()
	}
}
