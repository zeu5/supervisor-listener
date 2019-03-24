package events

import (
	"strconv"
	"strings"
)

type EventHeader struct {
	Ver        string
	Server     string
	Serial     int64
	Pool       string
	PoolSerial int64
	Eventtype  string
	Bodylength int64
}

type Event struct {
	Header  EventHeader
	Rawbody string
	Body    map[string]string
	Type    string
}

func (e *Event) ParseBody() {
	e.Body = parseEventBody(e.Rawbody, e.Type)
}

func ParseHeader(headerstring string) (EventHeader, bool) {

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

	if !valid {
		return EventHeader{}, false
	}

	serial, err := strconv.ParseInt(headermap["serial"], 10, 64)
	if err != nil {
		return EventHeader{}, false
	}
	poolserial, err := strconv.ParseInt(headermap["poolserial"], 10, 64)
	if err != nil {
		return EventHeader{}, false
	}
	bodylength, err := strconv.ParseInt(headermap["len"], 10, 64)
	if err != nil {
		return EventHeader{}, false
	}

	return EventHeader{
		Ver:        headermap["ver"],
		Server:     headermap["server"],
		Serial:     serial,
		Pool:       headermap["pool"],
		PoolSerial: poolserial,
		Eventtype:  headermap["eventname"],
		Bodylength: bodylength,
	}, true
}
