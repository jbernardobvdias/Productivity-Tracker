package view

import (
	"prod_tracker/data"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func DrawAcitivityView(a fyne.App) {
	w := a.NewWindow("Activities")
	w.Resize(fyne.NewSize(250, 250))

	// Gets the data and sets the table with all the different activities
	selectedName := ""
	activities := data.GetActivitiesString()
	list := widget.NewTable(
		func() (int, int) {
			return len(activities), 2
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(activities[i.Row][i.Col])
		})

	list.OnSelected = func(id widget.TableCellID) {
		selectedName = activities[id.Row][0]
	}
	// Place to input the activity name
	inputName := widget.NewEntry()

	saveButton := widget.NewButton("+", func() {
		data.AddActivity(inputName.Text)
		activities = data.GetActivitiesString()
		list.Refresh()
	})
	deleteButton := widget.NewButton("-", func() {
		data.DeleteActivity(selectedName)
		activities = data.GetActivitiesString()
		list.Refresh()
	})

	w.SetContent(
		container.NewGridWithRows(4,
			list,
			inputName,
			saveButton,
			deleteButton,
		),
	)
	w.Show()
}
