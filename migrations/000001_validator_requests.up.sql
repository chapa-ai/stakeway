CREATE TABLE IF NOT EXISTS validator_requests (
    id TEXT PRIMARY KEY,
    num_validators INTEGER NOT NULL,
    fee_recipient TEXT NOT NULL,
    status TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);