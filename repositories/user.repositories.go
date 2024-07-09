package repositories

import (
	"database/sql"
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

	stmt := `INSERT INTO users1 (first_name, last_name, email, occupation_area, password, telephone, refer_friend, created_at) 
			VALUES($1, $2, $3, $4, $5, $6, $7, $8)`

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

func (ur *UserRepository) GetAllUsers() ([]User, error) {

	var users []User

	stmt := "SELECT id, first_name, last_name, email, occupation_area FROM users1"

	rows, err := ur.UserStore.Db.Query(stmt)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var usr User
		if err := rows.Scan(&usr.ID, &usr.FirstName, &usr.LastName, &usr.Email, &usr.OccupationArea); err != nil {
			return nil, err
		}
		users = append(users, usr)
	}

	return users, nil
}

func (ur *UserRepository) GetUserByID(id int) (User, error) {

	var usr User

	stmt := "SELECT id, first_name, last_name, email, occupation_area, telephone, created_at, refer_friend FROM users1 WHERE id = $1"

	if err := ur.UserStore.Db.QueryRow(stmt, id).Scan(
		&usr.ID,
		&usr.FirstName,
		&usr.LastName,
		&usr.Email,
		&usr.OccupationArea,
		&usr.Telephone,
		&usr.CreatedAt,
		&usr.ReferFriend,
	); err != nil {
		return usr, err
	}

	return usr, nil
}
