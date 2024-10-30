package services

import (
	"bytes"
	"database/sql"
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/mail"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"golang.org/x/crypto/bcrypt"
	"nugu.dev/rd-vigor/repositories"
)

type UserRepository interface {
	CreateUser(u repositories.User) *repositories.RepositoryLayerErr
	UpdateUser(u repositories.User, newUserData repositories.User) *repositories.RepositoryLayerErr
	CheckEmailExists(email string) bool
	CheckUsernameExists(username string) bool

	GetAllUsers() ([]repositories.User, error)
	GetUsersByAny(any string) ([]repositories.User, error)
	GetUsersByTagID(tagId string) ([]repositories.User, error)

	GetUserByID(id string) (repositories.User, error)
	GetUserByStripeID(id string) (repositories.User, error)
	GetUserByEmail(email string) (repositories.User, *repositories.RepositoryLayerErr)
	GetUserByUsername(username string) (repositories.User, *repositories.RepositoryLayerErr)
	GetUserPasswordByID(id string) (string, *repositories.RepositoryLayerErr)

	GetUserTags(user repositories.User) ([]repositories.Tag, *repositories.RepositoryLayerErr)
	GetUserNotTags(user repositories.User) ([]repositories.Tag, *repositories.RepositoryLayerErr)
	SetNewTagUser(u repositories.User, t repositories.Tag) *repositories.RepositoryLayerErr
	DeleteUserTag(u repositories.User, tagId string) *repositories.RepositoryLayerErr
}

type UserService struct {
	Repository    UserRepository
	TagRepository TagRepository
}

func NewUserService(ur UserRepository, tr TagRepository) *UserService {
	return &UserService{
		Repository:    ur,
		TagRepository: tr,
	}
}

func (us *UserService) CreateUser(data map[string]string) *ServiceLayerErr {

	var generatedUsername string

	for {
		generatedUsername = fmt.Sprintf("%s-%s-%05d", data["FirstName"], data["LastName"], rand.Intn(10000))

		if !us.Repository.CheckUsernameExists(generatedUsername) {
			break
		}
	}

	u := repositories.User{
		Username:       generatedUsername,
		FirstName:      data["FirstName"],
		LastName:       data["LastName"],
		Email:          data["Email"],
		OccupationArea: data["OccupationArea"],
		Telephone:      data["Telephone"],
		Password:       data["Password"],
		ReferFriend:    data["ReferFriend"],
	}

	if _, err := mail.ParseAddress(u.Email); err != nil {
		return &ServiceLayerErr{nil, "E-mail inválido.", http.StatusBadRequest}
	}

	if us.Repository.CheckEmailExists(u.Email) {
		return &ServiceLayerErr{nil, "Email já cadastrado.", http.StatusBadRequest}
	}

	if us.Repository.CheckUsernameExists(u.Username) {
		return &ServiceLayerErr{nil, "Nome de usuário já cadastrado.", http.StatusBadRequest}
	}

	// CHECK ERROR
	us.Repository.CreateUser(u)

	return nil
}

func (us *UserService) GetAllUsers() ([]repositories.User, *ServiceLayerErr) {

	users, err := us.Repository.GetAllUsers()

	if err != nil {
		return nil, &ServiceLayerErr{err, "Query Err", http.StatusInternalServerError}
	}
	return users, nil
}

func (us *UserService) GetUserByID(id string) (repositories.User, *ServiceLayerErr) {

	users, err := us.Repository.GetUserByID(id)

	if err != nil {
		return repositories.User{}, &ServiceLayerErr{err, "Query Err", http.StatusInternalServerError}
	}
	return users, nil
}

func (us *UserService) GetUserByStripeID(id string) (repositories.User, *ServiceLayerErr) {

	users, err := us.Repository.GetUserByStripeID(id)

	if err != nil {
		return repositories.User{}, &ServiceLayerErr{err, "Query Err", http.StatusInternalServerError}
	}
	return users, nil
}

func (us *UserService) GetUserByUsername(username string) (repositories.User, *ServiceLayerErr) {

	usr, err := us.Repository.GetUserByUsername(username)

	if err != nil {
		return repositories.User{}, &ServiceLayerErr{err.Error, "Query Err", http.StatusInternalServerError}
	}
	return usr, nil
}

func (us *UserService) GetUsersByAny(any string) ([]repositories.User, *ServiceLayerErr) {

	found, err := us.Repository.GetUsersByAny(any)

	if err == nil {
		return found, nil
	}

	return []repositories.User{}, &ServiceLayerErr{err, "Query Err", http.StatusInternalServerError}
}

func (us *UserService) GetUsersByTagID(tagId string) ([]repositories.User, *ServiceLayerErr) {

	found, err := us.Repository.GetUsersByTagID(tagId)

	if err == nil {
		return found, nil
	}

	return []repositories.User{}, &ServiceLayerErr{err, "Query Err", http.StatusInternalServerError}
}

