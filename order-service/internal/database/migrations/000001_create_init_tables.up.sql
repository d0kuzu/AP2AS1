CREATE TABLE IF NOT EXISTS chats (
    id UUID PRIMARY KEY,
    assistant_id TEXT NOT NULL,
    customer TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS messages (
    id UUID PRIMARY KEY,
    chat_id UUID REFERENCES chats(id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    role TEXT NOT NULL
);
