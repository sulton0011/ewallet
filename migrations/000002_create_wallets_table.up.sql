CREATE TABLE IF NOT EXISTS wallets (
    wallet_id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(user_id),
    balance DECIMAL(15, 2) DEFAULT 0,
    is_detected BOOLEAN DEFAULT false,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL
);