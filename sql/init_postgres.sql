CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    login TEXT NOT NULL,
    password TEXT NOT NULL,
    name TEXT,
    surname TEXT,
    birthday DATE,
    mail TEXT,
    phone TEXT
);

CREATE TABLE tickets (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    author_id BIGINT
);
