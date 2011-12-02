package main

import (
	"fmt"
	mpd "github.com/jteeuwen/go-pkg-mpd"
)

func status() {
	client.RLock()
	defer client.RUnlock()
	status, _ := client.Status()

	if status.State != mpd.Stopped {
		song, err := client.Current()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Now Playing: ")

			file := song["file"]
			title := song["Title"]
			artist := song["Artist"]
			album := song["Album"]

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
