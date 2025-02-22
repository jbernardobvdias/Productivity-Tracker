package main

import (
	"prod_tracker/data"
	"prod_tracker/view"
)

func main() {
	data.CreateTable()
	data.LoadTable()
	data.AddActivity()
	data.AddRecord()

	view.DrawHomeView()
}