func (us *UserService) AuthUser(login string, password string) (repositories.User, *ServiceLayerErr) {

	var user repositories.User
	var queryErr *repositories.RepositoryLayerErr
	var isEmail = true

	if login == "" || password == "" {
		return repositories.User{}, &ServiceLayerErr{nil, "Por favor, preencha ambos os campos.", http.StatusBadRequest}
	}

	if _, err := mail.ParseAddress(login); err != nil {
		isEmail = false
	}

	if isEmail {
		user, queryErr = us.Repository.GetUserByEmail(login)
	} else {
		user, queryErr = us.Repository.GetUserByUsername(login)
	}

	if queryErr != nil {
		if queryErr.Error == sql.ErrNoRows {
			return repositories.User{}, &ServiceLayerErr{queryErr.Error, "Login ou senha incorretos.", http.StatusBadRequest}
		}
		return repositories.User{}, &ServiceLayerErr{queryErr.Error, "Query Err", http.StatusInternalServerError}
	}

	userPassword, queryErr := us.Repository.GetUserPasswordByID(user.ID)

	if queryErr != nil {
		return repositories.User{}, &ServiceLayerErr{queryErr.Error, "Query Err", http.StatusInternalServerError}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(password)); err != nil {
		return repositories.User{}, &ServiceLayerErr{err, "Login ou senha incorretos.", http.StatusBadRequest}
	}

	return user, nil
}

func putFileS3(f *multipart.FileHeader) *ServiceLayerErr {

	fileContent, err := f.Open()

	if err != nil {
		return &ServiceLayerErr{err, "file", http.StatusInternalServerError}
	}

	defer fileContent.Close()

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2"),
	})

	if err != nil {
		return &ServiceLayerErr{err, "New AWS Session", http.StatusInternalServerError}
	}

	byteContainer, err := io.ReadAll(fileContent)

	if err != nil {
		return &ServiceLayerErr{err, "Read File Content", http.StatusInternalServerError}
	}

	svc := s3.New(sess)

	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String("rd-vigor-20050411"),
		Key:    aws.String(fmt.Sprintf("%s", f.Filename)),
		Body:   bytes.NewReader(byteContainer),
	})

	if err != nil {
		return &ServiceLayerErr{err, "Put Object S3", http.StatusInternalServerError}
	}

	return nil
}

func (us *UserService) SetNewTagUser(username string, tagId string) *ServiceLayerErr {

	user, err := us.Repository.GetUserByUsername(username)

	if err != nil {
		return &ServiceLayerErr{err.Error, "Query Err 1", http.StatusInternalServerError}
	}

	tag, err := us.TagRepository.GetTagById(tagId)

	if err != nil {
		return &ServiceLayerErr{err.Error, "Query Err 2", http.StatusInternalServerError}
	}

	//CHECK IF USER ALREADY HAS TAG

	queryErr := us.Repository.SetNewTagUser(user, tag)

	if queryErr != nil {
		return &ServiceLayerErr{queryErr.Error, "Query Err 3", http.StatusInternalServerError}
	}

	return nil
}

func (us *UserService) DeleteUserTag(user repositories.User, tagId string) *ServiceLayerErr {

	queryErr := us.Repository.DeleteUserTag(user, tagId)

	if queryErr != nil {
		return &ServiceLayerErr{queryErr.Error, "Delete Err", http.StatusInternalServerError}
	}

	return nil
}

func (us *UserService) GetUserTags(user repositories.User) ([]repositories.Tag, *ServiceLayerErr) {

	tags, err := us.Repository.GetUserTags(user)

	if err != nil {
		return nil, &ServiceLayerErr{err.Error, "Query Err 3", http.StatusInternalServerError}
	}
	return tags, nil
}

func (us *UserService) GetUserNotTags(user repositories.User) ([]repositories.Tag, *ServiceLayerErr) {

	tags, err := us.Repository.GetUserNotTags(user)

	if err != nil {
		return nil, &ServiceLayerErr{err.Error, "Query Err 3", http.StatusInternalServerError}
	}
	return tags, nil
}

// Search Tag by name but only if user doesnt already have it
func (us *UserService) SearchTagByNameAvaiableToUser(user repositories.User, query string) ([]repositories.Tag, *ServiceLayerErr) {

	tags, err := us.TagRepository.SearchTagByNameAvaiableToUser(user, query)

	if err != nil {
		return nil, &ServiceLayerErr{err.Error, "Query Err", http.StatusInternalServerError}
	}
	return tags, nil
}

func (us *UserService) UpdateUser(u repositories.User, newUserData repositories.User) *ServiceLayerErr {

	if _, err := mail.ParseAddress(newUserData.Email); err != nil {
		return &ServiceLayerErr{nil, "E-mail inválido.", http.StatusBadRequest}
	}

	if u.Email != newUserData.Email && us.Repository.CheckEmailExists(newUserData.Email) {
		return &ServiceLayerErr{nil, "Email já cadastrado.", http.StatusBadRequest}
	}

	if u.Username != newUserData.Username && us.Repository.CheckUsernameExists(newUserData.Username) {
		return &ServiceLayerErr{nil, "Nome de usuário já cadastrado.", http.StatusBadRequest}
	}

	if err := us.Repository.UpdateUser(u, newUserData); err != nil {
		return &ServiceLayerErr{nil, "Update Error.", http.StatusBadRequest}
	}

	return nil
}
