package main

import (
	"prod_tracker/data"
	"prod_tracker/view"
)

func main() {
	data.CreateTable()
	view.DrawHomeView()
}
