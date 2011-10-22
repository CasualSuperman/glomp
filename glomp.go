package main

import (
	"fmt"
	"flag"
	mpd "github.com/jteeuwen/go-pkg-mpd"
	"os"
)

var addr *string = flag.String("p", "Port", ":6060")
var pass = flag.String("pass", "Password for connecting to mpd.", "")

func main() {
	flag.Parse()
	client, err := mpd.Dial(*addr, *pass)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(client.Status())
}
