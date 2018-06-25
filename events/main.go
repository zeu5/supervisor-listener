package events

type SupervisorEvent interface {
	Headers() map[string]string
	Data() map[string]string
}
