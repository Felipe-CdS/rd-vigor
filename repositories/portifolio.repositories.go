package repositories

import (
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

func (pr *PortifolioRepository) GetUserPortifolios(u User) *RepositoryLayerErr {

	stmt := `SELECT * FROM portifolios WHERE fk_user_id = $1;`

	_, err := pr.PortifolioStore.Db.Exec(stmt, u.ID)

	if err != nil {
		return &RepositoryLayerErr{err, "Insert Error"}
	}

	return nil
}
