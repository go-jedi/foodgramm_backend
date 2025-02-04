package recipe

import (
	"bytes"
	"errors"
	"text/template"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/recipe"
)

var (
	ErrNotFoundTemplateFunction    = errors.New("template function not found for the specified type")
	ErrTemplateFunctionReturnEmpty = errors.New("template function returned empty string")
)

type templateFunc func(recipe.GenerateRecipeDTO) (string, error)

type Recipe struct {
	temps map[int]templateFunc
}

func NewRecipe() *Recipe {
	r := &Recipe{}

	r.temps = map[int]templateFunc{
		1: r.getMenuForOneDay,
		2: r.getMenuForTheWeek,
		3: r.getFitnessMenu,
		4: r.getAvailableProductsMenu,
		5: r.getMenuByName,
	}

	return r
}

// Generate need template.
func (r *Recipe) Generate(data recipe.GenerateRecipeDTO) (string, error) {
	tFn, ok := r.temps[data.Type]
	if !ok {
		return "", ErrNotFoundTemplateFunction
	}

	t, err := tFn(data)
	if err != nil {
		return "", err
	}

	if len(t) == 0 {
		return "", ErrTemplateFunctionReturnEmpty
	}

	return t, nil
}

// executeTemplate executes the template with the provided data.
func (r *Recipe) executeTemplate(tmpl string, data recipe.GenerateRecipeDTO) (string, error) {
	t, err := template.New("recipe").Parse(tmpl)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// GetMenuForOneDay get menu for one day.
func (r *Recipe) getMenuForOneDay(data recipe.GenerateRecipeDTO) (string, error) {
	tmpl := `
Составь мне меню на день (завтрак, обед, полдник, ужин). Для каждого блюда напиши подробный рецепт: ингредиенты (в граммах/штуках), шаги приготовления, время готовки и калорийность. Учти, что у меня аллергия на {{.Allergies}} и я не употребляю {{.Products}} — исключи их из рецептов. Блюда должны быть простыми и разнообразными.

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

	return r.executeTemplate(tmpl, data)
}

// GetMenuForTheWeek get menu for the week.
func (r *Recipe) getMenuForTheWeek(data recipe.GenerateRecipeDTO) (string, error) {
	tmpl := `
Создай меню на 7 дней с 4 приемами пищи ежедневно. Для каждого дня предложи уникальные рецепты (без повторов) с точным количеством ингредиентов, инструкцией, временем готовки и калорийностью. Учти, что у меня аллергия на {{.Allergies}} и я не употребляю {{.Products}} — исключи их из рецептов.

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

	return r.executeTemplate(tmpl, data)
}

// GetFitnessMenu get fitness menu.
func (r *Recipe) getFitnessMenu(data recipe.GenerateRecipeDTO) (string, error) {
	tmpl := `
Составь фитнес-меню на день (4 приема пищи) с общим калоражем {{.AmountCalories}} ккал. Укажи для каждого блюда: точные граммовки, шаги, БЖУ и калорийность. Учти, что у меня аллергия на {{.Allergies}} и я не употребляю {{.Products}} — исключи их из рецептов. Сделай упор на белок и клетчатку.»

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

	return r.executeTemplate(tmpl, data)
}

// GetAvailableProductsMenu get menu of available products.
func (r *Recipe) getAvailableProductsMenu(data recipe.GenerateRecipeDTO) (string, error) {
	tmpl := `
У меня есть: {{.AvailableProducts}}. Придумай рецепт, используя только эти ингредиенты. Напиши пошагово с количествами, временем и калорийностью. Учти, что у меня аллергия на {{.Allergies}} и я не употребляю {{.Products}} — исключи их из рецептов.»

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

	return r.executeTemplate(tmpl, data)
}

// GetMenuByName get menu by name.
func (r *Recipe) getMenuByName(data recipe.GenerateRecipeDTO) (string, error) {
	tmpl := `
Я хочу приготовить {{.Name}}. Напиши рецепт: ингредиенты (граммы/штуки), шаги, время, сложность, калорийность. Учти, что у меня аллергия на {{.Allergies}} и я не употребляю {{.Products}} — исключи их из рецептов. Предложи замены ингредиентов, если это необходимо, но без указанных ограничений.»

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

	return r.executeTemplate(tmpl, data)
}
