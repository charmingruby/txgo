CREATE TABLE IF NOT EXISTS gifts (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    message TEXT NOT NULL,
    base_value_points INTEGER NOT NULL,
    status VARCHAR(255) NOT NULL,
    receiver_wallet_id VARCHAR(255) NOT NULL,
    sender_wallet_id VARCHAR(255) NOT NULL,
    payment_id VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    FOREIGN KEY (payment_id) REFERENCES payments(id),
    FOREIGN KEY (receiver_wallet_id) REFERENCES wallets(id),
    FOREIGN KEY (sender_wallet_id) REFERENCES wallets(id)
);
