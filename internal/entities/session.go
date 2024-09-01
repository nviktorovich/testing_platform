package entities

import "time"

type Session struct {
	SID       int
	UID       int
	CreatedAt time.Time
	TTL       time.Duration
	Topics    []Topic
	Question  []Question
	//State     State
}
