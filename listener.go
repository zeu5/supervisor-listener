package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"github.com/zeu5/supervisor-listener/config"
	"github.com/zeu5/supervisor-listener/events"
)

var (
	in  *bufio.Reader
	out *bufio.Writer
	log *bufio.Writer
	wg  sync.WaitGroup
)

func initListener(config *config.Config) {
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	// Need to figure out what to do with log
}

func processevents(eventchannel <-chan *events.Event) {
	for event := range eventchannel {
		event.ParseBody()
		wg.Add(1)
		go func(event *events.Event) {
			defer wg.Done()
			// Need to find handler for event and call the handler
		}(event)
	}
}

func runListener(sigint <-chan os.Signal) {
	eventchannel := make(chan *events.Event, 10)

	go processevents(eventchannel)

	for {
		select {
		case <-sigint:
			fmt.Println("Recieved Sigint")
			close(eventchannel)
			wg.Wait()
			return
		}
		//Keep reading from stdin

	}
}
