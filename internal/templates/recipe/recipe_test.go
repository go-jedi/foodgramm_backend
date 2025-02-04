package recipe

import (
	"errors"
	"testing"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/recipe"
)

func TestGenerate_MenuForOneDay(t *testing.T) {
	r := NewRecipe()
	data := recipe.GenerateRecipeDTO{
		Type:      1,
		Allergies: "шоколад",
		Products:  []string{"морковь", "лук"},
	}

	expected := `
Составь мне меню на день (завтрак, обед, полдник, ужин). Для каждого блюда напиши подробный рецепт: ингредиенты (в граммах/штуках), шаги приготовления, время готовки и калорийность. Учти, что у меня аллергия на шоколад и я не употребляю [морковь лук] — исключи их из рецептов. Блюда должны быть простыми и разнообразными.

Структура ответа:
Оформи ответ в следующем формате:

Меню на день
Завтрак: [Название блюда]

Ингредиенты:

[Продукт 1] — [количество в граммах/штуках]

[Продукт 2] — [количество]

Рецепт:

[Шаг 1]

[Шаг 2]
...

Время готовки: [время в минутах]

Калорийность: [ккал]

(Повтори структуру для обеда, полдника, ужина)

Соблюдай точные заголовки (Завтрак, Обед и т.д.), разделы (Ингредиенты, Рецепт и пр.) и форматирование. Не упоминай об аллергиях или исключенных продуктах в ответе. Не добавляй лишний текст.
	`

	result, err := r.Generate(data)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	t.Log("result:", result)
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestGenerate_MenuForTheWeek(t *testing.T) {
	r := NewRecipe()
	data := recipe.GenerateRecipeDTO{
		Type:      2,
		Allergies: "шоколад",
		Products:  []string{"морковь", "лук"},
	}

	expected := `
Создай меню на 7 дней с 4 приемами пищи ежедневно. Для каждого дня предложи уникальные рецепты (без повторов) с точным количеством ингредиентов, инструкцией, временем готовки и калорийностью. Учти, что у меня аллергия на шоколад и я не употребляю [морковь лук] — исключи их из рецептов.

Структура ответа:

markdown
Copy
### Меню на неделю  

День 1  
Завтрак: [Название блюда]  
- Ингредиенты:  
  - [Продукт] — [граммы/штуки]  
  ...  
- Рецепт:  
  1. [Шаг]  
  ...  
- Время готовки: [минуты]  
- Калорийность: [ккал]  

(Повтори для остальных приемов пищи и дней 2-7)  
(Исключить упоминание аллергенов и ограничений в ответе)
	`

	result, err := r.Generate(data)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestGenerate_FitnessMenu(t *testing.T) {
	r := NewRecipe()
	calories := 2000
	data := recipe.GenerateRecipeDTO{
		Type:           3,
		Allergies:      "шоколад",
		Products:       []string{"морковь", "лук"},
		AmountCalories: &calories,
	}

	expected := `
Составь фитнес-меню на день (4 приема пищи) с общим калоражем 2000 ккал. Укажи для каждого блюда: точные граммовки, шаги, БЖУ и калорийность. Учти, что у меня аллергия на шоколад и я не употребляю [морковь лук] — исключи их из рецептов. Сделай упор на белок и клетчатку.»

Структура ответа:

markdown
Copy
### Фитнес-меню  

Общий калораж: [число] ккал  

Завтрак: [Название блюда]  
- Ингредиенты:  
  - [Продукт] — [граммы]  
  ...  
- Рецепт:  
  1. [Шаг]  
  ...  
- Время готовки: [минуты]  
- БЖУ: Белки — [г], Жиры — [г], Углеводы — [г]  
- Калорийность: [ккал]  

(Повтори для остальных приемов пищи)
`

	result, err := r.Generate(data)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestGenerate_AvailableProductsMenu(t *testing.T) {
	r := NewRecipe()
	data := recipe.GenerateRecipeDTO{
		Type:              4,
		Allergies:         "шоколад",
		Products:          []string{"морковь", "лук"},
		AvailableProducts: []string{"яблоко", "банан"},
	}

	expected := `
У меня есть: [яблоко банан]. Придумай рецепт, используя только эти ингредиенты. Напиши пошагово с количествами, временем и калорийностью. Учти, что у меня аллергия на шоколад и я не употребляю [морковь лук] — исключи их из рецептов.»

Структура ответа:

markdown
Copy
### Рецепт из доступных продуктов  

Использованные ингредиенты:  
- [Продукт 1]  
- [Продукт 2]  
...  

Блюдо: [Название]  
- Ингредиенты:  
  - [Продукт] — [граммы/штуки]  
  ...  
- Рецепт:  
  1. [Шаг]  
  ...  
- Время готовки: [минуты]  
- Калорийность: [ккал]
`

	result, err := r.Generate(data)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestGenerate_MenuByName(t *testing.T) {
	r := NewRecipe()
	name := "Паста"
	data := recipe.GenerateRecipeDTO{
		Type:      5,
		Allergies: "шоколад",
		Products:  []string{"морковь", "лук"},
		Name:      &name,
	}

	expected := `
Я хочу приготовить Паста. Напиши рецепт: ингредиенты (граммы/штуки), шаги, время, сложность, калорийность. Учти, что у меня аллергия на шоколад и я не употребляю [морковь лук] — исключи их из рецептов. Предложи замены ингредиентов, если это необходимо, но без указанных ограничений.»

Структура ответа:

markdown
Copy
### Рецепт: [Название блюда]  

- Ингредиенты:  
  - [Продукт] — [граммы/штуки]  
  ...  
- Рецепт:  
  1. [Шаг]  
  ...  
- Время готовки: [минуты]  
- Сложность: [легкая/средняя/сложная]  
- Калорийность: [ккал]  

Возможные замены:  
- [Ингредиент] → [альтернатива]  
(Не упоминать аллергии/исключения, только замены)
`

	result, err := r.Generate(data)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestGenerate_InvalidType(t *testing.T) {
	r := NewRecipe()
	data := recipe.GenerateRecipeDTO{
		Type:      6,
		Allergies: "шоколад",
		Products:  []string{"морковь", "лук"},
	}

	_, err := r.Generate(data)
	if !errors.Is(err, ErrNotFoundTemplateFunction) {
		t.Errorf("Expected error %v, got %v", ErrNotFoundTemplateFunction, err)
	}
}
