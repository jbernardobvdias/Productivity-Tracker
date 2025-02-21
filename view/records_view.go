package view

import (
	"fyne.io/fyne/app"
)

func DrawRecordsView() {
	a := app.New()
	w := a.NewWindow("Productivity Tracker")

	w.ShowAndRun()
}
