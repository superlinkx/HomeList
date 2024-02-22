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
(1, 'Hello world 1!', 1024),
(1, 'Hello world 2!', 2048),
(1, 'Hello world 3!', 3072),
(1, 'Hello world 4!', 4096),
(1, 'Hello world 5!', 5120),
(1, 'Hello world 6!', 6144),
(1, 'Hello world 7!', 7168),
(1, 'Hello world 8!', 8192),
(1, 'Hello world 9!', 9216),
(1, 'Hello world 10!', 10240),
(1, 'Hello world 11!', 11264),
(1, 'Hello world 12!', 12288),
(1, 'Hello world 13!', 13312),
(1, 'Hello world 14!', 14336),
(1, 'Hello world 15!', 15360);

-- +migrate Down
DROP TABLE items;
DROP TABLE lists;
