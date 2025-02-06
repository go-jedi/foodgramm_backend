package recipe

import (
	"bytes"
	"errors"
	"text/template"
)

var (
	ErrNotFoundTemplateFunction    = errors.New("templates function not found for the specified type")
	ErrTemplateFunctionReturnEmpty = errors.New("templates function returned empty string")
)

type templateFunc func(GenerateRecipe) (string, error)

type GenerateRecipe struct {
	Type                  int      `json:"type"`
	Products              []string `json:"products"`
	NonConsumableProducts *string  `json:"non_consumable_products"`
	Name                  *string  `json:"name"`
	AmountCalories        *int     `json:"amount_calories"`
	AvailableProducts     []string `json:"available_products"`
}

type Template struct {
	temps map[int]templateFunc
}

func NewRecipe() *Template {
	r := &Template{}

	r.temps = map[int]templateFunc{
		1: r.getMenuForOneDay,
		2: r.getFitnessMenu,
		3: r.getAvailableProductsMenu,
		4: r.getMenuByName,
	}

	return r
}

// Generate need templates.
func (t *Template) Generate(data GenerateRecipe) (string, error) {
	tFn, ok := t.temps[data.Type]
	if !ok {
		return "", ErrNotFoundTemplateFunction
	}

	tr, err := tFn(data)
	if err != nil {
		return "", err
	}

	if len(tr) == 0 {
		return "", ErrTemplateFunctionReturnEmpty
	}

	return tr, nil
}

// executeTemplate executes the templates with the provided data.
func (t *Template) executeTemplate(tmpl string, data GenerateRecipe) (string, error) {
	tp, err := template.New("recipe").Parse(tmpl)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tp.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// GetMenuForOneDay get menu for one day.
func (t *Template) getMenuForOneDay(data GenerateRecipe) (string, error) {
	tmpl := `
Составь мне меню на день (завтрак, обед, полдник, ужин). Для каждого блюда напиши подробный рецепт: ингредиенты (в граммах/штуках), шаги приготовления, время готовки и калорийность. Учти, что у меня аллергия на {{.Products}} и я не употребляю {{.NonConsumableProducts}} — исключи их из рецептов. Блюда должны быть простыми и разнообразными.

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

БЖУ: Белки — [г], Жиры — [г], Углеводы — [г]

(Повтори структуру для обеда, полдника, ужина)

Соблюдай точные заголовки (Завтрак, Обед и т.д.), разделы (Ингредиенты, Рецепт и пр.) и форматирование. Не упоминай об аллергиях или исключенных продуктах в ответе. Не добавляй лишний текст.
`

	return t.executeTemplate(tmpl, data)
}

// GetFitnessMenu get fitness menu.
func (t *Template) getFitnessMenu(data GenerateRecipe) (string, error) {
	tmpl := `
Составь мне фитнес-меню на день (завтрак, обед, полдник, ужин) с общим калоражем {{.AmountCalories}} ккал. Укажи для каждого блюда: точные граммовки, шаги, БЖУ и калорийность. Учти, что у меня аллергия на {{.Products}} и я не употребляю {{.NonConsumableProducts}} — исключи их из рецептов. Сделай упор на белок и клетчатку.»

Структура ответа:
Оформи ответ в следующем формате:

Фитнес-меню
Завтрак: [Название блюда]

Общий калораж: [число] ккал  

Ингредиенты:

[Продукт 1] — [количество в граммах/штуках]

[Продукт 2] — [количество]

Рецепт:

[Шаг 1]

[Шаг 2]
...

Время готовки: [время в минутах]

Калорийность: [ккал]

БЖУ: Белки — [г], Жиры — [г], Углеводы — [г]

(Повтори структуру для обеда, полдника, ужина)

Соблюдай точные заголовки (Завтрак, Обед и т.д.), разделы (Ингредиенты, Рецепт и пр.) и форматирование. Не упоминай об аллергиях или исключенных продуктах в ответе. Не добавляй лишний текст.
`

	return t.executeTemplate(tmpl, data)
}

// GetAvailableProductsMenu get menu of available products.
func (t *Template) getAvailableProductsMenu(data GenerateRecipe) (string, error) {
	tmpl := `
У меня есть: {{.AvailableProducts}}. Придумай рецепт, используя только эти ингредиенты. Напиши пошагово с количествами, временем и калорийностью. Учти, что у меня аллергия на {{.Products}} и я не употребляю {{.NonConsumableProducts}} — исключи их из рецептов.»

Структура ответа:
Оформи ответ в следующем формате:

Меню из имеющихся продуктов
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

БЖУ: Белки — [г], Жиры — [г], Углеводы — [г]

Соблюдай точные разделы (Ингредиенты, Рецепт и пр.) и форматирование. Не упоминай об аллергиях или исключенных продуктах в ответе. Не добавляй лишний текст.
`

	return t.executeTemplate(tmpl, data)
}

// GetMenuByName get menu by name.
func (t *Template) getMenuByName(data GenerateRecipe) (string, error) {
	tmpl := `
Я хочу приготовить {{.Name}}. Напиши рецепт: ингредиенты (граммы/штуки), шаги, время, сложность, калорийность. Учти, что у меня аллергия на {{.Products}} и я не употребляю {{.NonConsumableProducts}} — исключи их из рецептов. Предложи замены ингредиентов, если это необходимо, но без указанных ограничений.»

Структура ответа:
Оформи ответ в следующем формате:

Получить рецепт
Завтрак: {{.Name}}

Ингредиенты:

[Продукт 1] — [количество в граммах/штуках]

[Продукт 2] — [количество]

Рецепт:

[Шаг 1]

[Шаг 2]
...

Время готовки: [время в минутах]

Калорийность: [ккал]

БЖУ: Белки — [г], Жиры — [г], Углеводы — [г]

Соблюдай точные разделы (Ингредиенты, Рецепт и пр.) и форматирование. Не упоминай об аллергиях или исключенных продуктах в ответе. Не добавляй лишний текст. Также, не обязательно использовать все продукты в блюдах, используй продукты адекватно и предлагай адекватные варианты блюд
`

	return t.executeTemplate(tmpl, data)
}
