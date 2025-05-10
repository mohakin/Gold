package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"testing"

	"goldwatcher/repository"

	"fyne.io/fyne/v2/test"
)

var testApp Config

func TestMain(m *testing.M) {
	a := test.NewApp()
	testApp.App = a
	testApp.MainWindow = a.NewWindow("")
	testApp.HTTPClient = client
	testApp.DB = repository.NewTestRepository()
	os.Exit(m.Run())
}

var jsonToReturn = `

{
  "ts": 1739654854858,
  "tsj": 1739654845681,
  "date": "Feb 15th 2025, 04:27:25 pm NY",
  "items": [
    {
      "curr": "USD",
      "xauPrice": 2882.735,
      "xagPrice": 32.146,
      "chgXau": -33.905,
      "chgXag": -0.0504,
      "pcXau": -1.1625,
      "pcXag": -0.1565,
      "xauClose": 2916.64,
      "xagClose": 32.19645
    }
  ]
}
	`

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

var client = NewTestClient(func(req *http.Request) *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(jsonToReturn)),
		Header:     make(http.Header),
	}
})
