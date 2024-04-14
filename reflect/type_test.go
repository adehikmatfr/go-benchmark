package reflect

import (
	"reflect"
	"testing"
)

// NonReflectiveSetType sets the value of an integer without using reflection.
func NonReflectiveSetType(i *int, value int) {
	*i = value
}

// ReflectiveSetType sets the value of an integer using reflection.
func ReflectiveSetType(i interface{}, value int) {
	reflect.ValueOf(i).Elem().SetInt(int64(value))
}

// BenchmarkNoReflectSetType benchmarks the NonReflectiveSet function.
func BenchmarkNoReflectSetType(b *testing.B) {
	var x int
	for n := 0; n < b.N; n++ {
		NonReflectiveSetType(&x, n)
	}
}

// BenchmarkWithReflectSetType benchmarks the ReflectiveSet function.
func BenchmarkWithReflectSetType(b *testing.B) {
	var x int
	for n := 0; n < b.N; n++ {
		ReflectiveSetType(&x, n)
	}
}
