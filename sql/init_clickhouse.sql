CREATE TABLE IF NOT EXISTS stats (
    user_id UInt64,
    ticket_id UInt64,
    type TEXT
)
ENGINE = MergeTree()
PRIMARY KEY (type, ticket_id)
