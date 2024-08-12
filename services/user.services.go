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
	CheckEmailExists(email string) bool
	CheckUsernameExists(username string) bool
	GetAllUsers() ([]repositories.User, error)
	GetUserByID(id string) (repositories.User, error)
	GetUserByEmail(email string) (repositories.User, *repositories.RepositoryLayerErr)
	GetUserByUsername(username string) (repositories.User, *repositories.RepositoryLayerErr)
	GetUserPasswordByID(id string) (string, *repositories.RepositoryLayerErr)
	SetNewTagUser(u repositories.User, t repositories.Tag) *repositories.RepositoryLayerErr

	GetUserTags(user repositories.User) ([]repositories.Tag, *repositories.RepositoryLayerErr)
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

func (us *UserService) GetUserByUsername(username string) (repositories.User, *ServiceLayerErr) {

	usr, err := us.Repository.GetUserByUsername(username)

	if err != nil {
		return repositories.User{}, &ServiceLayerErr{err.Error, "Query Err", http.StatusInternalServerError}
	}
	return usr, nil
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

func (us *UserService) SetNewTagUser(username string, tag_name string) *ServiceLayerErr {

	user, err := us.Repository.GetUserByUsername(username)

	if err != nil {
		return &ServiceLayerErr{err.Error, "Query Err 1", http.StatusInternalServerError}
	}

	tag, err := us.TagRepository.SearchTagByName(tag_name)

	if err != nil {
		return &ServiceLayerErr{err.Error, "Query Err 2", http.StatusInternalServerError}
	}

	//CHECK IF USER ALREADY HAS TAG

	queryErr := us.Repository.SetNewTagUser(user, tag[0])

	if queryErr != nil {
		return &ServiceLayerErr{queryErr.Error, "Query Err 3", http.StatusInternalServerError}
	}

	return nil
}

func (us *UserService) GetUserTags(user repositories.User) ([]repositories.Tag, *ServiceLayerErr) {

	tags, err := us.Repository.GetUserTags(user)

	if err != nil {

		fmt.Printf("%+v\n", err)
		return nil, &ServiceLayerErr{err.Error, "Query Err 3", http.StatusInternalServerError}
	}
	return tags, nil
}
