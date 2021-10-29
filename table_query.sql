CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    full_name TEXT NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    user_type TEXT DEFAULT('user') NOT NULL,
    created_at TIMESTAMP DEFAULT now() NOT NULL,
    updated_at TIMESTAMP
);

CREATE TABLE access_tokens(
    id SERIAL PRIMARY KEY,
    email TEXT NOT NULL,
    access_token TEXT,
    created_at TIMESTAMP DEFAULT now() NOT NULL
);

CREATE TABLE balances(
    id SERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    amount BIGINT NOT NULL
);

CREATE TABLE transfer_histories(
    id SERIAL PRIMARY KEY,
    sender_id BIGINT NOT NULL,
    receiver_id BIGINT NOT NULL,
    amount BIGINT NOT NULL,
    message TEXT,
    created_at TIMESTAMP DEFAULT now() NOT NULL
);