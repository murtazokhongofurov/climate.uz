CREATE TABLE IF NOT EXISTS products(
    id SERIAL NOT NULL PRIMARY KEY,
    category_id INT NOT NULL REFERENCES categories(id),
    title TEXT NOT NULL,
    media_link TEXT NOT NULL,
    description TEXT, 
    price FLOAT NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP WITHOUT TIME ZONE
);