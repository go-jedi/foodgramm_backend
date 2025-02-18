CREATE TABLE IF NOT EXISTS promo_code_uses(
    id SERIAL PRIMARY KEY, -- Уникальный идентификатор.
    promo_code_id INTEGER, -- Идентификатор промо-кода.
    telegram_id TEXT NOT NULL, -- Telegram id пользователя.
    created_at TIMESTAMP NOT NULL DEFAULT NOW(), -- Дата создания записи.
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(), -- Дата обновления записи.
    FOREIGN KEY (promo_code_id) REFERENCES promo_codes(id) ON DELETE CASCADE,
    FOREIGN KEY (telegram_id) REFERENCES users(telegram_id) ON DELETE CASCADE
);