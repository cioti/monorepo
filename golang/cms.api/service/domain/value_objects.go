package domain

type ApiID string

type Field struct {
	Name        string
	ApiID       ApiID
	Description string
	Type        string
}

type Content struct {
	Name        string
	ApiID       ApiID
	Description string
	Type        string
	Value       interface{}
}
