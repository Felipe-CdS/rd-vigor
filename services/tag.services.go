package services

import (
	"net/http"

	"nugu.dev/rd-vigor/repositories"
)

type TagRepository interface {
	CreateTag(t repositories.Tag, tc repositories.TagCategory) *repositories.RepositoryLayerErr
	CheckTagExists(name string) bool
	GetAllTags() ([]repositories.Tag, *repositories.RepositoryLayerErr)
	SearchTagByName(name string) ([]repositories.Tag, *repositories.RepositoryLayerErr)

	GetTagById(id string) (repositories.Tag, *repositories.RepositoryLayerErr)
	GetUserTags(u repositories.User) ([]repositories.Tag, *repositories.RepositoryLayerErr)

	SearchTagByNameAvaiableToUser(u repositories.User, name string) ([]repositories.Tag, *repositories.RepositoryLayerErr)
}

type TagService struct {
	Repository            TagRepository
	TagCategoryRepository TagCategoryRepository
}

func NewTagService(tr TagRepository, tcr TagCategoryRepository) *TagService {
	return &TagService{
		Repository:            tr,
		TagCategoryRepository: tcr,
	}
}

func (ts *TagService) CreateTag(n string, tc_id string) *ServiceLayerErr {

	newTag := repositories.Tag{
		Name: n,
	}

	category := repositories.TagCategory{
		ID: tc_id,
	}

	if ts.Repository.CheckTagExists(n) {
		return &ServiceLayerErr{nil, "Tag já cadastrada.", http.StatusBadRequest}
	}

	err := ts.Repository.CreateTag(newTag, category)

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

func (ts *TagService) SearchTagByName(name string) ([]repositories.Tag, *ServiceLayerErr) {

	tags, err := ts.Repository.SearchTagByName(name)

	if err != nil {
		return nil, &ServiceLayerErr{err.Error, "Query Err", http.StatusInternalServerError}
	}

	return tags, nil
}

func (ts *TagService) GetTagByID(id string) (repositories.Tag, *ServiceLayerErr) {

	tags, err := ts.Repository.GetTagById(id)

	if err != nil {
		return tags, &ServiceLayerErr{err.Error, "Query Err", http.StatusInternalServerError}
	}

	return tags, nil
}

func (ts *TagService) GetUserTags(u repositories.User) ([]repositories.Tag, *ServiceLayerErr) {

	tags, err := ts.Repository.GetUserTags(u)

	if err != nil {
		return nil, &ServiceLayerErr{err.Error, "Query Err", http.StatusInternalServerError}
	}

	return tags, nil
}
