package events

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	// Tick5Event - String alias to the TICK_5 event type of supervisor
	Tick5Event = "TICK_5"
	// Tick60Event - String alias to the TICK_60 event type of supervisor
	Tick60Event = "TICK_60"
	// Tick3600Event - String alias to the TICK_3600 event type of supervisor
	Tick3600Event = "TICK_3600"
	// ProcessStateStartingEvent - String alias to the PROCESS_STATE_STARTING event type of supervisor
	ProcessStateStartingEvent = "PROCESS_STATE_STARTING"
	// ProcessStateRunningEvent - String alias to the PROCESS_STATE_RUNNING event type of supervisor
	ProcessStateRunningEvent = "PROCESS_STATE_RUNNING"
	// ProcessStateBackoffEvent - String alias to the PROCESS_STATE_BACKOFF event type of supervisor
	ProcessStateBackoffEvent = "PROCESS_STATE_BACKOFF"
	// ProcessStateStoppingEvent - String alias to the PROCESS_STATE_STOPPING event type of supervisor
	ProcessStateStoppingEvent = "PROCESS_STATE_STOPPING"
	// ProcessStateExitedEvent - String alias to the PROCESS_STATE_EXITED event type of supervisor
	ProcessStateExitedEvent = "PROCESS_STATE_EXITED"
	// ProcessStateStoppedEvent - String alias to the PROCESS_STATE_STOPPED event type of supervisor
	ProcessStateStoppedEvent = "PROCESS_STATE_STOPPED"
	// ProcessStateFatalEvent - String alias to the PROCESS_STATE_FATAL event type of supervisor
	ProcessStateFatalEvent = "PROCESS_STATE_FATAL"
	// ProcessStateUnknownEvent - String alias to the PROCESS_STATE_UNKNOWN event type of supervisor
	ProcessStateUnknownEvent = "PROCESS_STATE_UNKNOWN"
	// RemoteCommunicationEvent - String alias to the REMOTE_COMMUNICATION event type of supervisor
	RemoteCommunicationEvent = "REMOTE_COMMUNICATION"
	// ProcessLogStdoutEvent - String alias to the PROCESS_LOG_STDOUT event type of supervisor
	ProcessLogStdoutEvent = "PROCESS_LOG_STDOUT"
	// ProcessLogStderrEvent - String alias to the PROCESS_LOG_STDERR event type of supervisor
	ProcessLogStderrEvent = "PROCESS_LOG_STDERR"
	// ProcessCommunicationStdoutEvent - String alias to the PROCESS_COMMUNICATION_STDOUT event type of supervisor
	ProcessCommunicationStdoutEvent = "PROCESS_COMMUNICATION_STDOUT"
	// ProcessCommunicationStderrEvent - String alias to the PROCESS_COMMUNICATION_STDERR event type of supervisor
	ProcessCommunicationStderrEvent = "PROCESS_COMMUNICATION_STDERR"
	// SupervisorStateChanceRunningEvent - String alias to the SUPERVISOR_STATE_CHANGE_RUNNING event type of supervisor
	SupervisorStateChanceRunningEvent = "SUPERVISOR_STATE_CHANGE_RUNNING"
	// SupervisorStateChancestoppingEvent - String alias to the SUPERVISOR_STATE_CHANGE_STOPPING event type of supervisor
	SupervisorStateChancestoppingEvent = "SUPERVISOR_STATE_CHANGE_STOPPING"
	// ProcessGroupAddedEvent - String alias to the PROCESS_GROUP_ADDED event type of supervisor
	ProcessGroupAddedEvent = "PROCESS_GROUP_ADDED"
	// ProcessGroupRemovedEvent - String alias to the PROCESS_GROUP_REMOVED event type of supervisor
	ProcessGroupRemovedEvent = "PROCESS_GROUP_REMOVED"
)

