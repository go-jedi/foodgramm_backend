CREATE TABLE IF NOT EXISTS recipe_ingredients (
    id SERIAL PRIMARY KEY, -- Уникальный идентификатор записи
    recipe_id INT NOT NULL, -- Ссылка на рецепт
    product_id INT NOT NULL, -- Ссылка на продукт
    quantity DECIMAL(10, 2) NOT NULL, -- Количество продукта, необходимое для рецепта
    FOREIGN KEY (recipe_id) REFERENCES recipes(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);