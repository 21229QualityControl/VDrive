package main

import (
	"github.com/andlabs/ui"
)

var isHosting = false
var prt int

var btn *ui.Button
var hostname *ui.Entry
var roomname *ui.Entry
var port *ui.Spinbox

func host() {
	// Create window
	win = ui.NewWindow("RemoDrive", 750, 600, false)
	win.OnClosing(func(*ui.Window) bool {
		if isHosting {
			cleanup()
		}
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		win.Destroy()
		return true
	})
	win.SetMargined(true)

	// Create form
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	group := ui.NewGroup("Settings")
	group.SetMargined(true)
	form := ui.NewForm()
	form.SetPadded(true)

	// Create form items
	roomname = ui.NewEntry()
	form.Append("Room", roomname, false)

	hostname = ui.NewEntry()
	hostname.SetText("localhost")
	form.Append("Host", hostname, false)

	port = ui.NewSpinbox(0, 65535)
	port.SetValue(8080)
	port.OnChanged(func(spin *ui.Spinbox) {
		// Make readonly
		if isHosting {
			spin.SetValue(prt)
		}
	})
	form.Append("Port", port, false)

	// Create listen btn
	btn = ui.NewButton("Host")

	// Disable if empty
	if len(roomname.Text()) == 0 {
		btn.Disable()
	}
	roomname.OnChanged(func(*ui.Entry) {
		if len(roomname.Text()) > 0 && !btn.Enabled() {
			btn.Enable()
		} else if len(roomname.Text()) == 0 && btn.Enabled() {
			btn.Disable()
		}
	})

	// onclick
	btn.OnClicked(func(*ui.Button) {
		if !isHosting {
			hostname.SetReadOnly(true)
			roomname.SetReadOnly(true)

			prt = port.Value()
			room = roomname.Text()

			btn.Disable()
			btn.SetText("Connecting...")

			go func() {
				listen()
				ui.QueueMain(func() {
					btn.Enable()
					btn.SetText("Stop")
				})
			}()
			return
		}

		cleanup()
	})

	// Make names
	nameModel = &NameModel{}
	uiNameModel := ui.NewTableModel(nameModel)
	nameModel.m = uiNameModel
	table := ui.NewTable(&ui.TableParams{
		Model: uiNameModel,
	})
	table.AppendTextColumn("Name", 0, ui.TableModelColumnNeverEditable, nil) // Name column

	nameGroup := ui.NewGroup("People")
	nameGroup.SetChild(table)

	// Make binds
	bindsModel = &BindsModel{}
	uiBindsModel := ui.NewTableModel(bindsModel)
	bindsModel.m = uiBindsModel
	table = ui.NewTable(&ui.TableParams{
		Model: uiBindsModel,
	})
	table.AppendTextColumn("Key", 0, ui.TableModelColumnNeverEditable, nil)         // Key column
	table.AppendTextColumn("Description", 1, ui.TableModelColumnNeverEditable, nil) // Description column

	bindsGroup := ui.NewGroup("Binds")
	bindsGroup.SetChild(table)

	// Make UI
	group.SetChild(form)
	vbox.Append(group, false)
	vbox.Append(nameGroup, true)
	vbox.Append(bindsGroup, true)
	vbox.Append(btn, false)
	win.SetChild(vbox)
	win.Show()
}
