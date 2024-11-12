CREATE TABLE users (
	id UUID PRIMARY KEY
	, username TEXT NOT NULL
	, password TEXT NOT NULL
	, first_name TEXT NOT NULL
	, last_name TEXT NOT NULL
	, email TEXT NOT NULL
	, occupation_area TEXT DEFAULT ''
	, telephone TEXT DEFAULT ''
	, refer_friend TEXT DEFAULT ''
	, role TEXT NOT NULL DEFAULT 'member'
	, registration_status TEXT NOT NULL DEFAULT 'pending'
	
	, profile_picture TEXT NOT NULL DEFAULT '/static/img/user.svg'
	, profile_description TEXT DEFAULT ''
	, company_logo TEXT DEFAULT ''
	, company_name TEXT DEFAULT ''
	, main_product TEXT DEFAULT ''
	, presentation_video TEXT DEFAULT ''
	, resume TEXT DEFAULT ''

	, address TEXT DEFAULT ''
	, address2 TEXT DEFAULT ''
	, city TEXT DEFAULT ''
	, state TEXT DEFAULT ''
	, zipcode TEXT DEFAULT ''
	
	, stripe_id TEXT DEFAULT ''
	, subsctription_status BOOLEAN DEFAULT false
	, subsctription_expires_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()

	, created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
	, updated_at TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE TABLE tag_categories (
	category_id UUID PRIMARY KEY
	, category_name TEXT NOT NULL
);

CREATE TABLE tags (
	tag_id UUID PRIMARY KEY
	, tag_name TEXT NOT NULL
	, fk_category_id UUID NOT NULL

	, FOREIGN KEY(fk_category_id) REFERENCES tag_categories(category_id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE users_tags (
	id UUID PRIMARY KEY
	, fk_tag_id UUID NOT NULL
	, fk_user_id UUID NOT NULL

	, FOREIGN KEY(fk_tag_id) REFERENCES tags(tag_id) ON UPDATE CASCADE ON DELETE CASCADE
	, FOREIGN KEY(fk_user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE chatrooms (
	chatroom_id UUID PRIMARY KEY
);

CREATE TABLE chatrooms_users (
	fk_user_id UUID NOT NULL
	, fk_chatroom_id UUID NOT NULL
	
	, FOREIGN KEY(fk_user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE
	, FOREIGN KEY(fk_chatroom_id) REFERENCES chatrooms(chatroom_id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE messages (
	message_id UUID PRIMARY KEY
	, fk_sender_id UUID NOT NULL
	, fk_chatroom_id UUID NOT NULL
	, content TEXT NOT NULL
	, created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()

	, FOREIGN KEY(fk_sender_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE
	, FOREIGN KEY(fk_chatroom_id) REFERENCES chatrooms(chatroom_id) ON UPDATE CASCADE ON DELETE CASCADE
);

ALTER TABLE chatrooms ADD COLUMN fk_last_message_id UUID;

ALTER TABLE chatrooms ADD CONSTRAINT
	last_message FOREIGN KEY(fk_last_message_id) REFERENCES messages(message_id) ON UPDATE CASCADE ON DELETE CASCADE;

CREATE TABLE portifolios (
	portifolio_id UUID PRIMARY KEY
	, fk_user_id UUID NOT NULL
	, title TEXT NOT NULL
	, description TEXT NOT NULL
	, photos TEXT []

	, FOREIGN KEY(fk_user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE events (
	id UUID PRIMARY KEY
	, title TEXT NOT NULL
	, description TEXT NOT NULL
	, price INTEGER NOT NULL DEFAULT 0
	, cover_path TEXT NOT NULL DEFAULT '/assets/img/events-1.png'

	, maps_link TEXT
	, address TEXT
	, address2 TEXT
	, city TEXT
	, state TEXT

	, date TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

CREATE TABLE user_event_confirmations (
	fk_user_id UUID NOT NULL
	, fk_event_id UUID NOT NULL

	, FOREIGN KEY(fk_user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE
	, FOREIGN KEY(fk_event_id) REFERENCES events(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE event_speakers (
	fk_user_id UUID NOT NULL
	, fk_event_id UUID NOT NULL

	, FOREIGN KEY(fk_user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE
	, FOREIGN KEY(fk_event_id) REFERENCES events(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE meetings (
	fk_user1_id UUID NOT NULL
	, fk_user2_id UUID NOT NULL
	, meeting_timestamp TIMESTAMP WITH TIME ZONE NOT NULL
	, status TEXT NOT NULL DEFAULT 'invited'

	, PRIMARY KEY(fk_user1_id, fk_user2_id, meeting_timestamp)
	, FOREIGN KEY(fk_user1_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE
	, FOREIGN KEY(fk_user2_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE reviews (
	fk_reviewer_id UUID NOT NULL
	, fk_recipient_id UUID NOT NULL
	, description TEXT NOT NULL
	, grade INTEGER NOT NULL

	, FOREIGN KEY(fk_reviewer_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE
	, FOREIGN KEY(fk_recipient_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE
);
