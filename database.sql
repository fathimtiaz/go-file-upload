CREATE TABLE IF NOT EXISTS file_ (
    id SERIAL PRIMARY KEY,
    name_ VARCHAR(255),
    chunk_size INT,
    num_of_chunks INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,

    UNIQUE(name_)
);

CREATE TABLE IF NOT EXISTS file_chunk_ (
    id SERIAL PRIMARY KEY,
    file_id INT NOT NULL,
    size INT,
    index_ INT,
    data_ BYTEA,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,

    FOREIGN KEY (file_id) REFERENCES file_ (id)
);