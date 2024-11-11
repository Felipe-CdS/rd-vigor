package repositories

import (
	"database/sql"
	"time"

	"nugu.dev/rd-vigor/db"
)

type Meeting struct {
	User1     string    `json:"fk_user1_id"`
	User2     string    `json:"fk_user2_id"`
	Timestamp time.Time `json:"meeting_timestamp"`
}

type MeetingRepository struct {
	Meeting      Meeting
	MeetingStore db.Store
}

func NewMeetingRepository(m Meeting, mStore db.Store) *MeetingRepository {
	return &MeetingRepository{
		Meeting:      m,
		MeetingStore: mStore,
	}
}

func (mr *MeetingRepository) CreateMeeting(m Meeting) *RepositoryLayerErr {

	var possibleDuplicate sql.NullTime

	stmt := `SELECT meeting_timestamp FROM meetings WHERE fk_user1_id = $1 AND fk_user2_id = $2;`

	if err := mr.MeetingStore.Db.QueryRow(stmt, m.User2, m.User1).Scan(
		&possibleDuplicate,
	); err != nil {
		return &RepositoryLayerErr{err, "Insert Error"}
	}

	if possibleDuplicate.Valid {
		return &RepositoryLayerErr{nil, "Duplicate found."}
	}

	stmt = `INSERT INTO meetings
			(fk_user1_id, fk_user2_id, meeting_timestamp) 
			VALUES ($1, $2, $3)`

	_, err := mr.MeetingStore.Db.Exec(
		stmt,
		m.User1,
		m.User2,
		m.Timestamp,
	)

	if err != nil {
		return &RepositoryLayerErr{err, "Insert Error"}
	}

	return nil
}
