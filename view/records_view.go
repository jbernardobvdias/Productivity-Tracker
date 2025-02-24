package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var dataMock2 = [][]string{{"top left", "top right"},
	{"bottom left", "bottom right"}}

func DrawRecordsView(a fyne.App) {
	w := a.NewWindow("Records")
	w.Resize(fyne.NewSize(250, 250))

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

	w.SetContent(
		list,
	)
	w.Show()
}
