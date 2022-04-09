CREATE TABLE IF NOT EXISTS wallets_incomes (
    history_id UUID PRIMARY KEY,
    wallet_id UUID REFERENCES wallets(wallet_id),
    amount DECIMAL(15, 2) NOT NULL,
    created_at TIMESTAMP NOT NULL
);
