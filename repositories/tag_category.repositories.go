package repositories

import (
	"database/sql"
	"strings"

	"github.com/google/uuid"
	"nugu.dev/rd-vigor/db"
)

type TagCategory struct {
	ID   string `JSON:"category_id"`
	Name string `JSON:"category_name"`
}

type TagCategoryRepository struct {
	TagCategory      TagCategory
	TagCategoryStore db.Store
}

func NewTagCategoryRepository(t TagCategory, tStore db.Store) *TagCategoryRepository {

	return &TagCategoryRepository{
		TagCategory:      t,
		TagCategoryStore: tStore,
	}
}

func (tr *TagCategoryRepository) CreateTagCategory(t TagCategory) *RepositoryLayerErr {

	stmt := "SELECT category_name FROM tag_categories WHERE category_name = $1"

	queryResult := tr.TagCategoryStore.Db.QueryRow(stmt, strings.ToLower(t.Name)).Scan()

	if queryResult != sql.ErrNoRows {
		return &RepositoryLayerErr{nil, "Category already exists"}
	}

	stmt = `INSERT INTO tag_categories (category_id, category_name) VALUES ($1, $2)`

	_, err := tr.TagCategoryStore.Db.Exec(
		stmt,
		uuid.New(),
		strings.ToLower(t.Name),
	)

	if err != nil {
		return &RepositoryLayerErr{err, "Insert Error"}
	}

	return nil
}

func (tr *TagCategoryRepository) GetAllTagCategories() ([]TagCategory, *RepositoryLayerErr) {

	var tags []TagCategory

	stmt := "SELECT * FROM tag_categories"

	rows, err := tr.TagCategoryStore.Db.Query(stmt)

	if err != nil {
		return nil, &RepositoryLayerErr{err, "Insert Error"}
	}

	for rows.Next() {
		var t TagCategory
		if err := rows.Scan(
			&t.ID,
			&t.Name,
		); err != nil {
			return nil, &RepositoryLayerErr{err, "Insert Error"}
		}
		tags = append(tags, t)
	}

	return tags, nil
}

func (tr *TagCategoryRepository) GetAllTagsByCategory(category_id string) ([]Tag, *RepositoryLayerErr) {

	var tags []Tag

	stmt := `SELECT tags.tag_id, tags.tag_name 
			FROM tags
			INNER JOIN tag_categories
			ON tags.fk_category_id = tag_categories.category_id
			WHERE tag_categories.category_id = $1;`

	rows, err := tr.TagCategoryStore.Db.Query(stmt, category_id)

	if err != nil {
		return nil, &RepositoryLayerErr{err, "Insert Error"}
	}

	for rows.Next() {
		var t Tag
		if err := rows.Scan(
			&t.ID,
			&t.Name,
		); err != nil {
			return nil, &RepositoryLayerErr{err, "Insert Error"}
		}
		tags = append(tags, t)
	}

	return tags, nil
}
