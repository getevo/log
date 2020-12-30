package main

import (
	"github.com/getevo/log"
	"net"
)

func main()  {
	log.Debug("test")
	log.Debug(net.ParseIP("192.168.1.168").DefaultMask())
	log.Print("test")
}