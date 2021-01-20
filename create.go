package kuantum

import (
	"fmt"
	"os/exec"
	"os"
	"strconv"

	"github.com/rivo/tview"
)

func CreateTUI() {

	// init app
	app := tview.NewApplication()

	// image details
	imgDetails := tview.NewForm().
		AddDropDown("Format", []string{"qcow2", "qcow", "qed"}, 0, nil).
		AddInputField("VM Name", "", 20, nil, nil).
		AddInputField("Size", "", 5, nil, nil).
		AddDropDown("Alloc Type", []string{"G", "M"}, 0, nil).
		AddButton("Create", func() {
			app.Stop()
		}).
		AddButton("Quit", func() {
			app.Stop()
			os.Exit(0)
		})

	imgDetails.SetBorder(true).SetTitle("VM Details").SetTitleAlign(tview.AlignLeft)

	if err := app.SetRoot(imgDetails, true).SetFocus(imgDetails).Run(); err != nil {
		panic(err)
	}

	_, format := (imgDetails.GetFormItemByLabel("Format").(*tview.DropDown).GetCurrentOption())
	name := imgDetails.GetFormItemByLabel("VM Name").(*tview.InputField).GetText()
	disk, _ := strconv.ParseFloat(imgDetails.GetFormItemByLabel("Size").(*tview.InputField).GetText(), 64)
	_, allocType := imgDetails.GetFormItemByLabel("Alloc Type").(*tview.DropDown).GetCurrentOption()

	output, err := Create(format, name, disk, allocType)

	if err != nil {
		panic(err)
	}

	fmt.Println(output)
}

func Create(format string, name string, disk float64, allocType string) (string, error) {
	args := []string{
		"qemu-img",
		"create",
		"-f", format,
		fmt.Sprintf("%s.%s", name, format),
		fmt.Sprintf("%.0f%s", disk, allocType),
	}

	cmd := exec.Command(args[0], args[1:]...)

	run, err := cmd.CombinedOutput()

	return string(run), err
}
