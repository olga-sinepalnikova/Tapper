-- 1_create_users_table.sql
CREATE TABLE IF NOT EXISTS users (
                       id SERIAL PRIMARY KEY,
                       username VARCHAR(50) NOT NULL,
                       email VARCHAR(100),
                       password_hash VARCHAR(255),
                       tokens INTEGER DEFAULT 0,
                       created_at TIMESTAMP DEFAULT NOW(),
                       updated_at TIMESTAMP DEFAULT NOW()
);

-- 2_create_messages_table.sql
CREATE TABLE IF NOT EXISTS messages (
                          id SERIAL PRIMARY KEY,
                          user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
                          content TEXT NOT NULL,
                          created_at TIMESTAMP DEFAULT NOW(),
                          length INTEGER GENERATED ALWAYS AS (LENGTH(content)) STORED
);

-- 3_create_transactions_table.sql
CREATE TABLE IF NOT EXISTS transactions (
                              id SERIAL PRIMARY KEY,
                              user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
                              type VARCHAR(20) NOT NULL,
                              amount INTEGER NOT NULL,
                              created_at TIMESTAMP DEFAULT NOW(),
                              description TEXT
);

-- 4_create_chats_table.sql
CREATE TABLE IF NOT EXISTS chats (
                       id SERIAL PRIMARY KEY,
                       user1_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
                       user2_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
                       created_at TIMESTAMP DEFAULT NOW(),
                       CONSTRAINT unique_chat UNIQUE (user1_id, user2_id)
);

-- 5_create_chat_members_table.sql
CREATE TABLE IF NOT EXISTS chat_members (
                              id SERIAL PRIMARY KEY,
                              chat_id INTEGER REFERENCES chats(id) ON DELETE CASCADE,
                              user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
                              joined_at TIMESTAMP DEFAULT NOW()
);