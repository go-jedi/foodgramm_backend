CREATE TABLE IF NOT EXISTS promo_codes(
    id SERIAL PRIMARY KEY, -- Уникальный идентификатор.
    code TEXT NOT NULL UNIQUE, -- Промо-код.
    discount_percent INTEGER NOT NULL, -- Процент скидки (например, 50 для 50%).
    max_uses_allowed INTEGER NOT NULL, -- Максимальное количество использований (-1 если без ограничений).
    amount_used INTEGER NOT NULL DEFAULT 0, -- Количество использований промокода.
    is_reusable BOOLEAN NOT NULL DEFAULT FALSE, -- Многоразовый ли промо-код.
    valid_from TIMESTAMP NOT NULL DEFAULT NOW(), -- Дата начала действия.
    valid_until TIMESTAMP, -- Дата окончания действия.
    created_at TIMESTAMP NOT NULL DEFAULT NOW(), -- Дата создания записи.
    updated_at TIMESTAMP NOT NULL DEFAULT NOW() -- Дата обновления записи.
);