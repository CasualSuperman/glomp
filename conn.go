package main

import (
	mpd "github.com/jteeuwen/go-pkg-mpd"
	"sync"
)

type Conn struct {
	sync.RWMutex
	*mpd.Client
}

func NewConn(c *mpd.Client) (conn Conn) {
	conn.Client = c
	return
}
