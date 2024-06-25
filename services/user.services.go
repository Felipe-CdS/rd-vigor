package services

import "nugu.dev/rd-vigor/db"

func NewUserService(u User, uStore db.Store) *UserService {
	return &UserService{
		User:      u,
		UserStore: uStore,
	}
}

type User struct {
	ID             int    `json:"id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	OccupationArea string `json:"occupation_area"`
	Password       string `json:"password"`
	CreatedAt      int    `json:"created_at"`
}

type UserService struct {
	User      User
	UserStore db.Store
}

func (us *UserService) CreateUser(u User) error {

	stmt := `INSERT INTO users (first_name, last_name, email, occupation_area, password, created_at) VALUES($1, $2, $3, $4, $5, $6)`

	_, err := us.UserStore.Db.Exec(
		stmt,
		u.FirstName,
		u.LastName,
		u.Email,
		u.OccupationArea,
		u.Password,
		u.CreatedAt,
	)

	return err
}
