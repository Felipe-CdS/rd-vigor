package repositories

import (
	"time"

	"github.com/google/uuid"
	"nugu.dev/rd-vigor/db"
)

type Event struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	CoverPath   string    `json:"cover_path"`
	MapsLink    string    `json:"maps_link"`
	Address     string    `json:"address"`
	Address2    string    `json:"address2"`
	City        string    `json:"city"`
	State       string    `json:"state"`
	Date        time.Time `json:"date"`
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

func (er *EventRepository) CreateEvent(e Event) *RepositoryLayerErr {

	stmt := `INSERT INTO events 
		(id, title, description, price, cover_path, maps_link, address, address2, city, state, date) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	_, err := er.EventStore.Db.Exec(
		stmt,
		uuid.New(),
		e.Title,
		e.Description,
		e.Price,
		e.CoverPath,
		e.MapsLink,
		e.Address,
		e.Address2,
		e.City,
		e.State,
		e.Date,
	)

	if err != nil {
		return &RepositoryLayerErr{err, "Insert Error"}
	}

	return nil
}

// If past == true, search old events. Otherwise search next events.
func (er *EventRepository) GetEvents(past bool) ([]Event, *RepositoryLayerErr) {

	var events []Event
	var stmt string

	if past {
		stmt = `SELECT * FROM events WHERE date < CURRENT_DATE`
	} else {
		stmt = `SELECT * FROM events WHERE date >= CURRENT_DATE`
	}

	rows, err := er.EventStore.Db.Query(stmt)

	if err != nil {
		return nil, &RepositoryLayerErr{err, "Query Error"}
	}

	for rows.Next() {
		var e Event
		if err := rows.Scan(
			&e.ID,
			&e.Title,
			&e.Description,
			&e.Price,
			&e.CoverPath,
			&e.MapsLink,
			&e.Address,
			&e.Address2,
			&e.City,
			&e.State,
			&e.Date,
		); err != nil {
			return nil, &RepositoryLayerErr{err, "Query Error"}
		}
		events = append(events, e)
	}

	return events, nil
}

func (er *EventRepository) GetEventByID(id string) (Event, *RepositoryLayerErr) {

	var e Event

	stmt := `SELECT * FROM events WHERE id = $1`

	if err := er.EventStore.Db.QueryRow(stmt, id).Scan(
		&e.ID,
		&e.Title,
		&e.Description,
		&e.Price,
		&e.CoverPath,
		&e.MapsLink,
		&e.Address,
		&e.Address2,
		&e.City,
		&e.State,
		&e.Date,
	); err != nil {
		return e, &RepositoryLayerErr{err, "Query Error"}
	}

	return e, nil
}

func (er *EventRepository) GetNextEvent() (Event, *RepositoryLayerErr) {

	var e Event

	today := time.Now()
	stmt := `SELECT * FROM events WHERE date >= $1 ORDER BY date DESC LIMIT 1;`

	if err := er.EventStore.Db.QueryRow(stmt, today).Scan(
		&e.ID,
		&e.Title,
		&e.Description,
		&e.Price,
		&e.CoverPath,
		&e.MapsLink,
		&e.Address,
		&e.Address2,
		&e.City,
		&e.State,
		&e.Date,
	); err != nil {
		return e, &RepositoryLayerErr{err, "Query Error"}
	}

	return e, nil
}
