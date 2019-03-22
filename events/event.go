package events

type EventHeader struct {
	Ver        string
	Server     string
	Serial     int
	Pool       string
	PoolSerial int
	Eventtype  string
	Len        int
}

type Event struct {
	Header  EventHeader
	Rawbody string
	Body    map[string]string
}

func (e *Event) ParseBody() {
	// Need to parse rawbody based on event type
}
