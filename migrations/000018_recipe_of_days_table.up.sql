CREATE TABLE IF NOT EXISTS recipe_of_days(
    id SERIAL PRIMARY KEY, -- Уникальный идентификатор.
    title VARCHAR(255) NOT NULL, -- Название меню.
    content JSONB NOT NULL, -- Рецепты.
    created_at TIMESTAMP NOT NULL DEFAULT NOW(), -- Дата создания записи.
    updated_at TIMESTAMP NOT NULL DEFAULT NOW() -- Дата обновления записи.
);