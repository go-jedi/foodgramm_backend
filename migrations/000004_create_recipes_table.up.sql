CREATE TABLE IF NOT EXISTS recipes (
    id SERIAL PRIMARY KEY,           -- Уникальный идентификатор рецепта.
    name VARCHAR(255) NOT NULL,      -- Название рецепта.
    description TEXT,                -- Описание рецепта.
    instructions TEXT,               -- Пошаговая инструкция.
    preparation_time INT NOT NULL    -- Время приготовления в минутах.
);