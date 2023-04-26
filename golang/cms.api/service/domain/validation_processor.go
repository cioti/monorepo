package domain

import (
	"reflect"
	"regexp"

	"github.com/cioti/monorepo/pkg/api"
)

type ValidationProcessor interface {
	Validate(value interface{}, field Field) (bool, error)
}

type validationProcessor struct {
	validators map[ValidationRuleLogic]ValidationFn
}

func NewValidationProcessor() ValidationProcessor {
	validators := make(map[ValidationRuleLogic]ValidationFn, 0)
	validators[Equal] = EqualValidation
	validators[GreatherOrEqual] = GreaterOrEqualValidation
	validators[GreatherThan] = GreaterThanValidation
	validators[LessThan] = LessThanValidation
	validators[LessOrEqual] = LessOrEqualValidation
	validators[MaxLen] = MaxLenValidation
	validators[MinLen] = MinLenValidation
	validators[Regex] = RegexValidation

	return &validationProcessor{
		validators: validators,
	}
}

func (vp validationProcessor) Validate(value interface{}, field Field) (bool, error) {
	err := validateDataType(value, field.Type.DataType)
	if err != nil {
		return false, err
	}

	for _, rule := range field.Type.ValidationRules {
		validator := vp.validators[rule.Logic]
		if validator != nil {
			isValid, err := validator(value, rule, field.Type.DataType)
			if err != nil {
				return false, err
			}
			if !isValid {
				return false, nil
			}
		}
	}

	return true, nil
}

type ValidationFn func(value interface{}, rule ValidationRule, dt DataType) (bool, error)

func EqualValidation(value interface{}, rule ValidationRule, dt DataType) (bool, error) {
	if rule.Logic != Equal {
		return false, api.NewInternalServerErrorf("Rule logic is not of type '%s'", Equal)
	}
	if value != rule.Target {
		return false, nil
	}

	return true, nil
}

func GreaterThanValidation(value interface{}, rule ValidationRule, dt DataType) (bool, error) {
	if rule.Logic != GreatherThan {
		return false, api.NewInternalServerErrorf("Rule logic is not of type '%s'", GreatherThan)
	}

	v, t, err := getComparableValues(value, rule.Target, dt)
	if err != nil {
		return false, err
	}
	if v > t {
		return true, nil
	}

	return true, nil
}

func GreaterOrEqualValidation(value interface{}, rule ValidationRule, dt DataType) (bool, error) {
	if rule.Logic != GreatherOrEqual {
		return false, api.NewInternalServerErrorf("Rule logic is not of type '%s'", GreatherOrEqual)
	}

	v, t, err := getComparableValues(value, rule.Target, dt)
	if err != nil {
		return false, err
	}
	if v >= t {
		return true, nil
	}

	return false, nil
}

func LessOrEqualValidation(value interface{}, rule ValidationRule, dt DataType) (bool, error) {
	if rule.Logic != LessOrEqual {
		return false, api.NewInternalServerErrorf("Rule logic is not of type '%s'", LessOrEqual)
	}

	v, t, err := getComparableValues(value, rule.Target, dt)
	if err != nil {
		return false, err
	}
	if v <= t {
		return true, nil
	}

	return false, nil
}

func LessThanValidation(value interface{}, rule ValidationRule, dt DataType) (bool, error) {
	if rule.Logic != LessThan {
		return false, api.NewInternalServerErrorf("Rule logic is not of type '%s'", LessThan)
	}

	v, t, err := getComparableValues(value, rule.Target, dt)
	if err != nil {
		return false, err
	}
	if v < t {
		return true, nil
	}

	return false, nil
}

func MinLenValidation(value interface{}, rule ValidationRule, dt DataType) (bool, error) {
	if rule.Logic != MinLen {
		return false, api.NewInternalServerErrorf("Rule logic is not of type '%s'", MinLen)
	}
	vstr, ok := value.(string)
	if !ok {
		return false, api.NewInternalServerErrorf("Unable to cast value to string type")
	}
	minLen, _ := rule.Target.(int)

	return len(vstr) >= minLen, nil
}

func MaxLenValidation(value interface{}, rule ValidationRule, dt DataType) (bool, error) {
	if rule.Logic != MaxLen {
		return false, api.NewInternalServerErrorf("Rule logic is not of type '%s'", MaxLen)
	}
	vstr, ok := value.(string)
	if !ok {
		return false, api.NewInternalServerErrorf("Unable to cast value to string type")
	}
	minLen, _ := rule.Target.(int)

	return len(vstr) >= minLen, nil
}

func RegexValidation(value interface{}, rule ValidationRule, dt DataType) (bool, error) {
	if rule.Logic != Regex {
		return false, api.NewInternalServerErrorf("Rule logic is not of type '%s'", MaxLen)
	}
	vstr, ok := value.(string)
	if !ok {
		return false, api.NewInternalServerErrorf("Unable to cast value to string type")
	}

	expr, _ := rule.Target.(string)
	regex, err := regexp.Compile(expr)
	if err != nil {
		return false, api.NewInternalServerError(err, "Regex expression '%s' does not compile", expr)
	}

	return regex.MatchString(vstr), nil
}

func getComparableValues(value interface{}, target interface{}, dt DataType) (v float64, t float64, err error) {
	switch dt {
	case IntType:
		vint, ok := value.(int)
		if !ok {
			return v, t, api.NewInternalServerErrorf("Value is not of type '%s'", dt)
		}
		tint, ok := value.(int)
		if !ok {
			return v, t, api.NewInternalServerErrorf("Target is not of type '%s'", dt)
		}

		return float64(vint), float64(tint), nil
	case FloatType:
		vfloat, ok := value.(float64)
		if !ok {
			return v, t, api.NewInternalServerErrorf("Value is not of type '%s'", dt)
		}
		tfloat, ok := value.(float64)
		if !ok {
			return v, t, api.NewInternalServerErrorf("Target is not of type '%s'", dt)
		}

		return vfloat, tfloat, nil
	}

	return v, t, api.NewInternalServerErrorf("Unable to get comparable values, data type must be of type '%s' or '%s'", IntType, FloatType)
}

func validateDataType(value interface{}, dt DataType) error {
	err := api.NewBadRequestErrorf("Content value dataType is not of type '%s'", dt)
	switch value.(type) {
	case string:
		if dt != StringType {
			return err
		}
	case int:
		if dt != IntType {
			return err
		}
	case float64:
		if dt != FloatType {
			return err
		}
	case bool:
		if dt != BoolType {
			return err
		}
	}

	if dt == ListType && reflect.TypeOf(value).Kind() != reflect.Slice {
		return err
	}

	return nil
}
