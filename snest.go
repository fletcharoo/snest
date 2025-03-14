package snest

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

// Load accepts a pointer to a struct and populates its fields with environment
// variables defined by the snest tag.
// If input is not a pointer, Load will return an error.
// If input is not a struct, Load will return an error.
func Load(input any) (err error) {
	val := reflect.ValueOf(input)
	if val.Kind() != reflect.Ptr {
		err = fmt.Errorf("input must be a pointer")
		return
	}
	if val.Elem().Kind() != reflect.Struct {
		err = fmt.Errorf("input must be a struct")
		return
	}

	val = val.Elem()
	typ := val.Type()
	for i := range val.NumField() {
		field := typ.Field(i)
		fieldVal := val.Field(i)

		if !fieldVal.CanSet() {
			err = fmt.Errorf("field %s can't be set", field.Name)
			return
		}

		envName, ok := field.Tag.Lookup("snest")
		if !ok {
			continue
		}

		envValue := os.Getenv(envName)
		switch fieldVal.Kind() {
		case reflect.String:
			envValue = `"` + envValue + `"`
		}

		fieldPtr := reflect.New(fieldVal.Type()).Interface()
		if err = json.Unmarshal([]byte(envValue), fieldPtr); err != nil {
			err = fmt.Errorf("failed to unmarshal to field %q: %w", field.Name, err)
			return
		}

		fieldVal.Set(reflect.ValueOf(fieldPtr).Elem())
	}

	return nil
}
