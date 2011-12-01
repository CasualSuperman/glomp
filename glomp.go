package main

import (
	"fmt"
	"flag"
	"json"
	mpd "github.com/jteeuwen/go-pkg-mpd"
	"os"
)

var file = flag.String("c", "$HOME/.config/glomp.conf", "Configuration file.")
var addr = flag.String("a", "127.0.0.1", "IP for mpd.")
var port = flag.String("p", ":6615", "Port used by mpd.")
var pass = flag.String("pass", "", "Password for connecting to mpd.")

func main() {
	flag.Parse()
	file, err := os.Open(os.ShellExpand(*file))
	if err != nil {
		panic(err)
	}
	json.NewDecoder(file)
	client, err := mpd.Dial(*addr + *port, *pass)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(client.Status())
}
