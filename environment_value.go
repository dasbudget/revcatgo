package revcatgo

import (
	"errors"
	"fmt"
	"strings"

	"gopkg.in/guregu/null.v4"
)

type environment struct {
	value null.String
}

const (
	EnvironmentSandbox    = "SANDBOX"
	EnvironmentProduction = "PRODUCTION"
)

var validEnvironmentValues = []string{
	EnvironmentSandbox,
	EnvironmentProduction,
}

func newEnvironment(s string) (*environment, error) {
	if s != "" && !contains(validEnvironmentValues, s) {
		return &environment{}, errors.New("environment value should be one of the following:" + strings.Join(validEnvironmentValues, ","))
	}
	return &environment{value: null.NewString(s, s != "")}, nil
}

func (e environment) String() string {
	return e.value.ValueOrZero()
}

func (e *environment) IsProduction() bool {
	return e.value.ValueOrZero() == EnvironmentProduction
}

func (e *environment) IsSandBox() bool {
	return e.value.ValueOrZero() == EnvironmentSandbox
}

// MarshalJSON serializes a store to JSON.
func (e environment) MarshalJSON() ([]byte, error) {
	return e.value.MarshalJSON()
}

// UnmarshalJSON deserializes a store from JSON
func (e *environment) UnmarshalJSON(b []byte) error {
	v := &environment{}
	err := v.value.UnmarshalJSON(b)
	if err != nil {
		return fmt.Errorf("failed to unmarshal the value of environment: %w", err)
	}

	_e, err := newEnvironment(v.value.ValueOrZero())
	if err != nil {
		return fmt.Errorf("failed to unmarshal the value of environment: %w", err)
	}
	e.value = _e.value

	return nil
}
