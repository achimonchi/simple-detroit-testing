CREATE TABLE todos (
    id SERIAL PRIMARY KEY,
    name varchar(255),
    description varchar(255),
    created_at timestamp
);