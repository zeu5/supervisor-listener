package events

import (
	"fmt"
	"strconv"
	"strings"
)

type EventHeader struct {
	Ver        string
	Server     string
	Serial     int
	Pool       string
	PoolSerial int
	Eventtype  string
	Bodylength int
}

type Event struct {
	Header  EventHeader
	Rawbody string
	Body    map[string]string
	Type    string
}

func (e *Event) ParseBody() error {
	eventbody, err := parseEventBody(e.Rawbody, e.Type)
	if err != nil {
		return err
	}
	e.Body = eventbody
	return nil
}

func ParseHeader(headerstring string) (EventHeader, error) {

	requiredkeys := []string{"ver", "server", "serial", "pool", "poolserial", "eventname", "len"}
	headermap := make(map[string]string)
	for _, keyvalue := range strings.Split(headerstring, " ") {
		if strings.Contains(keyvalue, ":") {
			s := strings.Split(keyvalue, ":")
			headermap[s[0]] = s[1]
		}
	}

	valid := true
	for _, key := range requiredkeys {
		if _, ok := headermap[key]; !ok {
			valid = false
			break
		}
	}

	headerErr := fmt.Errorf("Error parsing header data")
	if !valid {
		return EventHeader{}, headerErr
	}

	serial, err := strconv.Atoi(headermap["serial"])
	if err != nil {
		return EventHeader{}, headerErr
	}
	poolserial, err := strconv.Atoi(headermap["poolserial"])
	if err != nil {
		return EventHeader{}, headerErr
	}
	bodylength, err := strconv.Atoi(headermap["len"])
	if err != nil {
		return EventHeader{}, headerErr
	}

	return EventHeader{
		Ver:        headermap["ver"],
		Server:     headermap["server"],
		Serial:     serial,
		Pool:       headermap["pool"],
		PoolSerial: poolserial,
		Eventtype:  headermap["eventname"],
		Bodylength: bodylength,
	}, nil
}
