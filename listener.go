package main

import (
	"os"
	"sync"

	"github.com/zeu5/supervisor-listener/config"
	"github.com/zeu5/supervisor-listener/events"
	"github.com/zeu5/supervisor-listener/handlers"
)

var (
	wg sync.WaitGroup
)

func initListener(config *config.Config) {
	initIOBuffers()
	// Need to figure out what to do with log
}

func processevents(eventchannel <-chan *events.Event, wg *sync.WaitGroup) {
	for event := range eventchannel {
		wg.Add(1)
		go func(event *events.Event, wg *sync.WaitGroup) {
			defer wg.Done()
			event.ParseBody()
			handler, err := handlers.GetHandlerInstance(event)
			if err != nil {
				// Log error while getting handler
				return
			}
			err = handler.HandleEvent(event)
			if err != nil {
				// Log error when handling event
			}
		}(event, wg)
	}
}

func runListener(sigint <-chan os.Signal) {
	eventchannel := make(chan *events.Event, 10)

	go processevents(eventchannel, &wg)
	replyReady()

	for {
		select {
		case <-sigint:
			close(eventchannel)
			wg.Wait()
			return
		default:
			headerstring, err := readHeaderData()
			if err != nil {
				// Need to log error
			}
			if headerstring != "" {
				header, ok := events.ParseHeader(headerstring)
				bodystring, err := readEventData(header.Bodylength)
				if err != nil || !ok {
					// Log error
				} else {
					eventchannel <- &events.Event{
						Header:  header,
						Rawbody: bodystring,
						Type:    header.Eventtype,
					}
				}
				replyOk()
			}
		}
	}

}
