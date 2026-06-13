package event

type EventType string

const (
	URLCreated EventType = "URL_CREATED"
	URLVisited EventType = "URL_VISITED"
)
