package recipeofdays

import (
	"bufio"
	"errors"
	"strings"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/parser"
)

var ErrNoRecipeFound = errors.New("recipe not found")

const (
	lifehackTitle          = "1. Лайфхак:"
	nameTitle              = "Название:"
	descriptionTitle       = "Описание:"
	menuTitle              = "2. Меню:"
	dishTitle              = "Блюдо:"
	ingredientsTitle       = "Ингредиенты:"
	recipesTitle           = "Рецепт:"
	recipePreparationTitle = "Время готовки:"
	caloriesTitle          = "Калорийность:"
	bzhuTitle              = "БЖУ:"
)

//go:generate mockery --name=IParser --output=mocks --case=underscore
type IParser interface {
	ParseRecipe(input string) (parser.ParsedRecipeOfDays, error)
}

type Parser struct {
	contents            [][]parser.Content
	currentContent      parser.Content
	lifehack            parser.Lifehack
	isLifehack          bool
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
	p.contents = [][]parser.Content{}
	p.currentContent = parser.Content{}
	p.lifehack = parser.Lifehack{}
	p.isLifehack = false
	p.isIngredients = false
	p.isRecipes = false
	p.isRecipePreparation = false
	p.isCalories = false
	p.isBzhu = false
	p.idx = 0
}

func (p *Parser) addCurrentContent() {
	if p.currentContent.ID != 0 {
		p.contents[p.idx] = append(p.contents[p.idx], p.currentContent)
	}
}

func (p *Parser) newElementInSlice() {
	p.contents = append(p.contents, []parser.Content{})
}

func (p *Parser) setMealTitle(line string, id int64) {
	p.addCurrentContent()

	const lastID = 4
	if p.currentContent.ID == lastID {
		p.newElementInSlice()
		p.idx++
	}

	p.currentContent = parser.Content{ID: id, Type: "Меню"}
	p.currentContent.Title = strings.TrimSpace(strings.Split(line, ": ")[1])
}

func (p *Parser) handleLine(line string) {
	if len(line) == 0 {
		return
	}

	if strings.HasPrefix(line, nameTitle) && p.isLifehack {
		p.lifehack.Name = strings.TrimSpace(strings.Split(line, ": ")[1])
	}

	if strings.HasPrefix(line, descriptionTitle) && p.isLifehack {
		p.lifehack.Description = strings.TrimSpace(strings.Split(line, ": ")[1])
	}

	if strings.HasPrefix(line, dishTitle) {
		if len(p.contents) == 0 {
			p.newElementInSlice()
		}

		const id = int64(1)
		p.setMealTitle(line, id)
	}

	if strings.HasPrefix(line, lifehackTitle) {
		p.isLifehack = true
		p.isIngredients = false
		p.isRecipes = false
		p.isRecipePreparation = false
		p.isCalories = false
		p.isBzhu = false
		return
	}

	if strings.HasPrefix(line, menuTitle) {
		p.isLifehack = false
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

func (p *Parser) ParseRecipe(input string) (parser.ParsedRecipeOfDays, error) {
	s := bufio.NewScanner(strings.NewReader(input))

	for s.Scan() {
		p.handleLine(s.Text())
	}

	p.addCurrentContent()

	if err := s.Err(); err != nil {
		return parser.ParsedRecipeOfDays{}, err
	}

	if len(p.lifehack.Name) == 0 || len(p.lifehack.Description) == 0 || len(p.contents) == 0 {
		return parser.ParsedRecipeOfDays{}, ErrNoRecipeFound
	}

	prod := parser.ParsedRecipeOfDays{
		Title:    "Лайфхак дня",
		Lifehack: p.lifehack,
		Content:  p.contents,
	}

	p.Reset()

	return prod, nil
}
