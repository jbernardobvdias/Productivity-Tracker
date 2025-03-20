package view

import (
	"prod_tracker/data"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func DrawRecordsView(a fyne.App) {
	w := a.NewWindow("Records")
	w.Resize(fyne.NewSize(250, 250))

	// Gets the data and sets the table with all the different records
	selectedId := ""
	records := data.GetRecordsString()
	list := widget.NewTable(
		func() (int, int) {
			return len(records), 4
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(records[i.Row][i.Col])
		})
	list.OnSelected = func(id widget.TableCellID) {
		print()
		selectedId = records[id.Row][0]
	}

	deleteButton := widget.NewButton("-", func() {
		data.DeleteActivity(selectedId)
		records = data.GetRecordsString()
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
