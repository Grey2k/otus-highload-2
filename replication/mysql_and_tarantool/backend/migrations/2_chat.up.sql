CREATE TABLE IF NOT EXISTS chat (
    id          VARCHAR(100) NOT NULL DEFAULT (uuid()) PRIMARY KEY,
    create_time TIMESTAMP NOT NULL
) ENGINE=InnoDB;

