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
var client Conn

func main() {
	flag.Parse()
	config = make(map[string]string)
	getConfig()

	c, err := mpd.Dial(config["address"]+":"+config["port"], config["pass"])
	if err != nil {
		fmt.Printf("Could not connect to mpd instance on %s:%s\n", config["address"], config["port"])
		os.Exit(1)
	}
	client = NewConn(c)
	status()
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
