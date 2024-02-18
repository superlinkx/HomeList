-- +migrate Up
CREATE TABLE lists (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL
);

CREATE TABLE items (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  list_id INTEGER NOT NULL,
  content TEXT NOT NULL,
  checked BOOLEAN NOT NULL DEFAULT FALSE,
  sort INTEGER NOT NULL,
  CONSTRAINT fk_lists
    FOREIGN KEY (list_id)
    REFERENCES lists(id)
    ON DELETE CASCADE
);

INSERT INTO lists (name) VALUES ('Main List');
INSERT INTO items (list_id, content, sort) VALUES (1, 'Hello world!', 1024);

-- +migrate Down
DROP TABLE items;
DROP TABLE lists;
