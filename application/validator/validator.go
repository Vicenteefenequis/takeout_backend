package validator

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Input struct {
	tagMap       map[string]interface{}
	index        int
	value        reflect.Value
	reflectInput reflect.Value
}

type Validator struct {
}

func (v *Validator) Validate(input interface{}) error {
	inputData := reflect.ValueOf(input)

	for i := 0; i < inputData.NumField(); i++ {
		_input := &Input{
			index:        i,
			tagMap:       map[string]interface{}{},
			value:        inputData.Field(i),
			reflectInput: inputData,
		}

		_input.prepareTagMap()

		err := _input.isRequired()
		if err != nil {
			return err
		}

		err = _input.maxLength()
		if err != nil {
			return err
		}

	}

	return nil

}

func (i *Input) maxLength() error {
	if i.tagMap["maxlength"] != nil {
		maxLength, _ := strconv.Atoi(fmt.Sprintf("%v", i.tagMap["maxlength"]))
		if i.value.Len() > maxLength {
			return errors.New(fmt.Sprintf("%v Tem que ser no maximo: %v caracteres", i.tagMap["label"], maxLength))
		}
	}
	return nil
}

func (i *Input) isRequired() error {
	if i.tagMap["required"] == true && i.isZeroValue(i.value.Interface(), i.value.Type()) {
		return errors.New(fmt.Sprintf("%v e Requerido", i.tagMap["label"]))
	}
	return nil
}

func (i *Input) isZeroValue(value interface{}, fieldType reflect.Type) bool {
	return reflect.DeepEqual(value, reflect.Zero(fieldType).Interface())
}

func (i *Input) prepareTagMap() {
	validateTags := i.reflectInput.Type().Field(i.index).Tag.Get("validate")
	tags := strings.Split(validateTags, ",")

	for _, tag := range tags {
		tagKeyValue := strings.Split(tag, ":")
		if len(tagKeyValue) > 1 {
			i.tagMap[tagKeyValue[0]] = tagKeyValue[1]
		}

		if len(tagKeyValue) == 1 {
			i.tagMap[tagKeyValue[0]] = true
		}

	}
}
