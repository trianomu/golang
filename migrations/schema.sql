CREATE TABLE customers (
    id VARCHAR(36) PRIMARY KEY,
    nik VARCHAR(50),
    full_name VARCHAR(100),
    legal_name VARCHAR(100),
    birth_place VARCHAR(100),
    birth_date DATE,
    salary BIGINT,
    photo_ktp TEXT,
    photo_selfie TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE transactions (
    id VARCHAR(36) PRIMARY KEY,
    customer_id VARCHAR(36),
    contract_number VARCHAR(100),
    otr BIGINT,
    admin_fee BIGINT,
    installment BIGINT,
    interest BIGINT,
    asset_name VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (customer_id) REFERENCES customers(id)
);

CREATE TABLE limits (
    id VARCHAR(36) PRIMARY KEY,
    customer_id VARCHAR(36),
    tenor INT,
    amount BIGINT,
    FOREIGN KEY (customer_id) REFERENCES customers(id)
);
