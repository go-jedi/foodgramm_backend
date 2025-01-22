CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,  -- Уникальный идентификатор продукта
    name VARCHAR(255) NOT NULL, -- Название продукта
    category_id INT, -- Ссылка на категорию продукта
    unit VARCHAR(50), -- Единица измерения продукта (например, "граммы", "литры", "штуки"). Текстовое поле длиной до 50 символов
    FOREIGN KEY (category_id) REFERENCES categories(id)
);