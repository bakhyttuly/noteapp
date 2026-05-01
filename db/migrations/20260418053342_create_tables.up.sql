CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username VARCHAR(255) NOT NULL UNIQUE,
                       password VARCHAR(255) NOT NULL
);

CREATE TABLE categories (
                            id SERIAL PRIMARY KEY,
                            name VARCHAR(255) NOT NULL
);

CREATE TABLE notes (
                       id SERIAL PRIMARY KEY,
                       title VARCHAR(255) NOT NULL,
                       content TEXT NOT NULL,
                       category_id INT NOT NULL,
                       CONSTRAINT fk_notes_category
                           FOREIGN KEY (category_id)
                               REFERENCES categories(id)
                               ON DELETE CASCADE
);