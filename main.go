package main

import (
	"prod_tracker/data"
	"prod_tracker/view"
)

func main() {
	data.CreateTable()
	data.AddActivity("Test1")
	data.AddActivity("Test2")

	view.DrawHomeView()
}
