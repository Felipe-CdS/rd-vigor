package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"nugu.dev/rd-vigor/db"
)

type Role string
type RegistrationStatus string

const (
	Admin       Role = "admin"
	GroupLeader Role = "group_leader"
	Member      Role = "member"
)

const (
	Pending  RegistrationStatus = "pending"
	Accepted RegistrationStatus = "accepted"
	Rejected RegistrationStatus = "rejected"
)

type User struct {
	ID                 string    `json:"id"`
	Username           string    `json:"username"`
	FirstName          string    `json:"first_name"`
	LastName           string    `json:"last_name"`
	Email              string    `json:"email"`
	OccupationArea     string    `json:"occupation_area"`
	Telephone          string    `json:"telephone"`
	ReferFriend        string    `json:"refer_friend"`
	Password           string    `json:"password"`
	Role               string    `json:"role"`
	RegistrationStatus string    `json:"registration_status"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
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

	stmt := `INSERT INTO users 
		(id, username, first_name, last_name, email, occupation_area, telephone, refer_friend, password, role, registration_status, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)

	if err != nil {
		return &RepositoryLayerErr{err, "Hash password Error"}
	}

	_, err = ur.UserStore.Db.Exec(
		stmt,
		uuid.New(),
		u.Username,
		u.FirstName,
		u.LastName,
		u.Email,
		u.OccupationArea,
		u.Telephone,
		u.ReferFriend,
		hashPassword,
		Member,
		Pending,
		time.Now(),
		time.Time{},
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

func (ur *UserRepository) CheckUsernameExists(username string) bool {

	stmt := "SELECT username FROM users WHERE username = $1"

	queryResult := ur.UserStore.Db.QueryRow(stmt, username).Scan()

	if queryResult != sql.ErrNoRows {
		return true
	}

	return false
}

func (ur *UserRepository) GetAllUsers() ([]User, error) {

	var users []User

	stmt := "SELECT id, username, first_name, last_name, email, occupation_area FROM users"

	rows, err := ur.UserStore.Db.Query(stmt)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var usr User
		if err := rows.Scan(
			&usr.ID,
			&usr.Username,
			&usr.FirstName,
			&usr.LastName,
			&usr.Email,
			&usr.OccupationArea,
		); err != nil {
			return nil, err
		}
		users = append(users, usr)
	}

	return users, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (User, *RepositoryLayerErr) {

	var usr User

	stmt := "SELECT id FROM users WHERE email = $1"

	if err := ur.UserStore.Db.QueryRow(stmt, email).Scan(&usr.ID); err != nil {
		return User{}, &RepositoryLayerErr{sql.ErrNoRows, "Usuario inexistente."}
	}
	usr, _ = ur.GetUserByID(usr.ID)

	return usr, nil
}

func (ur *UserRepository) GetUserByUsername(username string) (User, *RepositoryLayerErr) {

	var usr User

	stmt := "SELECT id FROM users WHERE username = $1"

	if err := ur.UserStore.Db.QueryRow(stmt, username).Scan(&usr.ID); err != nil {
		return User{}, &RepositoryLayerErr{sql.ErrNoRows, "Usuario inexistente."}
	}
	usr, _ = ur.GetUserByID(usr.ID)

	return usr, nil
}

func (ur *UserRepository) GetUserPasswordByID(id string) (string, *RepositoryLayerErr) {

	var password string

	stmt := "SELECT password FROM users WHERE id = $1"

	if err := ur.UserStore.Db.QueryRow(stmt, id).Scan(&password); err != nil {
		return "", &RepositoryLayerErr{sql.ErrNoRows, "Usuario inexistente."}
	}
	return password, nil
}

func (ur *UserRepository) GetUserByID(id string) (User, error) {

	var usr User

	stmt := `SELECT id
		, username
		, first_name
		, last_name
		, email
		, occupation_area
		, telephone
		, refer_friend
		, role
		, registration_status
		, created_at
		, updated_at
		FROM users WHERE id = $1`

	if err := ur.UserStore.Db.QueryRow(stmt, id).Scan(
		&usr.ID,
		&usr.Username,
		&usr.FirstName,
		&usr.LastName,
		&usr.Email,
		&usr.OccupationArea,
		&usr.Telephone,
		&usr.ReferFriend,
		&usr.Role,
		&usr.RegistrationStatus,
		&usr.CreatedAt,
		&usr.UpdatedAt,
	); err != nil {
		return usr, err
	}

	return usr, nil
}

func (ur *UserRepository) SetNewTagUser(u User, t Tag) *RepositoryLayerErr {

	stmt := `INSERT INTO users_tags (id, fk_tag_id, fk_user_id) VALUES ($1, $2, $3)`

	_, err := ur.UserStore.Db.Exec(
		stmt,
		uuid.New(),
		t.ID,
		u.ID,
	)

	if err != nil {
		return &RepositoryLayerErr{err, "Insert Error"}
	}

	return nil
}
