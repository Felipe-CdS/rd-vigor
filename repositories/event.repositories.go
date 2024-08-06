package repositories

import "nugu.dev/rd-vigor/db"

type Event struct {
}
type EventRepository struct {
	Event      Event
	EventStore db.Store
}

func NewEventRepository(e Event, eStore db.Store) *EventRepository {
	return &EventRepository{
		Event:      e,
		EventStore: eStore,
	}
}
