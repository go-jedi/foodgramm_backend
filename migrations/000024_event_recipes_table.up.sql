CREATE TABLE IF NOT EXISTS event_recipes(
    id SERIAL PRIMARY KEY, -- Уникальный идентификатор.
    type_id INTEGER, -- Тип рецепта.
    title VARCHAR(255) NOT NULL, -- название меню.
    content JSONB NOT NULL, -- рецепт.
    created_at TIMESTAMP NOT NULL DEFAULT NOW(), -- Дата создания записи.
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(), -- Дата обновления записи.
    FOREIGN KEY (type_id) REFERENCES recipe_types(id)
);