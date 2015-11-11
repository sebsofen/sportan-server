// Package sportan package for sportan server
package main

import "sportan/srvr"

func main() {
	cnfg := srvr.NewConfiguration("config.json")

	server := srvr.NewAppServer(*cnfg)
	server.Run()
}
