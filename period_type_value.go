package revcatgo

import (
	"errors"
	"fmt"
	"strings"

	"gopkg.in/guregu/null.v4"
)

const (
	PeriodTypeTrial       = "TRIAL"
	PeriodTypeIntro       = "INTRO"
	PeriodTypeNormal      = "NORMAL"
	PeriodTypePromotional = "PROMOTIONAL"
)

var validperiodTypeValues = []string{
	PeriodTypeTrial,
	PeriodTypeIntro,
	PeriodTypeNormal,
	PeriodTypePromotional,
}

type periodType struct {
	value null.String
}

func newPeriodType(s string) (*periodType, error) {
	if s != "" && !contains(validperiodTypeValues, s) {
		return &periodType{}, errors.New("periodType value should be one of the following:" + strings.Join(validperiodTypeValues, ","))
	}
	return &periodType{value: null.NewString(s, s != "")}, nil
}

func (p periodType) String() string {
	return p.value.ValueOrZero()
}

func (p periodType) MarshalJSON() ([]byte, error) {
	return p.value.MarshalJSON()
}

// UnmarshalJSON deserializes a store from JSON
func (p *periodType) UnmarshalJSON(b []byte) error {
	v := &periodType{}
	err := v.value.UnmarshalJSON(b)
	if err != nil {
		return fmt.Errorf("failed to unmarshal the value of period_type: %w", err)
	}

	_p, err := newPeriodType(strings.ToUpper(v.value.ValueOrZero()))
	if err != nil {
		return fmt.Errorf("failed to unmarshal the value of period_type: %w", err)
	}
	p.value = _p.value

	return nil
}
