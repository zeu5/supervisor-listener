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
			return t.String(), nil
		}
	case ProcessStateBackoffEvent, ProcessStateExitedEvent, ProcessStateFatalEvent, ProcessStateRunningEvent, ProcessStateStartingEvent, ProcessStateStoppedEvent, ProcessStateStoppingEvent, ProcessStateUnknownEvent:

	}
	return message, nil
}
