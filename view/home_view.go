package view

import (
	"fmt"
	"prod_tracker/data"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func DrawHomeView() {
	a := app.New()
	w := a.NewWindow("Productivity Tracker")
	w.Resize(fyne.NewSize(250, 250))
	t := 0
	isRunning := false
	selectedClass := ""

	activities := data.GetActivitiesString()

	selected := widget.NewSelect(activities, func(s string) {
		selectedClass = s
		fmt.Println(selectedClass)
	})
	counter := widget.NewLabel("00:00:00")

	// Set the buttons that interact with time.
	startCount := widget.NewButton("Start", func() {
		isRunning = true
		go StartTimer(&t, counter, &isRunning)
	})
	stopCount := widget.NewButton("Stop", func() {
		isRunning = false
	})
	resetCount := widget.NewButton("Reset", func() {
		isRunning = false
		t = 0
		counter.SetText("00:00:00")
	})
	saveButton := widget.NewButton("Save", func() {
		data.AddRecord(selectedClass, time.Now().GoString(), t)
		isRunning = false
		t = 0
		counter.SetText("00:00:00")
	})

	// Set the buttons that open windows.
	getRecords := widget.NewButton("Records", func() {
		DrawRecordsView(a)
	})
	getClasses := widget.NewButton("Classes", func() {
		DrawAcitivityView(a)
	})

	w.SetContent(
		container.NewVBox(
			selected,
			counter,
			startCount,
			stopCount,
			resetCount,
			saveButton,
			getRecords,
			getClasses,
		))

	w.ShowAndRun()
}

func StartTimer(t *int, counter *widget.Label, isRunning *bool) {
	for {
		time.Sleep(1 * time.Second)
		if !*isRunning {
			return
		}
		*t += 1
		counter.SetText(TranslateSeconds(*t))
	}
}

func TranslateSeconds(secs int) string {
	hours := (secs - secs%3600) / 3600
	mins := (secs - secs%60) / 60
	secs = secs % 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, mins, secs)
}
