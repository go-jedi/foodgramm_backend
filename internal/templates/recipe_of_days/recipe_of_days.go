package recipeofdays

import "github.com/go-jedi/foodgrammm-backend/internal/domain/templates"

//go:generate mockery --name=ITemplate --output=mocks --case=underscore
type ITemplate interface {
	Generate() (templates.ApplyDataToTemplateResponse, error)
}

type Template struct{}

func NewTemplate() *Template {
	return &Template{}
}

func (t *Template) Generate() (templates.ApplyDataToTemplateResponse, error) {
	str, err := t.getRecipeOfDays()
	if err != nil {
		return templates.ApplyDataToTemplateResponse{}, err
	}

	return templates.ApplyDataToTemplateResponse{
		Content: str,
	}, nil
}

// getRecipeOfDays get menu recipe of days.
func (t *Template) getRecipeOfDays() (string, error) {
	tmpl := `
Выдай один неочевидный кулинарный лайфхак и один случайный рецепт блюда. Ответ должен быть максимально подробным, структурированным и без лишнего текста. Строго соблюдай указанную ниже структуру и форматирование:

1. Лайфхак:
Название: [Название лайфхака]
Описание: [Подробное описание лайфхака, включающее интересные детали и советы]


2. Меню:
Блюдо: [Название блюда]


Ингредиенты:
[Продукт 1] — [количество в граммах/штуках]
[Продукт 2] — [количество]
...

Рецепт:
[Шаг 1]
[Шаг 2]
...
Время готовки: [количество минут]
Калорийность: [ккал]
БЖУ: Белки — [г], Жиры — [г], Углеводы — [г]

Не добавляй ничего, кроме перечисленных разделов и формата.
`

	return tmpl, nil
}
