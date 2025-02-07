package recipeofdays

import "github.com/go-jedi/foodgrammm-backend/internal/domain/parser"

type Parser struct {
	contents       [][]parser.Content
	currentContent parser.Content
	title          string
}

func NewRecipe() *Parser {
	return &Parser{}
}

func (p *Parser) Reset() {}
