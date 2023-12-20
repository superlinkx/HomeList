-- +migrate Up
CREATE TABLE lists (
  id BIGSERIAL PRIMARY KEY,
  name TEXT NOT NULL
);

CREATE TABLE list_items (
  id BIGSERIAL PRIMARY KEY,
  list_id BIGINT NOT NULL,
  content TEXT NOT NULL,
  checked BOOLEAN NOT NULL DEFAULT FALSE,
  sort BIGINT NOT NULL,
  CONSTRAINT fk_lists
    FOREIGN KEY (list_id)
    REFERENCES lists(id)
    ON DELETE CASCADE
);

INSERT INTO lists (name) VALUES ('Main List');
INSERT INTO list_items (list_id, content, sort) VALUES (1, 'Hello world!', 1024);

-- +migrate Down
DROP TABLE list_items;
DROP TABLE lists;
