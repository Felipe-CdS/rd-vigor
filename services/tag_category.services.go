package services

import (
	"fmt"
	"net/http"

	"nugu.dev/rd-vigor/repositories"
)

type TagCategoryRepository interface {
	CreateTagCategory(t repositories.TagCategory) *repositories.RepositoryLayerErr

	GetAllTagCategories() ([]repositories.TagCategory, *repositories.RepositoryLayerErr)
	GetAllTagsByCategory(categoryId string) ([]repositories.Tag, *repositories.RepositoryLayerErr)
}

type TagCategoryService struct {
	Repository TagCategoryRepository
}

func NewTagCategoryService(tr TagCategoryRepository) *TagCategoryService {
	return &TagCategoryService{
		Repository: tr,
	}
}

func (ts *TagCategoryService) CreateTagCategory(n string) *ServiceLayerErr {

	c := repositories.TagCategory{
		Name: n,
	}

	err := ts.Repository.CreateTagCategory(c)

	if err != nil {
		fmt.Println(err)
		return &ServiceLayerErr{err.Error, "Query Err", http.StatusInternalServerError}
	}

	return nil
}

func (ts *TagCategoryService) GetAllTagCategories() ([]repositories.TagCategory, *ServiceLayerErr) {

	c, err := ts.Repository.GetAllTagCategories()

	if err != nil {
		return nil, &ServiceLayerErr{err.Error, "Query Err", http.StatusInternalServerError}
	}

	return c, nil
}

func (ts *TagCategoryService) GetAllTagsByCategory(category_id string) ([]repositories.Tag, *ServiceLayerErr) {

	c, err := ts.Repository.GetAllTagsByCategory(category_id)

	if err != nil {
		return nil, &ServiceLayerErr{err.Error, "Query Err", http.StatusInternalServerError}
	}

	return c, nil
}
