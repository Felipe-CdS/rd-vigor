package services

import (
	"net/http"
	"time"

	"nugu.dev/rd-vigor/repositories"
)

type MeetingRepository interface {
	CreateMeeting(m repositories.Meeting) *repositories.RepositoryLayerErr
}

type MeetingService struct {
	Repository MeetingRepository
}

func NewMeetingService(mr MeetingRepository) *MeetingService {
	return &MeetingService{
		Repository: mr,
	}
}

func (ms *MeetingService) CreateMeeting(user1ID string, user2ID string, timestamp time.Time) *ServiceLayerErr {

	newMeeting := repositories.Meeting{
		User1:     user1ID,
		User2:     user2ID,
		Timestamp: timestamp,
	}

	err := ms.Repository.CreateMeeting(newMeeting)

	if err != nil {
		return &ServiceLayerErr{err.Error, "Query Err", http.StatusInternalServerError}
	}

	return nil
}
