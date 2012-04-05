package main

import (
	"fmt"
	mpd "github.com/jteeuwen/go-pkg-mpd"
)

func action(args []string, conn int) {
	mp := client[conn]
	if len(args) == 1 {
		switch args[0] {
			case "pause":
				mp.Lock()
				mp.Pause(true)
				mp.Unlock()

			case "play":
				mp.Lock()
				mp.Pause(false)
				mp.Unlock()

			case "toggle", "t":
				mp.Lock()
				mp.Toggle()
				mp.Unlock()

			case "status", "stat", "s":
				status(conn)

			case "next", "n":
				mp.Lock()
				mp.Next()
				mp.Unlock()

			case "prev", "previous", "p":
				mp.Lock()
				mp.Previous()
				mp.Unlock()
		}
	}
}

func status(conn int) {
	client[conn].RLock()
	defer client[conn].RUnlock()
	status, _ := client[conn].Status()

	if status.State != mpd.Stopped {
		song, err := client[conn].Current()
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
