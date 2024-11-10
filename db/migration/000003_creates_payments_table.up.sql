CREATE TABLE IF NOT EXISTS payments (
    id VARCHAR(255) PRIMARY KEY,
    installments INTEGER NOT NULL,
    tax_percent INTEGER NOT NULL,
    partial_value INTEGER NOT NULL,
    total_value INTEGER NOT NULL,
    status VARCHAR(255) NOT NULL,
    transaction_id VARCHAR(255) NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    FOREIGN KEY (transaction_id) REFERENCES transactions(id)
);
