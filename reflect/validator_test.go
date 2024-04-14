package reflect

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

type UserDto struct {
	Name string `validate:"max-len=5,allow-symbol=false"`
}

// validateStructReflective validates fields of a struct based on `validate` tags.
func validateStructReflective(s interface{}) error {
	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Ptr {
		val = val.Elem() // Handle pointer to struct by dereferencing.
	}

	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		tag := fieldType.Tag.Get("validate")
		if tag == "" {
			continue // No validation rules for this field.
		}

		// Process the tag string and perform validations.
		if err := validateField(field, tag); err != nil {
			return fmt.Errorf("%s: %v", fieldType.Name, err)
		}
	}
	return nil
}

// validateField checks a single field
func validateField(field reflect.Value, tag string) error {
	rules := strings.Split(tag, ",")
	for _, rule := range rules {
		parts := strings.Split(rule, "=")
		if len(parts) != 2 {
			return fmt.Errorf("invalid rule format")
		}
		key, value := parts[0], parts[1]
		err := validateTag(key, value, field)
		if err != nil {
			return err
		}
	}
	return nil
}

// validateTag checks a single field against its validation rules.
func validateTag(key string, value string, field reflect.Value) error {
	switch key {
	case "max-len":
		maxLen, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("invalid max-len value")
		}
		if field.Kind() == reflect.String && len(field.String()) > maxLen {
			return fmt.Errorf("length greater than %d", maxLen)
		}
	case "allow-symbol":
		allowSymbol, err := strconv.ParseBool(value)
		if err != nil {
			return fmt.Errorf("invalid symbol value")
		}
		if !allowSymbol && strings.ContainsAny(field.String(), "!@#$%^&*()") {
			return fmt.Errorf("symbols are not allowed")
		}
	}
	return nil
}

func validateStructNonReflective(d UserDto, rules UserDtoRules) error {
	if len(d.Name) > rules.MaxLen {
		return fmt.Errorf("length greater than %d", rules.MaxLen)
	}
	if !rules.AllowSymbol && strings.ContainsAny(d.Name, "!@#$%^&*()") {
		return fmt.Errorf("symbols are not allowed")
	}
	return nil
}

type UserDtoRules struct {
	MaxLen      int
	AllowSymbol bool
}

// Benchmark for reflective validation.
func BenchmarkValidateStructReflect(b *testing.B) {
	user := UserDto{Name: "John#"}
	for n := 0; n < b.N; n++ {
		_ = validateStructReflective(&user)
	}
}

// Benchmark for non-reflective validation.
func BenchmarkValidateStructNoReflect(b *testing.B) {
	user := UserDto{Name: "John#"}
	rules := UserDtoRules{MaxLen: 5, AllowSymbol: false}
	for n := 0; n < b.N; n++ {
		_ = validateStructNonReflective(user, rules)
	}
}
