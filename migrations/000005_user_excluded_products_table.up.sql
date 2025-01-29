CREATE TABLE IF NOT EXISTS user_excluded_products_table(
    id SERIAL PRIMARY KEY, -- Уникальный идентификатор.
    user_id INTEGER, -- Идентификатор пользователя.
    telegram_id TEXT, -- Telegram id пользователя.
    products TEXT[] NOT NULL DEFAULT '{}', -- Исключенные продукты пользователя.
    created_at TIMESTAMP NOT NULL DEFAULT NOW(), -- Дата создания записи.
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(), -- Дата обновления записи.
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (telegram_id) REFERENCES users(telegram_id) ON DELETE CASCADE
);