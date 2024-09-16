package repositories

import (
	"github.com/google/uuid"
	"nugu.dev/rd-vigor/db"
)

type Chatroom struct {
	ChatroomId string `JSON:"chatroom_id"`
}

type ChatroomRepository struct {
	Chatroom      Chatroom
	ChatroomStore db.Store
}

func NewChatroomRepository(c Chatroom, cStore db.Store) *ChatroomRepository {

	return &ChatroomRepository{
		Chatroom:      c,
		ChatroomStore: cStore,
	}
}

func (cr *ChatroomRepository) CreateChatroom(sender_id string, recipient_id string) (string, *RepositoryLayerErr) {

	tx, err := cr.ChatroomStore.Db.Begin()
	chatroomId := uuid.New()

	if err != nil {
		return "", &RepositoryLayerErr{err, "Error Creating transaction."}
	}

	_, err = tx.Exec(
		`INSERT INTO chatrooms (chatroom_id) VALUES ($1)`,
		chatroomId,
	)

	if err != nil {
		tx.Rollback()
		return "", &RepositoryLayerErr{err, "Error Creating transaction."}
	}

	_, err = tx.Exec(
		`INSERT INTO chatrooms_users (fk_user_id, fk_chatroom_id) VALUES ($1, $2);`,
		sender_id,
		chatroomId,
	)

	if err != nil {
		tx.Rollback()
		return "", &RepositoryLayerErr{err, "Error Creating transaction."}
	}

	_, err = tx.Exec(
		`INSERT INTO chatrooms_users (fk_user_id, fk_chatroom_id) VALUES ($1, $2);`,
		recipient_id,
		chatroomId,
	)

	if err != nil {
		tx.Rollback()
		return "", &RepositoryLayerErr{err, "Error Creating transaction."}
	}

	err = tx.Commit()

	if err != nil {
		return "", &RepositoryLayerErr{err, "Error Creating transaction."}
	}

	return chatroomId.String(), nil
}

func (cr *ChatroomRepository) GetChatroomsByUserID(id string) ([]Chatroom, *RepositoryLayerErr) {

	var chatrooms []Chatroom

	stmt := "SELECT fk_chatroom_id FROM chatrooms_users WHERE fk_user_id = $1"

	rows, err := cr.ChatroomStore.Db.Query(stmt, id)

	if err != nil {
		return nil, &RepositoryLayerErr{err, "Search Error"}
	}

	for rows.Next() {
		var c Chatroom
		if err := rows.Scan(
			&c.ChatroomId,
		); err != nil {
			return nil, &RepositoryLayerErr{err, "Insert Error"}
		}
		chatrooms = append(chatrooms, c)
	}

	return chatrooms, nil
}

func (cr *ChatroomRepository) GetChatroomRecipient(user1_id string, chatroom_id string) (string, *RepositoryLayerErr) {

	var recipient string

	stmt := "SELECT fk_user_id FROM chatrooms_users WHERE fk_chatroom_id = $1"

	rows, err := cr.ChatroomStore.Db.Query(stmt, chatroom_id)

	if err != nil {
		return recipient, &RepositoryLayerErr{err, "Search Error"}
	}

	for rows.Next() {
		var u User
		if err := rows.Scan(
			&u.ID,
		); err != nil {
			return recipient, &RepositoryLayerErr{err, "Insert Error"}
		}

		if u.ID != user1_id {
			recipient = u.ID
			return recipient, nil
		}
	}

	return recipient, nil
}
