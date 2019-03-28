package events

// EventHeader contains the properties of the header of an event dispatched by supervisor
type EventHeader struct {
	Ver        string
	Server     string
	Serial     int64
	Pool       string
	PoolSerial int64
	Eventtype  string
	Bodylength int64
}

// Event stores the properties of an event dispatched by supervisor
type Event struct {
	Header  EventHeader
	Rawbody string
	Body    map[string]string
	Type    string
}

// ParseBody converts the string representaion of the body of the supervisor event
// into a map of key values.
func (e *Event) ParseBody() {
	e.Body = parseEventBody(e.Rawbody, e.Type)
}

// GetEventMessage returns the representation of the event as a string to be communicated
// through the handlers. It abstracts out dealing with different eventtypes
func (e *Event) GetEventMessage() (string, error) {
	return getEventMessage(e)
}
