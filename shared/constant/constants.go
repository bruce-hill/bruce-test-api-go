// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	shimjson "github.com/bruce-hill/bruce-test-api-go/internal/encoding/json"
)

type Constant[T any] interface {
	Default() T
}

// ValueOf gives the default value of a constant from its type. It's helpful when
// constructing constants as variants in a one-of. Note that empty structs are
// marshalled by default. Usage: constant.ValueOf[constant.Foo]()
func ValueOf[T Constant[T]]() T {
	var t T
	return t.Default()
}

type Bird string // Always "bird"
type Cat string  // Always "cat"
type Dog string  // Always "dog"

func (c Bird) Default() Bird { return "bird" }
func (c Cat) Default() Cat   { return "cat" }
func (c Dog) Default() Dog   { return "dog" }

func (c Bird) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c Cat) MarshalJSON() ([]byte, error)  { return marshalString(c) }
func (c Dog) MarshalJSON() ([]byte, error)  { return marshalString(c) }

type constant[T any] interface {
	Constant[T]
	*T
}

func marshalString[T ~string, PT constant[T]](v T) ([]byte, error) {
	var zero T
	if v == zero {
		v = PT(&v).Default()
	}
	return shimjson.Marshal(string(v))
}
