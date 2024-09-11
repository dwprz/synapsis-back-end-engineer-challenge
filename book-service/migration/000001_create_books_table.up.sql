CREATE TABLE IF NOT EXISTS books (
    book_id SERIAL NOT NULL PRIMARY KEY,
    title VARCHAR(100) NOT NULL UNIQUE,
    author VARCHAR(100) NOT NULL,
    isbn VARCHAR(50) NOT NULL,
    synopsis TEXT NULL,
    published_year INT NOT NULL,
    stock INT NOT NULL,
    location VARCHAR(100) NOT NULL DEFAULT 'Unknown',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS product_title_fts ON books USING GIN ((to_tsvector('indonesian', title)));

CREATE INDEX IF NOT EXISTS author_hash_index ON books USING HASH (author);