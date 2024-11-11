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
	ID                 string `json:"id"`
	Username           string `json:"username"`
	Password           string `json:"password"`
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	Email              string `json:"email"`
	OccupationArea     string `json:"occupation_area"`
	Telephone          string `json:"telephone"`
	ReferFriend        string `json:"refer_friend"`
	Role               string `json:"role"`
	RegistrationStatus string `json:"registration_status"`

	ProfilePic         string `json:"profile_picture"`
	ProfileDescription string `json:"profile_description"`
	CompanyLogo        string `json:"company_logo"`
	CompanyName        string `json:"company_name"`
	MainProduct        string `json:"main_product"`
	PresentationVideo  string `json:"presentation_video"`
	Resume             string `json:"resume"`

	Address  string `json:"address"`
	Address2 string `json:"address2"`
	City     string `json:"city"`
	State    string `json:"state"`
	Zipcode  string `json:"zipcode"`

	StripeID              string    `json:"stripe_id"`
	SubscriptionStatus    bool      `json:"subscription_status"`
	SubscriptionExpiresAt time.Time `json:"subsctription_expires_at"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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

func (ur *UserRepository) UpdateUser(u User, newUserData User) *RepositoryLayerErr {

	var aux = func(oldValue, newValue, column any) error {
		if newValue != "" && oldValue != newValue {
			stmt := fmt.Sprintf("UPDATE users SET %s = $1 WHERE id = $2;", column)
			if _, err := ur.UserStore.Db.Exec(stmt, newValue, u.ID); err != nil {
				return err
			}
		}
		return nil
	}

	if err := aux(u.Username, newUserData.Username, "username"); err != nil {
		return &RepositoryLayerErr{err, "Insert Error"}
	}

	if err := aux(u.Email, newUserData.Email, "email"); err != nil {
		return &RepositoryLayerErr{err, "Insert Error"}
	}

	if err := aux(u.FirstName, newUserData.FirstName, "first_name"); err != nil {
		return &RepositoryLayerErr{err, "Insert Error"}
	}

	if err := aux(u.LastName, newUserData.LastName, "last_name"); err != nil {
		return &RepositoryLayerErr{err, "Insert Error"}
	}

	if err := aux(u.Address, newUserData.Address, "address"); err != nil {
		return &RepositoryLayerErr{err, "Insert Error"}
	}

	if err := aux(u.Address2, newUserData.Address2, "address2"); err != nil {
		return &RepositoryLayerErr{err, "Insert Error"}
	}

	if err := aux(u.City, newUserData.City, "city"); err != nil {
		return &RepositoryLayerErr{err, "Insert Error"}
	}

	if err := aux(u.State, newUserData.State, "state"); err != nil {
		return &RepositoryLayerErr{err, "Insert Error"}
	}

	if err := aux(u.Zipcode, newUserData.Zipcode, "zipcode"); err != nil {
		return &RepositoryLayerErr{err, "Insert Error"}
	}

	if err := aux(u.Telephone, newUserData.Telephone, "telephone"); err != nil {
		return &RepositoryLayerErr{err, "Insert Error"}
	}

	if err := aux(u.StripeID, newUserData.StripeID, "stripe_id"); err != nil {
		return &RepositoryLayerErr{err, "Insert Error"}
	}

	if err := aux(u.SubscriptionStatus, newUserData.SubscriptionStatus, "subscription_status"); err != nil {
		return &RepositoryLayerErr{err, "Insert Error"}
	}

	return nil
}

func (ur *UserRepository) CheckEmailExists(email string) bool {

	stmt := "SELECT email FROM users WHERE email = $1"

	queryResult := ur.UserStore.Db.QueryRow(stmt, email).Scan()

	return queryResult != sql.ErrNoRows
}

func (ur *UserRepository) CheckUsernameExists(username string) bool {

	stmt := "SELECT username FROM users WHERE username = $1"

	queryResult := ur.UserStore.Db.QueryRow(stmt, username).Scan()

	return queryResult != sql.ErrNoRows
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

func (ur *UserRepository) GetUsersByAny(any string) ([]User, error) {

	var users []User

	stmt := `SELECT id, username, first_name, last_name, email, occupation_area
		FROM users
		WHERE LOWER(username)
		LIKE CONCAT(LOWER($1::text),'%%')
		OR LOWER(CONCAT(first_name,' ',last_name))
		LIKE CONCAT('%%',LOWER($1::text),'%%');`

	rows, err := ur.UserStore.Db.Query(stmt, any)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var usr User
		if err = rows.Scan(
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

func (ur *UserRepository) GetUsersByTagID(tagId string) ([]User, error) {

	var users []User

	stmt := `SELECT id, username, first_name, last_name, email, occupation_area
		FROM users
		WHERE id
		IN (
			SELECT fk_user_id
			FROM users_tags
			WHERE fk_tag_id = $1
		);`

	rows, err := ur.UserStore.Db.Query(stmt, tagId)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var usr User
		if err = rows.Scan(
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
	var trashPassword string

	stmt := "SELECT * FROM users WHERE username = $1"

	if err := ur.UserStore.Db.QueryRow(stmt, username).Scan(
		&usr.ID,
		&usr.Username,
		&usr.FirstName,
		&usr.LastName,
		&usr.Email,
		&usr.OccupationArea,
		&usr.Telephone,
		&usr.ReferFriend,
		&trashPassword,
		&usr.Role,
		&usr.RegistrationStatus,
		&usr.CreatedAt,
		&usr.UpdatedAt,
		&usr.Address,
		&usr.Address2,
		&usr.City,
		&usr.State,
		&usr.Zipcode,
		&usr.StripeID,
		&usr.SubscriptionStatus,
	); err != nil {
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

func (ur *UserRepository) GetUser(id string) (User, *RepositoryLayerErr) {

	var usr User
	stmt := "SELECT * FROM users WHERE id::text = $1 OR email = $1 OR username = $1;"

	if err := ur.UserStore.Db.QueryRow(stmt, id).Scan(
		&usr.ID,
		&usr.Username,
		&usr.Password,
		&usr.FirstName,
		&usr.LastName,
		&usr.Email,
		&usr.OccupationArea,
		&usr.Telephone,
		&usr.ReferFriend,
		&usr.Role,
		&usr.RegistrationStatus,
		&usr.ProfilePic,
		&usr.ProfileDescription,
		&usr.CompanyLogo,
		&usr.CompanyName,
		&usr.MainProduct,
		&usr.PresentationVideo,
		&usr.Resume,
		&usr.Address,
		&usr.Address2,
		&usr.City,
		&usr.State,
		&usr.Zipcode,
		&usr.StripeID,
		&usr.SubscriptionStatus,
		&usr.SubscriptionExpiresAt,
		&usr.CreatedAt,
		&usr.UpdatedAt,
	); err != nil {
		return User{}, &RepositoryLayerErr{sql.ErrNoRows, "Usuario inexistente."}
	}

	return usr, nil
}

func (ur *UserRepository) GetUserByID(id string) (User, error) {

	var usr User
	var trashPassword string

	stmt := `SELECT * FROM users WHERE id = $1`

	if err := ur.UserStore.Db.QueryRow(stmt, id).Scan(
		&usr.ID,
		&usr.Username,
		&usr.FirstName,
		&usr.LastName,
		&usr.Email,
		&usr.OccupationArea,
		&usr.Telephone,
		&usr.ReferFriend,
		&trashPassword,
		&usr.Role,
		&usr.RegistrationStatus,
		&usr.CreatedAt,
		&usr.UpdatedAt,
		&usr.Address,
		&usr.Address2,
		&usr.City,
		&usr.State,
		&usr.Zipcode,
		&usr.StripeID,
		&usr.SubscriptionStatus,
	); err != nil {
		return usr, err
	}

	return usr, nil
}

func (ur *UserRepository) GetUserByStripeID(id string) (User, error) {

	var usr User
	var trashPassword string

	stmt := `SELECT * FROM users WHERE stripe_id = $1`

	if err := ur.UserStore.Db.QueryRow(stmt, id).Scan(
		&usr.ID,
		&usr.Username,
		&usr.FirstName,
		&usr.LastName,
		&usr.Email,
		&usr.OccupationArea,
		&usr.Telephone,
		&usr.ReferFriend,
		&trashPassword,
		&usr.Role,
		&usr.RegistrationStatus,
		&usr.CreatedAt,
		&usr.UpdatedAt,
		&usr.Address,
		&usr.Address2,
		&usr.City,
		&usr.State,
		&usr.Zipcode,
		&usr.StripeID,
		&usr.SubscriptionStatus,
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

func (ur *UserRepository) DeleteUserTag(u User, tagId string) *RepositoryLayerErr {

	stmt := `DELETE FROM users_tags WHERE fk_user_id = $1 AND fk_tag_id = $2;`

	_, err := ur.UserStore.Db.Exec(
		stmt,
		u.ID,
		tagId,
	)

	if err != nil {
		return &RepositoryLayerErr{err, "Insert Error"}
	}

	return nil
}

func (ur *UserRepository) GetUserTags(user User) ([]Tag, *RepositoryLayerErr) {

	var tagsIdList []string
	var tags []Tag

	stmt := "SELECT fk_tag_id FROM users_tags WHERE fk_user_id = $1"

	rows, err := ur.UserStore.Db.Query(stmt, user.ID)

	if err != nil {
		return nil, &RepositoryLayerErr{err, "Insert Error"}
	}

	for rows.Next() {
		var tagId string
		if err := rows.Scan(
			&tagId,
		); err != nil {
			return nil, &RepositoryLayerErr{err, "Insert Error"}
		}
		tagsIdList = append(tagsIdList, tagId)
	}

	stmt = "SELECT * FROM tags WHERE tag_id = $1"

	for _, id := range tagsIdList {
		var tag Tag
		if err := ur.UserStore.Db.QueryRow(stmt, id).Scan(
			&tag.ID,
			&tag.Name,
		); err != nil {
			return nil, &RepositoryLayerErr{err, "Insert Error"}
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

func (ur *UserRepository) GetUserNotTags(user User) ([]Tag, *RepositoryLayerErr) {

	var tagsIdList []string
	var tags []Tag

	stmt := `SELECT t.tag_id, t.tag_name, ut.fk_user_id
		FROM tags t
		INNER JOIN users_tags ut
		ON t.tag_id = ut.fk_tag_id
		WHERE ut.fk_user_id != '55a399f3-a7d3-4c9b-bdf1-300ca6bbcc50'
		AND tag_name LIKE CONCAT('%%', 'pro', '%%');`

	rows, err := ur.UserStore.Db.Query(stmt, user.ID)

	if err != nil {
		return nil, &RepositoryLayerErr{err, "Insert Error"}
	}

	for rows.Next() {
		var tagId string
		if err := rows.Scan(
			&tagId,
		); err != nil {
			return nil, &RepositoryLayerErr{err, "Insert Error"}
		}
		tagsIdList = append(tagsIdList, tagId)
	}

	stmt = "SELECT * FROM tags WHERE tag_id != $1"

	for _, id := range tagsIdList {
		var tag Tag
		if err := ur.UserStore.Db.QueryRow(stmt, id).Scan(
			&tag.ID,
			&tag.Name,
		); err != nil {
			return nil, &RepositoryLayerErr{err, "Insert Error"}
		}
		tags = append(tags, tag)
	}

	return tags, nil
}
