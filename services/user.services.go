package services

import (
	"bytes"
	"database/sql"
	"fmt"
	"io"
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
	CheckEmailExists(email string) bool
	CheckUsernameExists(username string) bool
	GetAllUsers() ([]repositories.User, error)
	GetUserByID(id string) (repositories.User, error)
	GetUserByEmail(email string) (repositories.User, *repositories.RepositoryLayerErr)
	GetUserByUsername(username string) (repositories.User, *repositories.RepositoryLayerErr)
	GetUserPasswordByID(id string) (string, *repositories.RepositoryLayerErr)
}

type UserService struct {
	Repository UserRepository
}

func NewUserService(ur UserRepository) *UserService {
	return &UserService{
		Repository: ur,
	}
}

func (us *UserService) CreateUser(u repositories.User) *ServiceLayerErr {

	if _, err := mail.ParseAddress(u.Email); err != nil {
		return &ServiceLayerErr{nil, "E-mail inválido.", http.StatusBadRequest}
	}

	if us.Repository.CheckEmailExists(u.Email) {
		return &ServiceLayerErr{nil, "Email já cadastrado.", http.StatusBadRequest}
	}

	if us.Repository.CheckUsernameExists(u.Username) {
		return &ServiceLayerErr{nil, "Nome de usuário já cadastrado.", http.StatusBadRequest}
	}

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

func (us *UserService) AuthUser(login string, password string) (repositories.User, *ServiceLayerErr) {

	var user repositories.User
	var queryErr *repositories.RepositoryLayerErr
	var isEmail bool = true

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
