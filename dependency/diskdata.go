package dependency

import "encoding/json"



// POJO to wrap data stored in disk
type DiskData struct {
	topic string  `json:"topic"`
	bData string  `json:"bData"`
}

type Event struct {
	Header  EventHeader     `json:"header"`
	Message json.RawMessage `json:"event"`
}

// EventHeader header part of the event
type EventHeader struct {
	Name          string `json:"name"`
	SchemaVersion string `json:"schemaVersion"`
	AppName       string `json:"appName"`
	AppVersion    string `json:"appVersion"`
	ExperimentID  string `json:"experimentId"`
	UUID          string `json:"uuid"`
	EventID       int64  `json:"eventId"`
	Timestamp     int64  `json:"timestamp"`
}