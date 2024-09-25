package repositories

import (
	"database/sql"

	"github.com/google/uuid"
	"nugu.dev/rd-vigor/db"
)

type Portifolio struct {
	Portifolio_ID string `json:"portifolio_id"`
	Fk_user_ID    string `json:"fk_user_id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
}

type PortifolioRepository struct {
	Portifolio      Portifolio
	PortifolioStore db.Store
}

func NewPortifolioRepository(p Portifolio, pStore db.Store) *PortifolioRepository {
	return &PortifolioRepository{
		Portifolio:      p,
		PortifolioStore: pStore,
	}
}

func (pr *PortifolioRepository) CreatePortifolio(u User, p Portifolio) *RepositoryLayerErr {

	stmt := `INSERT INTO portifolios 
		(portifolio_id, fk_user_id, title, description) 
		VALUES ($1, $2, $3, $4)`

	_, err := pr.PortifolioStore.Db.Exec(
		stmt,
		uuid.New(),
		u.ID,
		p.Title,
		p.Description,
	)

	if err != nil {
		return &RepositoryLayerErr{err, "Insert Error"}
	}

	return nil
}

func (pr *PortifolioRepository) EditPortifolio(u User, p Portifolio) *RepositoryLayerErr {

	var ownerID string

	stmt := "SELECT fk_user_id FROM portifolios WHERE portifolio_id = $1"

	err := pr.PortifolioStore.Db.QueryRow(stmt, p.Portifolio_ID).Scan(&ownerID)

	if err != nil {
		return &RepositoryLayerErr{sql.ErrNoRows, "Portifolio inexistente."}
	}

	if u.ID != ownerID {
		return &RepositoryLayerErr{nil, "Permissão negada."}
	}

	stmt = `UPDATE portifolios SET title = $1, description = $2
		WHERE  portifolio_id = $3;`

	_, err = pr.PortifolioStore.Db.Exec(
		stmt,
		p.Title,
		p.Description,
		p.Portifolio_ID,
	)

	if err != nil {
		return &RepositoryLayerErr{err, "Insert Error"}
	}

	return nil
}

func (pr *PortifolioRepository) GetUserPortifolios(u User) ([]Portifolio, *RepositoryLayerErr) {

	var list []Portifolio

	stmt := `SELECT portifolio_id, title, description FROM portifolios WHERE fk_user_id = $1;`

	rows, err := pr.PortifolioStore.Db.Query(stmt, u.ID)

	if err != nil {
		return nil, &RepositoryLayerErr{err, "Insert Error"}
	}

	for rows.Next() {
		var p Portifolio
		if err := rows.Scan(
			&p.Portifolio_ID,
			&p.Title,
			&p.Description,
		); err != nil {
			return nil, &RepositoryLayerErr{err, "Insert Error"}
		}
		list = append(list, p)
	}

	return list, nil
}

func (pr *PortifolioRepository) DeletePortifolio(u User, portifolioId string) *RepositoryLayerErr {

	stmt := `DELETE FROM portifolios WHERE portifolio_id = $1`

	_, err := pr.PortifolioStore.Db.Exec(stmt, portifolioId)

	if err != nil {
		return &RepositoryLayerErr{err, "Delete Error"}
	}

	return nil
}
