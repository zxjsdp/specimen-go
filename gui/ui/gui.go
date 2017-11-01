package ui

import (
	"strconv"

	"github.com/ProtonMail/ui"
)

func main() {
	err := ui.Main(func() {
		entry1 := ui.NewEntry()
		combo1 := ui.NewCombobox()
		combo1.Append("the dog")
		combo1.Append("jumped over")
		combo1.Append("the lazy dog")
		label1 := ui.NewLabel(" query file:  ")

		entry2 := ui.NewEntry()
		combo2 := ui.NewCombobox()
		combo2.Append("the dog")
		combo2.Append("jumped over")
		combo2.Append("the lazy dog")
		label2 := ui.NewLabel(" data file:  ")

		entry3 := ui.NewEntry()
		label3 := ui.NewLabel(" output file:  ")

		greeting := ui.NewLabel("-")

		multiLineEntry := ui.NewMultilineEntry()

		grid := ui.NewGrid()

		grid.Append(label1, 0, 0, 1, 1, false, ui.AlignFill, false, ui.AlignFill)
		grid.Append(entry1, 1, 0, 1, 1, true, ui.AlignFill, false, ui.AlignFill)
		grid.Append(combo1, 2, 0, 1, 1, false, ui.AlignFill, false, ui.AlignFill)

		grid.Append(label2, 0, 1, 1, 1, false, ui.AlignFill, false, ui.AlignFill)
		grid.Append(entry2, 1, 1, 1, 1, true, ui.AlignFill, false, ui.AlignFill)
		grid.Append(combo2, 2, 1, 1, 1, false, ui.AlignFill, false, ui.AlignFill)

		grid.Append(label3, 0, 2, 1, 1, false, ui.AlignFill, false, ui.AlignFill)
		grid.Append(entry3, 1, 2, 2, 1, true, ui.AlignFill, false, ui.AlignFill)

		grid.Append(greeting, 0, 3, 9, 1, true, ui.AlignFill, false, ui.AlignFill)

		grid.Append(multiLineEntry, 0, 4, 9, 1, true, ui.AlignFill, true, ui.AlignFill)

		window := ui.NewWindow("Specimen Info", 800, 600, false)
		window.SetChild(grid)

		combo1.OnSelected(func(*ui.Combobox) {
			selectedIndex := combo1.Selected()
			entry1.SetText(strconv.Itoa(selectedIndex))
		})

		combo2.OnSelected(func(*ui.Combobox) {
			selectedIndex := combo2.Selected()
			entry2.SetText(strconv.Itoa(selectedIndex))
		})

		//button1.OnClicked(func(*ui.Button) {
		//	greeting.SetText("Hello, " + entry1.Text() + "!")
		//})

		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}
