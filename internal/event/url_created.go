package event

import "time"

type URLCreatedEvent struct {
	EventType EventType `json:"event_type"`
	URLID     int64     `json:"url_id"`
	ShortCode string    `json:"short_code"`
	LongURL   string    `json:"long_url"`
	CreatedAt time.Time `json:"created_at"`
}
