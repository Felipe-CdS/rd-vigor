package services

import (
	"net/http"

	"nugu.dev/rd-vigor/repositories"
)

type TagRepository interface {
	GetAllTags() ([]repositories.Tag, *repositories.RepositoryLayerErr)
	CreateTag(t repositories.Tag) *repositories.RepositoryLayerErr
	CheckTagExists(name string) bool
}

type TagService struct {
	Repository TagRepository
}

func NewTagService(tr TagRepository) *TagService {
	return &TagService{
		Repository: tr,
	}
}

func (ts *TagService) CreateTag(n string) *ServiceLayerErr {

	newTag := repositories.Tag{
		Name: n,
	}

	if ts.Repository.CheckTagExists(n) {
		return &ServiceLayerErr{nil, "Tag já cadastrada.", http.StatusBadRequest}
	}

	err := ts.Repository.CreateTag(newTag)

	if err != nil {
		return &ServiceLayerErr{err.Error, "Query Err", http.StatusInternalServerError}
	}

	return nil
}

func (ts *TagService) GetAllTags() ([]repositories.Tag, *ServiceLayerErr) {

	tags, err := ts.Repository.GetAllTags()

	if err != nil {
		return nil, &ServiceLayerErr{err.Error, "Query Err", http.StatusInternalServerError}
	}

	return tags, nil
}
