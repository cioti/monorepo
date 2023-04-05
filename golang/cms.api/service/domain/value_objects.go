package domain

type ApiID string

type ModelField struct {
	Name        string
	ApiID       ApiID
	Description string
	Type        string
}

type ContentField struct {
	Name        string
	ApiID       ApiID
	Description string
	Type        string
	Value       interface{}
}
