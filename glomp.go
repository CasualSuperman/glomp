package main

import (
	"encoding/json"
	"flag"
	"fmt"
	mpd "github.com/jteeuwen/go-pkg-mpd"
	"log"
	"log/syslog"
	"os"
	"os/user"
	"strconv"
)

var config map[string]string
var client []Conn
var ErrLogger *log.Logger
var WarnLogger *log.Logger

func main() {
	client = make([]Conn, 0)
	flag.Parse()
	config = make(map[string]string)
	getConfig()

	c, err := mpd.Dial(config["address"]+":"+config["port"], config["pass"])
	if err != nil {
		log.Fatalf("Could not connect to mpd instance on %s:%s\n", config["address"], config["port"])
	}
	client = append(client, NewConn(c))
	if len(flag.Args()) == 0 {
		showGui(0)
	} else {
		action(os.Args[1:], &client[0])
	}
}

func getConfig() {
	uid := os.Getuid()

	if uid == 0 {
		ErrLogger, _ = syslog.NewLogger(syslog.LOG_ERR, log.LstdFlags)
		WarnLogger, _ = syslog.NewLogger(syslog.LOG_WARNING, log.LstdFlags)
	} else {
		logFile := os.Stderr
		ErrLogger = log.New(logFile, "Error", log.LstdFlags)
		WarnLogger = log.New(logFile, "Warning", log.LstdFlags)
	}
	usr, _ := user.LookupId(strconv.Itoa(uid))
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
