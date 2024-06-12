CREATE TABLE IF NOT EXISTS stats (
    user_id BIGINT,
    ticket_id BIGINT,
    type TEXT
)
ENGINE = MergeTree()
PRIMARY KEY (type, ticket_id)
