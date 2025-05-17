package utils

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// JSONBArray is a generic type to handle JSONB array data
type JSONBArray[T any] []T

// Value implements the driver.Valuer interface
func (j JSONBArray[T]) Value() (driver.Value, error) {
	return json.Marshal(j)
}

// Scan implements the sql.Scanner interface
func (j *JSONBArray[T]) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	// First try to unmarshal as an array
	var arr []T
	if err := json.Unmarshal(bytes, &arr); err == nil {
		*j = arr
		return nil
	}

	// If that fails, try to unmarshal as a single object
	var single T
	if err := json.Unmarshal(bytes, &single); err == nil {
		*j = []T{single}
		return nil
	}

	return errors.New("failed to unmarshal JSONB data")
}

// JSONBObject is a type to handle JSONB object data
type JSONBObject struct {
	Data interface{}
}

// Value implements the driver.Valuer interface
func (j JSONBObject) Value() (driver.Value, error) {
	return json.Marshal(j.Data)
}

// Scan implements the sql.Scanner interface
func (j *JSONBObject) Scan(value interface{}) error {
	if value == nil {
		j.Data = nil
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(bytes, &j.Data)
}
