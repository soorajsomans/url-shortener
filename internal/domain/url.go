package domain

import "time"

type URL struct {
	ID          int64
	LongURL     string
	ShortCode   string
	CreatedAt   time.Time
	ExpiresAt   *time.Time
	CustomAlias string
}

func (u URL) IsExpired() bool {
	if u.ExpiresAt == nil {
		return false
	}
	return time.Now().After(*u.ExpiresAt)
}
