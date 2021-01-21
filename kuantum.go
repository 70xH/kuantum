package kuantum

import (
	"os"

	"github.com/rivo/tview"
)

func Kuantum() {

	// new app
	app := tview.NewApplication()

	// choice
	choice := tview.NewForm().
		AddButton("New VM", func () {
			app.Stop()
			CreateTUI()
			WriteTUI()
		}).
		AddButton("Existing VMs", nil).
		AddButton("Write to an existing vm", nil).
		AddButton("Run VM", nil).
		AddButton("Running VMs", nil).
		AddButton("Help", nil).
		AddButton("Quit", func() {
			app.Stop()
			os.Exit(0)
		})

	choice.SetBorder(true).SetTitle("Menu").SetTitleAlign(tview.AlignLeft)
	if err := app.SetRoot(choice, true).SetFocus(choice).Run(); err != nil {
		panic(err)
	}
}
