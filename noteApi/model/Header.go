package model

import (
	"gopkg.in/go-playground/validator.v8"
	"reflect"
)

type (
	Header struct {
		User     string `header:"user" binding:"required,ValidateUser"`
		Password string `header:"password"  binding:"required,ValidatePassword"`
	}
)

func ValidateHeaderUser(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if user, ok := field.Interface().(string); ok {
		if user != "admin"{
			return false
		}
	}
	return true
}

func ValidateHeaderPassword(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if password, ok := field.Interface().(string); ok {
		if password != "admin"{
			return false
		}
	}
	return true
}
