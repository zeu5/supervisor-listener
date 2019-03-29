package main

import (
	"fmt"
	"os"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/zeu5/supervisor-listener/config"
	"github.com/zeu5/supervisor-listener/events"
	"github.com/zeu5/supervisor-listener/handlers"
)

var (
	wg          sync.WaitGroup
	globalprops = make(map[string]string)
)

func initListener(config *config.Config) {
	initIOBuffers()
	globalprops = config.GlobalProps
}

func processevents(eventchannel <-chan *events.Event, wg *sync.WaitGroup) {
	for event := range eventchannel {
		wg.Add(1)
		go func(event *events.Event, wg *sync.WaitGroup) {
			defer wg.Done()
			event.ParseBody()
			handlers, err := handlers.GetHandlerInstances(event)
			if err != nil {
				log.Warn(fmt.Sprintf("Error fetching handler instance : "), err)
				return
			}
			for _, handler := range handlers {
				err = handler.HandleEvent(event, globalprops)
				if err != nil {
					log.Warn(fmt.Sprintf("Error handling event: %s", event.Type), err)
				}
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
				log.Warn(err)
			}
			header, err1 := events.ParseHeader(headerstring)
			bodystring, err2 := readEventData(header.Bodylength)
			if err1 != nil {
				log.Warn(err1)
			} else if err2 != nil {
				log.Warn(err2)
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
