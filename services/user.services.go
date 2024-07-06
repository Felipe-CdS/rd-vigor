package services

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/mail"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"nugu.dev/rd-vigor/repositories"
)

type UserRepository interface {
	CreateUser(u repositories.User) *repositories.RepositoryLayerErr
	CheckEmailExists(email string) bool
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

	us.Repository.CreateUser(u)

	return nil
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
