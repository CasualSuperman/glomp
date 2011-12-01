package main

import (
	config "github.com/kless/goconfig/config"
	"fmt"
	"flag"
	mpd "github.com/jteeuwen/go-pkg-mpd"
	"os"
	"strings"
)

var file = "~/.config/glomp.conf"
var addr = flag.String("p", ":6615", "Port used by mpd.")
var pass = flag.String("pass", "", "Password for connecting to mpd.")

func main() {
	flag.Parse()
	c, err := config.ReadDefault(file)
	if err != nil {
		fmt.Println(err)
	}
	port, _ := c.String("default", "port")
	if strings.Index(port, ":") != 0 {
		port = ":" + port
	}
	*addr = port
	client, err := mpd.Dial(*addr, *pass)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(client.Status())
}
