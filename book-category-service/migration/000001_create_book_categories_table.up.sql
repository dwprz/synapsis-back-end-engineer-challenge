CREATE TYPE "BookCategory" AS ENUM(
    'FICTION',
    'NON-FICTION',
    'SCIENCE',
    'HISTORY',
    'BIOGRAPHY',
    'TECHNOLOGY',
    'FANTASY',
    'MYSTERY',
    'THRILLER',
    'CHILDREN',
    'YOUNG ADULT',
    'ROMANCE',
    'ADVENTURE'
);

CREATE TABLE IF NOT EXISTS book_categories (
    category "BookCategory" NOT NULL,
    book_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (category, book_id)
)