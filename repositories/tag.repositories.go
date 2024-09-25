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

	return queryResult != sql.ErrNoRows
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

func (tr *TagRepository) GetUserTags(u User) ([]Tag, *RepositoryLayerErr) {

	var tags []Tag
	var relationsIDs []string

	stmt := "SELECT fk_tag_id FROM users_tags WHERE fk_user_id = $1"

	rows, err := tr.TagStore.Db.Query(stmt, u.ID)

	if err != nil {
		return nil, &RepositoryLayerErr{err, "Insert Error"}
	}

	for rows.Next() {
		var s string
		if err := rows.Scan(
			&s,
		); err != nil {
			return nil, &RepositoryLayerErr{err, "Insert Error"}
		}
		relationsIDs = append(relationsIDs, s)
	}

	stmt = "SELECT * FROM tags WHERE tag_id = $1"

	for i := range relationsIDs {

		rows, err = tr.TagStore.Db.Query(stmt, i)

		if err != nil {
			return nil, &RepositoryLayerErr{err, "Search Error"}
		}

		for rows.Next() {
			var t Tag
			if err := rows.Scan(
				&t.ID,
				&t.Name,
			); err != nil {
				return nil, &RepositoryLayerErr{err, "Search Error"}
			}
			tags = append(tags, t)
		}
	}

	return tags, nil
}
