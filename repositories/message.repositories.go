package repositories

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"nugu.dev/rd-vigor/db"
)

type Message struct {
	MessageId  string `JSON:"message_id"`
	chatroomId string `JSON:"fk_chatroom_id"`

	UserId     string    `JSON:"fk_user_id"`
	Content    string    `JSON:"content"`
	Created_at time.Time `JSON:"created_at"`
}

type MessageRepository struct {
	Message      Message
	MessageStore db.Store
}

func NewMessageRepository(m Message, mStore db.Store) *MessageRepository {

	return &MessageRepository{
		Message:      m,
		MessageStore: mStore,
	}
}

func (mr *MessageRepository) GetChatroomMessages(chatroomId string) ([]Message, *RepositoryLayerErr) {

	var messages []Message

	stmt := "SELECT message_id, fk_sender_id, content, created_at FROM messages WHERE fk_chatroom_id = $1"

	rows, err := mr.MessageStore.Db.Query(stmt, chatroomId)

	if err != nil {
		return nil, &RepositoryLayerErr{err, "Search Error"}
	}

	for rows.Next() {
		var msg Message
		if err := rows.Scan(
			&msg.MessageId,
			&msg.UserId,
			&msg.Content,
			&msg.Created_at,
		); err != nil {
			return nil, &RepositoryLayerErr{err, "Insert Error"}
		}
		messages = append(messages, msg)
	}

	return messages, nil
}

func (mr *MessageRepository) GetLastChatroomMessage(chatroom_id string) (Message, *RepositoryLayerErr) {

	var msg Message

	stmt := "SELECT message_id, fk_sender_id, content, created_at FROM messages WHERE fk_chatroom_id = $1 ORDER BY created_at DESC LIMIT 1"

	err := mr.MessageStore.Db.QueryRow(stmt, chatroom_id).Scan(
		&msg.MessageId,
		&msg.UserId,
		&msg.Content,
		&msg.Created_at)

	if err != nil {
		return msg, &RepositoryLayerErr{Error: err}
	}

	return msg, nil
}

func (mr *MessageRepository) CreateMessage(sender_id string, content string, chatroom_id string) *RepositoryLayerErr {

	fmt.Println(sender_id, content, chatroom_id)
	stmt := `INSERT INTO messages 
		(message_id, fk_sender_id, fk_chatroom_id, content, created_at) 
		VALUES ($1, $2, $3, $4, $5)`

	_, err := mr.MessageStore.Db.Exec(
		stmt,
		uuid.New(),
		sender_id,
		chatroom_id,
		content,
		time.Now(),
	)

	if err != nil {
		return &RepositoryLayerErr{err, "Insert Error"}
	}

	return nil
}
