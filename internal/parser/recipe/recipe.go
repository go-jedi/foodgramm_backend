package recipe

import (
	"bufio"
	"strings"
	"time"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/recipe"
)

const (
	breakfastTitle         = "Завтрак:"
	lunchTitle             = "Обед:"
	afternoonSnack         = "Полдник:"
	dinnerTitle            = "Ужин:"
	ingredientsTitle       = "Ингредиенты:"
	recipesTitle           = "Рецепт:"
	recipePreparationTitle = "Время готовки:"
	caloriesTitle          = "Калорийность:"
	bzhuTitle              = "БЖУ:"
)

type Parser struct {
	contents            [][]recipe.Content
	currentContent      recipe.Content
	title               string
	isIngredients       bool
	isRecipes           bool
	isRecipePreparation bool
	isCalories          bool
	isBzhu              bool
	idx                 int
}

func NewRecipe() *Parser {
	return &Parser{}
}

func (p *Parser) Reset() {
	p.contents = [][]recipe.Content{}
	p.currentContent = recipe.Content{}
	p.title = ""
	p.isIngredients = false
	p.isRecipes = false
	p.isRecipePreparation = false
	p.isCalories = false
	p.isBzhu = false
}

func (p *Parser) addCurrentContent() {
	if p.currentContent.ID != 0 {
		p.contents[p.idx] = append(p.contents[p.idx], p.currentContent)
	}
}

func (p *Parser) newElementInSlice() {
	p.contents = append(p.contents, []recipe.Content{})
}

func (p *Parser) setMealTitle(line string, title string, id int64) {
	p.addCurrentContent()

	const lastID = 4
	if p.currentContent.ID == lastID {
		p.newElementInSlice()
		p.idx++
	}

	p.currentContent = recipe.Content{ID: id, Type: strings.Split(title, ":")[0]}
	p.currentContent.Title = strings.TrimSpace(strings.Split(line, ": ")[1])
}

func (p *Parser) handleLine(line string) {
	if len(line) == 0 {
		return
	}

	if strings.HasPrefix(line, "Меню на день") {
		p.title = strings.TrimSpace(line)
		return
	}

	if strings.HasPrefix(line, "Фитнес-меню") {
		p.title = strings.TrimSpace(line)
		return
	}

	if strings.HasPrefix(line, breakfastTitle) {
		if len(p.contents) == 0 {
			p.newElementInSlice()
		}

		const id = int64(1)
		p.setMealTitle(line, breakfastTitle, id)

		return
	}

	if strings.HasPrefix(line, lunchTitle) {
		const id = int64(2)
		p.setMealTitle(line, lunchTitle, id)

		return
	}

	if strings.HasPrefix(line, afternoonSnack) {
		const id = int64(3)
		p.setMealTitle(line, afternoonSnack, id)

		return
	}

	if strings.HasPrefix(line, dinnerTitle) {
		const id = int64(4)
		p.setMealTitle(line, dinnerTitle, id)

		return
	}

	if strings.HasPrefix(line, ingredientsTitle) {
		p.isIngredients = true
		p.isRecipes = false
		p.isRecipePreparation = false
		p.isCalories = false
		p.isBzhu = false
		return
	}

	if strings.HasPrefix(line, recipesTitle) {
		p.isRecipes = true
		p.isIngredients = false
		p.isRecipePreparation = false
		p.isCalories = false
		p.isBzhu = false
		return
	}

	if strings.HasPrefix(line, recipePreparationTitle) {
		p.isRecipes = false
		p.isIngredients = false
		p.isRecipePreparation = true
		p.isCalories = false
		p.isBzhu = false
		p.currentContent.RecipePreparation = strings.TrimSpace(strings.Split(line, ": ")[1])

		return
	}

	if strings.HasPrefix(line, caloriesTitle) {
		p.isRecipes = false
		p.isIngredients = false
		p.isRecipePreparation = false
		p.isCalories = true
		p.isBzhu = false
		p.currentContent.Calories = strings.TrimSpace(strings.Split(line, ": ")[1])

		return
	}

	if strings.HasPrefix(line, bzhuTitle) {
		p.isRecipes = false
		p.isIngredients = false
		p.isRecipePreparation = false
		p.isCalories = false
		p.isBzhu = true
		p.currentContent.Bzhu = strings.TrimSpace(strings.Split(line, ": ")[1])

		return
	}

	switch {
	case p.isIngredients:
		p.currentContent.Ingredients = append(p.currentContent.Ingredients, strings.TrimSpace(line))
	case p.isRecipes:
		p.currentContent.MethodPreparation = append(p.currentContent.MethodPreparation, strings.TrimSpace(line))
	}
}

func (p *Parser) ParseRecipe(telegramID string, input string) (recipe.GenerateRecipeResponse, error) {
	s := bufio.NewScanner(strings.NewReader(input))

	for s.Scan() {
		p.handleLine(s.Text())
	}

	p.addCurrentContent()

	if err := s.Err(); err != nil {
		return recipe.GenerateRecipeResponse{}, err
	}

	grr := recipe.GenerateRecipeResponse{
		TelegramID: telegramID,
		Title:      p.title,
		Content:    p.contents,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	p.Reset()

	return grr, nil
}
