CREATE TABLE player (
  id CHAR(36) PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  account_balance DECIMAL(12,2) NOT NULL
);
