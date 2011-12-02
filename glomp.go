package main

import (
	"encoding/json"
	"flag"
	"fmt"
	mpd "github.com/jteeuwen/go-pkg-mpd"
	"os"
	"os/user"
	fp "path/filepath"
)

var config map[string]string

func main() {
	flag.Parse()
	config = make(map[string]string)
	getConfig()

	client, err := mpd.Dial(config["address"]+":"+config["port"], config["pass"])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	status, err := client.Status()
	if status.State == mpd.Playing {
		song, err := client.Current()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Now Playing: %s - %s from %s\n", song.S("Title"), song.S("Artist"), song.S("Album"))
		}
	} else {
		fmt.Println("Nothing playing.")
	}

}

func getConfig() {
	usr, _ := user.LookupId(os.Getuid())
	var conf = usr.HomeDir + "/.config/glomp.conf"

	if len(flag.Args()) == 1 {
		conf = flag.Args()[0]
	}

	confs, err := fp.Glob(conf)
	if len(confs) == 0 || err != nil {
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Configuration file could not be found at %s.", conf)
		}
		os.Exit(1)
	}

	err = json.Unmarshal([]byte(defaults), &config)
	if err != nil {
		fmt.Println(err)
	}

	file, _ := os.Open(confs[0])
	decoder := json.NewDecoder(file)
	decoder.Decode(&config)
}
