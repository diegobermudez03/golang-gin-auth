CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    password TEXT NOT NULL, 
    email TEXT UNIQUE NOT NULL,
    phone TEXT UNIQUE,
    token TEXT,
    user_type TEXT CHECK (user_type IN ('ADMIN', 'USER')),
    refresh_token TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
