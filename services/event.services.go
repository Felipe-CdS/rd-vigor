package services

import (
	"net/http"

	"nugu.dev/rd-vigor/repositories"
)

type EventRepository interface {
	CreateEvent(e repositories.Event) *repositories.RepositoryLayerErr

	GetEvents(past bool) ([]repositories.Event, *repositories.RepositoryLayerErr)
	GetEventByID(id string) (repositories.Event, *repositories.RepositoryLayerErr)
	GetNextEvent() (repositories.Event, *repositories.RepositoryLayerErr)
}

type EventService struct {
	Repository EventRepository
}

func NewEventService(er EventRepository) *EventService {
	return &EventService{
		Repository: er,
	}
}

func (es *EventService) CreateEvent(e repositories.Event) *ServiceLayerErr {

	if err := es.Repository.CreateEvent(e); err != nil {
		return &ServiceLayerErr{nil, "Error Creating Event.", http.StatusInternalServerError}
	}

	return nil
}

func (es *EventService) GetEvents(past bool) ([]repositories.Event, *ServiceLayerErr) {

	users, err := es.Repository.GetEvents(past)

	if err != nil {
		return nil, &ServiceLayerErr{err.Error, "Query Err", http.StatusInternalServerError}
	}

	return users, nil
}

func (es *EventService) GetEventByID(id string) (repositories.Event, *ServiceLayerErr) {

	e, err := es.Repository.GetEventByID(id)

	if err != nil {
		return e, &ServiceLayerErr{err.Error, "Query Err", http.StatusInternalServerError}
	}

	return e, nil
}

func (es *EventService) GetNextEvent() (repositories.Event, *ServiceLayerErr) {

	e, err := es.Repository.GetNextEvent()

	if err != nil {
		return e, &ServiceLayerErr{err.Error, "Query Err", http.StatusInternalServerError}
	}

	return e, nil
}
