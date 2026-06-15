CREATE TABLE images (
    id SERIAL PRIMARY KEY,
    url TEXT NOT NULL
);

CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    name TEXT,
    description TEXT,
    id_image INTEGER REFERENCES images(id) ON DELETE SET NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE post_images (
    id SERIAL PRIMARY KEY,
    post_id INTEGER NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
    image_id INTEGER NOT NULL REFERENCES images(id) ON DELETE CASCADE,
    UNIQUE (post_id, image_id)
);