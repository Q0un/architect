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
