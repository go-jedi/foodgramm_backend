CREATE TABLE IF NOT EXISTS recipe_types(
    id SERIAL PRIMARY KEY, -- Уникальный идентификатор.
    title TEXT NOT NULL, -- Название
    created_at TIMESTAMP NOT NULL DEFAULT NOW(), -- Дата создания записи.
    updated_at TIMESTAMP NOT NULL DEFAULT NOW() -- Дата обновления записи.
);