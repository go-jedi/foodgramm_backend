CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY, -- Уникальный идентификатор
    telegramId BIGINT NOT NULL, -- Telegram id пользователя
    created_at TIMESTAMP NOT NULL DEFAULT NOW(), -- Дата создания записи
    updated_at TIMESTAMP NOT NULL DEFAULT NOW() -- Дата обновления записи
);