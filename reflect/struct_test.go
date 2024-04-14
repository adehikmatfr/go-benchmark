package reflect

import (
	"reflect"
	"testing"
)

type Source struct {
	Name    string
	Age     int
	Balance float64
}

type Target struct {
	Name    string
	Age     int64
	Balance float32
}

// CopyDataReflective uses reflection to copy data from one struct to another.
func CopyDataReflective(src interface{}, dst interface{}) {
	srcVal := reflect.ValueOf(src)
	dstVal := reflect.ValueOf(dst).Elem()

	// Ensure srcVal points to a struct and not a pointer.
	if srcVal.Kind() == reflect.Ptr {
		srcVal = srcVal.Elem() // Dereference the pointer to get the struct.
	}

	for i := 0; i < srcVal.NumField(); i++ {
		srcField := srcVal.Field(i)
		dstField := dstVal.FieldByName(srcVal.Type().Field(i).Name)
		copyField(srcField, dstField)
	}
}

// copyField handles the setting of a single field from srcField to dstField.
func copyField(srcField, dstField reflect.Value) {
	if dstField.IsValid() && dstField.CanSet() {
		if srcField.Type() == dstField.Type() {
			dstField.Set(srcField)
		} else {
			convertAndSet(srcField, dstField)
		}
	}
}

// convertAndSet performs type conversion if necessary and sets the dstField with the converted value.
func convertAndSet(srcField, dstField reflect.Value) {
	switch dstField.Kind() {
	case reflect.Int, reflect.Int64:
		if srcField.Kind() == reflect.Int {
			dstField.SetInt(int64(srcField.Int()))
		}
	case reflect.Float32, reflect.Float64:
		if srcField.Kind() == reflect.Float64 {
			dstField.SetFloat(float64(srcField.Float()))
		}
	}
}

// This function directly copies data from one struct to another without using reflection.
func CopyDataNonReflective(src Source, dst *Target) {
	dst.Name = src.Name
	dst.Age = int64(src.Age)
	dst.Balance = float32(src.Balance)
}

// Benchmark using reflection.
func BenchmarkWithReflectCopyData(b *testing.B) {
	var target Target
	src := &Source{
		Name:    "John Doe",
		Age:     30,
		Balance: 100.5,
	}

	for i := 0; i < b.N; i++ {
		CopyDataReflective(src, &target)
	}
}

// Benchmark without using reflection.
func BenchmarkNonReflectCopyData(b *testing.B) {
	var target Target
	src := Source{
		Name:    "John Doe",
		Age:     30,
		Balance: 100.5,
	}

	for i := 0; i < b.N; i++ {
		CopyDataNonReflective(src, &target)
	}
}
