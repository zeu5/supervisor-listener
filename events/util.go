package events

const (
	TICK_5_EVENT = "TICK_5"
)

func parseEventBody(bodystring, eventtype string) (map[string]string, error) {
	eventmap := make(map[string]string)

	switch eventtype {
	case TICK_5_EVENT:

	}
	return eventmap, nil
}
