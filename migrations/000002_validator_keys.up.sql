CREATE TABLE IF NOT EXISTS validator_keys (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    request_id TEXT NOT NULL,
    key TEXT NOT NULL,
    fee_recipient TEXT NOT NULL,
    FOREIGN KEY (request_id) REFERENCES validator_requests(id)
    );