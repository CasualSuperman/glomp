package main

import (
	"encoding/json"
	"flag"
	"fmt"
	gtk "github.com/mattn/go-gtk/gtk"
	mpd "github.com/jteeuwen/go-pkg-mpd"
	"log"
	"log/syslog"
	"os"
	"os/user"
)

var config map[string]string
var client Conn
var ErrLogger *log.Logger
var WarnLogger *log.Logger

func main() {
	ErrLogger = syslog.NewLogger(syslog.LOG_ERR, log.LstdFlags)
	WarnLogger = syslog.NewLogger(syslog.LOG_WARNING, log.LstdFlags)
	flag.Parse()
	config = make(map[string]string)
	getConfig()

	c, err := mpd.Dial(config["address"]+":"+config["port"], config["pass"])
	if err != nil {
		log.Fatalf("Could not connect to mpd instance on %s:%s\n", config["address"], config["port"])
	}
	client = NewConn(c)
	if len(flag.Args()) == 0 {
		gtk.Init(&os.Args)
		window := gtk.Window(gtk.GTK_WINDOW_TOPLEVEL)
		window.SetTitle("glomp")
		window.Connect("destroy", gtk.MainQuit)
		window.SetSizeRequest(400, 200)
		window.ShowAll()

		gtk.Main()
	} else {
		status()
	}
}

func getConfig() {
	usr, _ := user.LookupId(os.Getuid())
	var conf = usr.HomeDir + "/.config/glomp.conf"

	err := json.Unmarshal([]byte(defaults), &config)
	if err != nil {
		fmt.Println("Warning: default settings corrupted", err)
		ErrLogger.Fatalln("Warning: default settings corrupted", err)
	}

	file, err := os.Open(conf)
	if err != nil {
		WarnLogger.Println("Configuration file could not be found. Creating default.")
		file, err := os.Create(conf)
		if err != nil {
			ErrLogger.Printf("Could not create config file at (%s)", conf)
		} else {
			file.WriteString(defaults)
			file.Close()
		}
	}
	decoder := json.NewDecoder(file)
	decoder.Decode(&config)
}
