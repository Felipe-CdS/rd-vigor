package repositories

import (
	"database/sql"

	"github.com/google/uuid"
	"nugu.dev/rd-vigor/db"
)

type Tag struct {
	ID   string `JSON:"tag_id"`
	Name string `JSON:"name"`
}

type TagRepository struct {
	Tag      Tag
	TagStore db.Store
}

func NewTagRepository(t Tag, tStore db.Store) *TagRepository {

	return &TagRepository{
		Tag:      t,
		TagStore: tStore,
	}
}

func (tr *TagRepository) CreateTag(t Tag) *RepositoryLayerErr {

	stmt := `INSERT INTO tags (tag_id, tag_name) VALUES ($1, $2)`

	_, err := tr.TagStore.Db.Exec(
		stmt,
		uuid.New(),
		t.Name,
	)

	if err != nil {
		return &RepositoryLayerErr{err, "Insert Error"}
	}

	return nil
}

func (tr *TagRepository) CheckTagExists(name string) bool {

	stmt := "SELECT tag_name FROM tags WHERE tag_name = $1"

	queryResult := tr.TagStore.Db.QueryRow(stmt, name).Scan()

	if queryResult != sql.ErrNoRows {
		return true
	}

	return false
}

func (tr *TagRepository) GetAllTags() ([]Tag, *RepositoryLayerErr) {

	var tags []Tag

	stmt := "SELECT tag_name FROM tags"

	rows, err := tr.TagStore.Db.Query(stmt)

	if err != nil {
		return nil, &RepositoryLayerErr{err, "Insert Error"}
	}

	for rows.Next() {
		var tag Tag
		if err := rows.Scan(
			&tag.Name,
		); err != nil {
			return nil, &RepositoryLayerErr{err, "Insert Error"}
		}
		tags = append(tags, tag)
	}

	return tags, nil
}

func (tr *TagRepository) SearchTagByName(name string) ([]Tag, *RepositoryLayerErr) {

	var tags []Tag

	stmt := "SELECT * FROM tags WHERE tag_name LIKE CONCAT('%%',$1::text,'%%')"

	rows, err := tr.TagStore.Db.Query(stmt, name)

	if err != nil {
		return nil, &RepositoryLayerErr{err, "Insert Error"}
	}

	for rows.Next() {
		var tag Tag
		if err := rows.Scan(
			&tag.ID,
			&tag.Name,
		); err != nil {
			return nil, &RepositoryLayerErr{err, "Insert Error"}
		}
		tags = append(tags, tag)
	}

	return tags, nil
}
