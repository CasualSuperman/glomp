package main

import (
	gtk "github.com/mattn/go-gtk/gtk"
	"os"
)

func showGui(conn int) {
		gtk.Init(&os.Args)
		window := gtk.Window(gtk.GTK_WINDOW_TOPLEVEL)
		window.SetTitle("glomp")
		window.Connect("destroy", gtk.MainQuit)
		window.SetSizeRequest(400, 200)
		window.ShowAll()

		gtk.Main()
}
