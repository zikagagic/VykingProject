CREATE TABLE tournament (
  id CHAR(36) PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  prize_pool DECIMAL(12,2) NOT NULL,
  start_date DATETIME NOT NULL,
  end_date DATETIME NOT NULL
);
