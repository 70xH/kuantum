package kuantum

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/rivo/tview"
)

// for the moment the file path will be taken as a direct input
func WriteTUI() {

	// new app
	app := tview.NewApplication()

	// write details
	writed := tview.NewForm().
		AddInputField("OS ISO", "", 100, nil, nil).
		AddDropDown("CPU", []string{"host"}, 0, nil).
		AddInputField("Memory", "", 10, nil, nil).
		AddDropDown("Memory Type", []string{"G", "M", "K"}, 0, nil).
		AddInputField("Number of CPUs", "", 10, nil, nil).
		AddButton("Save", func() {
			app.Stop()
		}).
		AddButton("Quit", func() {
			app.Stop()
			os.Exit(0)
		})

	writed.SetBorder(true).SetTitle("Installing OS").SetTitleAlign(tview.AlignLeft)

	if err := app.SetRoot(writed, true).SetFocus(writed).Run(); err != nil {
		panic(err)
	}
}

func Write(cd string, cpu string, enable_kvm bool, memory float32, memAlloc string, numcpu int, image string, format string, sandBox string) (string, error){

	// args
	// qemu-system-x86_64 -cdrom "$rom" -cpu "$cpu_model" -enable-kvm -m "$mem$mem_alloc" -smp "$num_cpu" -vga virtio -display sdl -drive file="$image",format=$format -soundhw all --sandbox "$sandx"
	args := []string{
		"qemu-system-x86_64",
		"-cdrom", cd,
		"-cpu", cpu,
		"-m", fmt.Sprintf("%s%s", memory, memAlloc),
		"-smp", string(numcpu),
		"enable-kvm",
		"-vga", "virtio",
		"-display", "sdl",
		"-drive", fmt.Sprintf("file=%s,format=%s", image, format),
		"-soundhw", "all",
		"--sandbox", sandBox,
	}

	cmd := exec.Command(args[0], args[1:]...)

	output, err := cmd.CombinedOutput()

	return string(output), err
}
