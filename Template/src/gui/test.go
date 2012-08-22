package main

import(
	"github.com/mattn/go-gtk/gtk"
	"os"
)

func main () {

	gtk.Init(&os.Args)
	window := gtk.Window(gtk.GTK_WINDOW_TOPLEVEL)

	window.Connect("delete-event", func(){
		println("Delete-event occured\n")
		})

	window.Connect("destroy", func(){ gtk.MainQuit()})

	button := gtk.ButtonWithLabel("Hello, World")

	button.Clicked(func() {
		println("Hello, World\n")
	})

	button.Clicked(func() {
		gtk.MainQuit()
	})

	window.Add(button)
	window.ShowAll()
	gtk.Main()
}