// ParseHeader parses the string representaion of the supervisor event header string and validates it
func ParseHeader(headerstring string) (EventHeader, bool) {

	emptyEventHeader := EventHeader{
		Bodylength: 0,
	}

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
		return emptyEventHeader, false
	}

	serial, err := strconv.ParseInt(headermap["serial"], 10, 64)
	if err != nil {
		return emptyEventHeader, false
	}
	poolserial, err := strconv.ParseInt(headermap["poolserial"], 10, 64)
	if err != nil {
		return emptyEventHeader, false
	}
	bodylength, err := strconv.ParseInt(headermap["len"], 10, 64)
	if err != nil {
		return emptyEventHeader, false
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

func parseEventBody(bodystring, eventtype string) map[string]string {
	eventmap := make(map[string]string)

	switch eventtype {
	case Tick5Event, Tick60Event, Tick3600Event:
		a := strings.Split(bodystring, ":")
		eventmap[a[0]] = a[1]
	case ProcessStateBackoffEvent, ProcessStateExitedEvent, ProcessStateFatalEvent, ProcessStateRunningEvent, ProcessStateStartingEvent, ProcessStateStoppedEvent, ProcessStateStoppingEvent, ProcessStateUnknownEvent:
		keyvalues := strings.Split(bodystring, " ")
		for _, keyvalue := range keyvalues {
			s := strings.Split(keyvalue, ":")
			eventmap[s[0]] = s[1]
		}
	case RemoteCommunicationEvent:
		lines := strings.Split(bodystring, "\n")
		keyvalues := strings.Split(lines[0], ":")
		eventmap[keyvalues[0]] = keyvalues[1]
		eventmap["data"] = lines[1]
	case ProcessLogStderrEvent, ProcessLogStdoutEvent, ProcessCommunicationStderrEvent, ProcessCommunicationStdoutEvent:
		lines := strings.Split(bodystring, "\n")
		for _, keyvalues := range strings.Split(lines[0], " ") {
			s := strings.Split(keyvalues, ":")
			eventmap[s[0]] = s[1]
		}
		eventmap["data"] = lines[1]
	case ProcessGroupAddedEvent, ProcessGroupRemovedEvent:
		s := strings.Split(bodystring, ":")
		eventmap[s[0]] = s[1]
	}
	return eventmap
}

func getEventMessage(event *Event) (string, error) {
	var message string
	errEventBody := fmt.Errorf("Could not parse event body data")
	switch event.Type {
	case Tick5Event, Tick60Event, Tick3600Event:
		if epochstring, ok := event.Body["when"]; ok {
			epoch, err := strconv.ParseInt(epochstring, 10, 64)
			if err != nil {
				return "", errEventBody
			}
			t := time.Unix(epoch, 0)
			message = fmt.Sprintf("Event %s at %s", event.Type, t.String())
		}
	case ProcessStateBackoffEvent, ProcessStateExitedEvent, ProcessStateFatalEvent, ProcessStateRunningEvent, ProcessStateStartingEvent, ProcessStateStoppedEvent, ProcessStateStoppingEvent, ProcessStateUnknownEvent:
		processname, ok1 := event.Body["processname"]
		fromstate, ok2 := event.Body["from_state"]
		if ok1 && ok2 {
			curstate := strings.Split(event.Type, "_")[2]
			message = fmt.Sprintf("Process %s transitioned from state %s to %s", processname, fromstate, curstate)
		} else {
			return "", errEventBody
		}
	case RemoteCommunicationEvent:
		comtype, ok1 := event.Body["type"]
		data, ok2 := event.Body["data"]
		if ok1 && ok2 {
			message = fmt.Sprintf("Remote communication of type %s was sent with data %s", comtype, data)
		} else {
			return "", errEventBody
		}
	case ProcessLogStderrEvent, ProcessLogStdoutEvent:
		processname, ok1 := event.Body["processname"]
		data, ok2 := event.Body["data"]
		if ok1 && ok2 {
			out := strings.Split(event.Type, "_")[2]
			message = fmt.Sprintf("Process %s output \"%s\" to its %s", processname, data, out)
		} else {
			return "", errEventBody
		}
	case ProcessCommunicationStderrEvent, ProcessCommunicationStdoutEvent:
		processname, ok1 := event.Body["processname"]
		data, ok2 := event.Body["data"]
		if ok1 && ok2 {
			out := strings.Split(event.Type, "_")[2]
			message = fmt.Sprintf("Process %s communicated \"%s\" to supervisor on %s", processname, data, out)
		} else {
			return "", errEventBody
		}
	case SupervisorStateChanceRunningEvent:
		message = "Supervisor started running"
	case SupervisorStateChancestoppingEvent:
		message = "Supervisor has stopped"
	}
	return message, nil
}
