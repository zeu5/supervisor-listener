package handlers

// Handler to a supervisor event
type Handler interface {
	Run(headers interface{}, data interface{}) error
}
