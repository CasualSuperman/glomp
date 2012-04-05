package main

import (
	gtk "github.com/norisatir/go-gtk3/gtk"
)

func showGui(conn int) {
		gtk.Init()
		window := gtk.NewWindow(gtk.GtkWindowType.TOPLEVEL)
		window.SetTitle("glomp")
		window.Connect("destroy", gtk.MainQuit)
		window.SetDefaultSize(400, 200)
		window.ShowAll()

		gtk.Main()
}
