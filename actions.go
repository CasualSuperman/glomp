package main

import (
	"fmt"
	mpd "github.com/jteeuwen/go-pkg-mpd"
)

func action(args []string, conn *Conn) {
	if len(args) == 1 {
		switch args[0] {
			case "pause":
				conn.Lock()
				conn.Pause(true)
				conn.Unlock()

			case "play":
				conn.Lock()
				conn.Pause(false)
				conn.Unlock()

			case "toggle", "t":
				conn.Lock()
				conn.Toggle()
				conn.Unlock()

			case "status", "stat", "s":
				status(conn)

			case "next", "n":
				conn.Lock()
				conn.Next()
				conn.Unlock()

			case "prev", "previous", "p":
				conn.Lock()
				conn.Previous()
				conn.Unlock()
		}
	}
}

func status(conn *Conn) {
	conn.RLock()
	defer conn.RUnlock()
	status, _ := conn.Status()

	if status.State != mpd.Stopped {
		song, err := conn.Current()
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
					fmt.Printf(" (%s)", album)
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
