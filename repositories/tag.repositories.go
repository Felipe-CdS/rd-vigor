package repositories

import (
	"database/sql"
	"strings"

	"github.com/google/uuid"
	"nugu.dev/rd-vigor/db"
)

type Tag struct {
	ID         string `JSON:"tag_id"`
	Name       string `JSON:"name"`
	CategoryID string `JSON:"fk_category_id"`
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

func (tr *TagRepository) CreateTag(t Tag, tc TagCategory) *RepositoryLayerErr {

	stmt := `INSERT INTO tags (tag_id, tag_name, fk_category_id) VALUES ($1, $2, $3)`

	_, err := tr.TagStore.Db.Exec(
		stmt,
		uuid.New(),
		strings.ToLower(t.Name),
		tc.ID,
	)

	if err != nil {
		return &RepositoryLayerErr{err, "Insert Error"}
	}

	return nil
}

func (tr *TagRepository) CheckTagExists(name string) bool {

	stmt := "SELECT tag_name FROM tags WHERE tag_name = $1"

	queryResult := tr.TagStore.Db.QueryRow(stmt, strings.ToLower(name)).Scan()

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

	if name == "" {
		return tags, nil
	}

	stmt := "SELECT * FROM tags WHERE tag_name LIKE CONCAT('%%',$1::text,'%%')"

	rows, err := tr.TagStore.Db.Query(stmt, strings.ToLower(name))

	if err != nil {
		return nil, &RepositoryLayerErr{err, "Insert Error"}
	}

	for rows.Next() {
		var tag Tag
		if err := rows.Scan(
			&tag.ID,
			&tag.Name,
			&tag.CategoryID,
		); err != nil {
			return nil, &RepositoryLayerErr{err, "Insert Error"}
		}
		tags = append(tags, tag)
	}

	return tags, nil
}

func (tr *TagRepository) SearchTagByNameAvaiableToUser(user User, name string) ([]Tag, *RepositoryLayerErr) {

	var tags []Tag

	if name == "" {
		return tags, nil
	}

	stmt := `SELECT t.tag_id, t.tag_name
		FROM tags t
		LEFT JOIN users_tags ut
		ON t.tag_id = ut.fk_tag_id
		WHERE ut.fk_user_id != $1
		AND t.tag_name LIKE CONCAT('%%',$2::text,'%%');`

	rows, err := tr.TagStore.Db.Query(stmt, user.ID, strings.ToLower(name))

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

func (tr *TagRepository) GetTagById(id string) (Tag, *RepositoryLayerErr) {

	var t Tag
	stmt := "SELECT * FROM tags WHERE tag_id = $1;"

	if err := tr.TagStore.Db.QueryRow(stmt, id).Scan(
		&t.ID,
		&t.Name,
		&t.CategoryID,
	); err != nil {
		return t, &RepositoryLayerErr{err, "Insert Error"}
	}

	return t, nil
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
		if err = rows.Scan(
			&s,
		); err != nil {
			return nil, &RepositoryLayerErr{err, "Insert Error"}
		}
		relationsIDs = append(relationsIDs, s)
	}

	stmt = "SELECT * FROM tags WHERE tag_id = $1"

	for _, i := range relationsIDs {

		rows, err = tr.TagStore.Db.Query(stmt, i)

		if err != nil {
			return nil, &RepositoryLayerErr{err, "Search Error"}
		}

		for rows.Next() {
			var t Tag
			if err := rows.Scan(
				&t.ID,
				&t.Name,
				&t.CategoryID,
			); err != nil {
				return nil, &RepositoryLayerErr{err, "Search Error"}
			}
			tags = append(tags, t)
		}
	}

	return tags, nil
}
