CREATE TABLE IF NOT EXISTS user_free_recipes_table(
    id SERIAL PRIMARY KEY, -- Уникальный идентификатор.
    telegram_id TEXT NOT NULL UNIQUE, -- Telegram id пользователя.
    free_recipes_allowed INTEGER NOT NULL DEFAULT 3, -- Разрешенное количество.
    free_recipes_used INTEGER NOT NULL DEFAULT 0, -- Использованное количество.
    created_at TIMESTAMP NOT NULL DEFAULT NOW(), -- Дата создания записи.
    updated_at TIMESTAMP NOT NULL DEFAULT NOW() -- Дата обновления записи.
);