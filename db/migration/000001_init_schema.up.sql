CREATE TABLE "books" (
    "id" bigserial PRIMARY KEY,
    "title" varchar NOT NULL,
    "author" varchar NOT NULL,
    "price" bigint NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

INSERT INTO  books (title, author, price)
VALUES ('Harry Potter', 'J. K. Rowling', 4000),
       ('Moby Dick', 'Herman Melville',2200),
       ('Don Quixote', 'Miguel de Cervantes', 6220);
