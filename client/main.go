package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

var win *ui.Window

func handle(err error) {
	if err != nil {
		done := make(chan struct{})
		ui.QueueMain(func() {
			ui.MsgBoxError(win, "Error!", err.Error())
			win.Destroy()
			ui.Quit()
			done <- struct{}{}
		})
		<-done // So that the error stays in original scope
		panic(err)
	}
}

func main() {
	ui.Main(host)
}
