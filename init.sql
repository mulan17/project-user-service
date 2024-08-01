CREATE TABLE users (
    id VARCHAR PRIMARY KEY,
    email VARCHAR UNIQUE NOT NULL,
    password VARCHAR NOT NULL,
    role VARCHAR NOT NULL,
    name VARCHAR,
    lastname VARCHAR,
    status VARCHAR NOT NULL DEFAULT 'active'
);