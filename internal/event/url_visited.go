package event

import "time"

type URLVisitedEvent struct {
	EventType EventType `json:"event_type"`
	URLID     int64     `json:"url_id"`
	ShortCode string    `json:"short_code"`
	VisitedAt time.Time `json:"visited_at"`
}
