CREATE TABLE users (
    id serial primary key,
    name text,
    email varchar,
    password varchar
);

CREATE TABLE books (
    id serial primary key,
    name text,
    price decimal,
    publisher text,
    user_id serial,
    FOREIGN KEY (user_id) REFERENCES users (id)
);


