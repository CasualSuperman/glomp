package main

import (
	"encoding/json"
	"flag"
	"fmt"
	mpd "github.com/jteeuwen/go-pkg-mpd"
	"os"
	"os/user"
)

var config map[string]string

func main() {
	flag.Parse()
	config = make(map[string]string)
	getConfig()

	client, err := mpd.Dial(config["address"]+":"+config["port"], config["pass"])
	if err != nil {
		fmt.Printf("Could not connect to mpd instance on %s:%s\n", config["address"], config["port"])
		os.Exit(1)
	}
	status, err := client.Status()
	if status.State != mpd.Stopped {
		song, err := client.Current()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Now Playing: ")

			file := song.S("file")
			title := song.S("Title")
			artist := song.S("Artist")
			album := song.S("Album")

			if title == "" {
				fmt.Print(file)
			} else {
				fmt.Print(title)
				if artist != "" {
					fmt.Printf(" - %s", artist)
				}
				if album != "" {
					fmt.Printf(" from %s", album)
				}
			}

			if status.State == mpd.Paused {
				fmt.Print(" [Paused]")
			}
			fmt.Println()
		}
	} else {
		fmt.Println("Nothing playing.")
	}
}

func getConfig() {
	usr, _ := user.LookupId(os.Getuid())
	var conf = usr.HomeDir + "/.config/glomp.conf"

	err := json.Unmarshal([]byte(defaults), &config)
	if err != nil {
		fmt.Println("Warning: default settings corrupted", err)
	}

	file, err := os.Open(conf)
	if err != nil {
		fmt.Println("Configuration file could not be found. Continuing with default settings...")
	}
	decoder := json.NewDecoder(file)
	decoder.Decode(&config)
}
