-- +migrate Up
CREATE TABLE lists (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name text NOT NULL
);

CREATE TABLE list_items (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  list_id INTEGER NOT NULL,
  content text,
  FOREIGN KEY (list_id) REFERENCES lists(id)
);

INSERT INTO lists (name) VALUES ('Main List');
INSERT INTO list_items (list_id, content) VALUES (1, 'Hello world!');

-- +migrate Down
DROP TABLE list_items;
DROP TABLE lists;