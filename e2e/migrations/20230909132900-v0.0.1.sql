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

INSERT INTO lists (name)
VALUES
('List 1'),
('List 2'),
('List 3'),
('List 4'),
('List 5'),
('List 6'),
('List 7'),
('List 8'),
('List 9'),
('List 10'),
('List 11'),
('List 12'),
('List 13'),
('List 14'),
('List 15');

INSERT INTO items (list_id, content, sort)
VALUES
(1, 'Hello world!', 1024),
(2, 'Hello world!', 2048),
(3, 'Hello world!', 3072),
(4, 'Hello world!', 4096),
(5, 'Hello world!', 5120),
(6, 'Hello world!', 6144),
(7, 'Hello world!', 7168),
(8, 'Hello world!', 8192),
(9, 'Hello world!', 9216),
(10, 'Hello world!', 10240),
(11, 'Hello world!', 11264),
(12, 'Hello world!', 12288),
(13, 'Hello world!', 13312),
(14, 'Hello world!', 14336),
(15, 'Hello world!', 15360);

-- +migrate Down
DROP TABLE items;
DROP TABLE lists;
