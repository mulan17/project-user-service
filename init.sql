CREATE TABLE users (
    id SERIAL PRIMARY KEY UNIQUE,
    email VARCHAR UNIQUE NOT NULL,
    password VARCHAR NOT NULL,
    role VARCHAR NOT NULL,
    name VARCHAR,
    lastname VARCHAR,
    status VARCHAR NOT NULL DEFAULT 'active'
);

INSERT INTO users (email, password, role, name, lastname, status)
VALUES ('admin@example.com', '$2a$10$WpbBrnpLVTuAMQn86SgJ1OzJxjrz9WYFh9my8sYvyGdfG6Zvxfo8K', 'admin', 'NameAdmin', 'LastNameAdmin', 'active')
ON CONFLICT (email) DO NOTHING;