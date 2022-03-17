package main

//Reveal2Pushover - Created by Thomas Cope - Copyright(C) Thomas Cope 2022 All Rights Reserved

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
)

var (
	conf *Config
)

func WaitForCtrlC() {
	var end_waiter sync.WaitGroup
	end_waiter.Add(1)
	var signal_channel chan os.Signal
	signal_channel = make(chan os.Signal, 1)
	signal.Notify(signal_channel, os.Interrupt)
	go func() {
		<-signal_channel
		end_waiter.Done()
	}()
	end_waiter.Wait()
}

func main() {
	fmt.Println("Reveal1Pushover - Copyright (C) Thomas Cope 2022")
	LoggingInit()
	if !ConfigExists() {
		Warn.Println("Config File not detected. Generating a new one. Please edit 'anonae_config.json'")
		WriteConfig()
		Error.Panic("No Config file to load")
	}
	c, err := LoadConfig()
	conf = &c
	if err != nil {
		Error.Panicf("Failed to load config file. Reason = %s", err)
	}
	Info.Printf("Starting Webserver...")
	server := RunServer()
	WaitForCtrlC()
	Info.Printf("Caught Control-C - Initiating Shutdown Sequence...")
	Info.Printf("Webserver Going Down...")
	if err := server.Shutdown(nil); err != nil {
		panic(fmt.Sprintf("Failed to Stop Web Server gracefully. Server was stoped forcefuly. Error Message = %s", err))
	}
	Info.Printf("Shutdown Complete")
}
