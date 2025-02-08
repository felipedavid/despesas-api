package null

import (
	"bytes"
	"encoding/json"
	"errors"
)

type Nullable[T any] map[bool]T

// NewNullableWithValue is a convenience helper to allow constructing a `Nullable` with a given value, for instance to construct a field inside a struct, without introducing an intermediate variable
func NewNullableWithValue[T any](t T) Nullable[T] {
	var n Nullable[T]
	n.Set(t)
	return n
}

// NewNullNullable is a convenience helper to allow constructing a `Nullable` with an explicit `null`, for instance to construct a field inside a struct, without introducing an intermediate variable
func NewNullNullable[T any]() Nullable[T] {
	var n Nullable[T]
	n.SetNull()
	return n
}

var ErrValueIsNull = errors.New("value is null")
var ErrValueIsNotSpecified = errors.New("value is not specified")

// Get retrieves the underlying value, if present, and returns an error if the value was not present
func (t Nullable[T]) Get() (T, error) {
	var empty T
	if t.IsNull() {
		return empty, ErrValueIsNull
	}
	if !t.IsSpecified() {
		return empty, ErrValueIsNotSpecified
	}
	return t[true], nil
}

// MustGet retrieves the underlying value, if present, and panics if the value was not present
func (t Nullable[T]) MustGet() T {
	v, err := t.Get()
	if err != nil {
		panic(err)
	}
	return v
}

// Set sets the underlying value to a given value
func (t *Nullable[T]) Set(value T) {
	*t = map[bool]T{true: value}
}

// IsNull indicate whether the field was sent, and had a value of `null`
func (t Nullable[T]) IsNull() bool {
	_, foundNull := t[false]
	return foundNull
}

// SetNull indicate that the field was sent, and had a value of `null`
func (t *Nullable[T]) SetNull() {
	var empty T
	*t = map[bool]T{false: empty}
}

// IsSpecified indicates whether the field was sent
func (t Nullable[T]) IsSpecified() bool {
	return len(t) != 0
}

// SetUnspecified indicate whether the field was sent
func (t *Nullable[T]) SetUnspecified() {
	*t = map[bool]T{}
}

func (t Nullable[T]) MarshalJSON() ([]byte, error) {
	// if field was specified, and `null`, marshal it
	if t.IsNull() {
		return []byte("null"), nil
	}

	// if field was unspecified, and `omitempty` is set on the field's tags, `json.Marshal` will omit this field

	// otherwise: we have a value, so marshal it
	return json.Marshal(t[true])
}

func (t *Nullable[T]) UnmarshalJSON(data []byte) error {
	// if field is unspecified, UnmarshalJSON won't be called

	// if field is specified, and `null`
	if bytes.Equal(data, []byte("null")) {
		t.SetNull()
		return nil
	}
	// otherwise, we have an actual value, so parse it
	var v T
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	t.Set(v)
	return nil
}
