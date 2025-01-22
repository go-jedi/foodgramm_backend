CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY, -- Уникальный идентификатор категории
    name VARCHAR(255) NOT NULL -- Название категории
);