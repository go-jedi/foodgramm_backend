CREATE TABLE IF NOT EXISTS recipes(
    id SERIAL PRIMARY KEY, -- Уникальный идентификатор.
    telegram_id TEXT NOT NULL UNIQUE, -- Telegram id пользователя.
    title VARCHAR(255) NOT NULL, -- название меню
    content JSONB NOT NULL, -- рецепты
    created_at TIMESTAMP NOT NULL DEFAULT NOW(), -- Дата создания записи.
    updated_at TIMESTAMP NOT NULL DEFAULT NOW() -- Дата обновления записи.
);