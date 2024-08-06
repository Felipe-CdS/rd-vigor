package services

type EventRepository interface {
}

type EventService struct {
	Repository EventRepository
}

func NewEventService(er EventRepository) *EventService {
	return &EventService{
		Repository: er,
	}
}
