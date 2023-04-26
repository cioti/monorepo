package domain

type ValidationRuleLogic int
type DataType string

const (
	Equal ValidationRuleLogic = iota
	GreatherThan
	GreatherOrEqual
	LessThan
	LessOrEqual
	Regex
	MaxLen
	MinLen

	StringType DataType = "string"
	IntType    DataType = "integer"
	FloatType  DataType = "float"
	BoolType   DataType = "boolean"
	ListType   DataType = "list"
)

var ValidationRuleLogics []ValidationRuleLogic = []ValidationRuleLogic{Equal, GreatherThan, GreatherOrEqual, LessThan, LessOrEqual, Regex, MaxLen, MinLen}
