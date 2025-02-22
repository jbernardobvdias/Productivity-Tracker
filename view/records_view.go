package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func DrawRecordsView(a fyne.App) {
	w := a.NewWindow("Records")

	getClasses := widget.NewButton("Classes", func() {
		DrawAcitivityView(a)
	})

	w.SetContent(
		getClasses,
	)
	w.Show()
}
