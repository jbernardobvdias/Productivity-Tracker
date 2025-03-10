package view

import (
	"prod_tracker/data"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var dataMock2 = [][]string{{"Activity", "Date", "Time"},
	{"Coding", "10/03/2024", "00:01:30"},
	{"Gym", "10/03/2024", "01:00:30"},
	{"Coding", "10/03/2024", "00:01:30"},
	{"Gym", "10/03/2024", "01:00:30"}}

func DrawRecordsView(a fyne.App) {
	w := a.NewWindow("Records")
	w.Resize(fyne.NewSize(250, 250))

	// Gets the data and sets the table with all the different records
	selectedName := ""
	activities := []string{"", ""}
	list := widget.NewTable(
		func() (int, int) {
			return len(dataMock2), len(dataMock2[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(dataMock2[i.Row][i.Col])
		})
	list.OnSelected = func(id widget.TableCellID) {
		selectedName = activities[id.Row]
	}

	deleteButton := widget.NewButton("-", func() {
		data.DeleteActivity(selectedName)
		activities = data.GetActivitiesString()
		list.Refresh()
	})

	w.SetContent(
		container.NewGridWithRows(2,
			list,
			deleteButton,
		),
	)
	w.Show()
}
