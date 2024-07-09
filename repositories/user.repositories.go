package repositories

import (
	"database/sql"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
	"nugu.dev/rd-vigor/db"
)

type User struct {
	ID             int       `json:"id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email"`
	OccupationArea string    `json:"occupation_area"`
	Telephone      string    `json:"telephone"`
	ReferFriend    string    `json:"refer_friend"`
	Password       string    `json:"password"`
	CreatedAt      time.Time `json:"created_at"`
}

type UserRepository struct {
	User      User
	UserStore db.Store
}

func NewUserRepository(u User, uStore db.Store) *UserRepository {
	return &UserRepository{
		User:      u,
		UserStore: uStore,
	}
}

func (ur *UserRepository) CreateUser(u User) *RepositoryLayerErr {

	stmt := `INSERT INTO users1 (first_name, last_name, email, occupation_area, password, telephone, refer_friend, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)

	if err != nil {
		return &RepositoryLayerErr{err, "Hash password Error"}
	}

	_, err = ur.UserStore.Db.Exec(
		stmt,
		u.FirstName,
		u.LastName,
		u.Email,
		u.OccupationArea,
		hashPassword,
		u.Telephone,
		u.ReferFriend,
		u.CreatedAt,
	)

	if err != nil {
		return &RepositoryLayerErr{err, "Insert Error"}
	}

	log.Printf("%s | New User Created: %s.\n", u.CreatedAt.Format(time.RFC822Z), u.Email)

	return nil
}

func (ur *UserRepository) CheckEmailExists(email string) bool {

	stmt := "SELECT email FROM users WHERE email = $1"

	queryResult := ur.UserStore.Db.QueryRow(stmt, email).Scan()

	if queryResult != sql.ErrNoRows {
		return true
	}

	return false
}
