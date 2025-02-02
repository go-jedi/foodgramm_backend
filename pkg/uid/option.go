package uid

type Option struct {
	Chars string `json:"chars" yaml:"chars"` // we use a string instead of a rune slice for convenience
	Count int    `json:"cnt" yaml:"cnt"`
}
