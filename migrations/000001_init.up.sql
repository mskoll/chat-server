CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    username   VARCHAR(255)            NOT NULL,
    created_at TIMESTAMP DEFAULT now() NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE chats
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(50)             NOT NULL,
    created_at TIMESTAMP DEFAULT now() NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE user_chat
(
    id         SERIAL PRIMARY KEY,
    user_id    INT                     NOT NULL,
    chat_id    INT                     NOT NULL,
    created_at TIMESTAMP DEFAULT now() NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (chat_id) REFERENCES chats (id)
);

CREATE TABLE messages
(
    id         SERIAL PRIMARY KEY,
    chat_id    INT                     NOT NULL,
    user_id    INT                     NOT NULL,
    content    TEXT                    NOT NULL,
    created_at TIMESTAMP DEFAULT now() NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (chat_id) REFERENCES chats (id)
);