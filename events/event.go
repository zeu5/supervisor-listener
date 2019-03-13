package events

type Event struct {
	header map[string]string
	body   map[string]string
}

func GetEvent(header map[string]string, event string) *Event {
	eventtype := header["eventname"]
	body := parseBody(event, eventtype)
	return &Event{
		header: header,
		body:   body,
	}
}

func parseBody(bodystring string, eventtype string) map[string]string {
	var body map[string]string
	switch eventtype {
	default:

	}
	return body
}
