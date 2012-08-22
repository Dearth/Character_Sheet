package main

import(
	"github.com/mattn/go-gtk/gtk"
	"os"
)

func main () {
	
	gtk.Init(&os.Args)
	window := gtk.Window(gtk.GTK_WINDOW_TOPLEVEL)

	window.ShowAll()
	gtk.Main()
}
