package revcatgo

import (
	"errors"
	"fmt"
	"strings"

	"gopkg.in/guregu/null.v4"
)

type cancelReason struct {
	value null.String
}

const (
	cancelReasonUnsubscribe        = "UNSUBSCRIBE"
	cancelReasonBillingError       = "BILLING_ERROR"
	cancelReasonDeveloperInitiated = "DEVELOPER_INITIATED"
	cancelReasonPriceIncrease      = "PRICE_INCREASE"
	cancelReasonCustomerSupport    = "CUSTOMER_SUPPORT"
	cancelReasonUnknown            = "UNKNOWN"
)

var validcancelReasonValues = []string{
	cancelReasonUnsubscribe,
	cancelReasonBillingError,
	cancelReasonDeveloperInitiated,
	cancelReasonPriceIncrease,
	cancelReasonCustomerSupport,
	cancelReasonUnknown,
}

func newCancelReason(v string) (*cancelReason, error) {
	if !contains(validcancelReasonValues, v) {
		return &cancelReason{}, errors.New("cancelReason value should be one of the following:" + strings.Join(validcancelReasonValues, ","))
	}
	return &cancelReason{value: null.StringFrom(v)}, nil
}

func (c cancelReason) String() string {
	return c.value.ValueOrZero()
}

// MarshalJSON serializes a store to JSON.
func (c cancelReason) MarshalJSON() ([]byte, error) {
	return c.value.MarshalJSON()
}

// UnmarshalJSON deserializes a store from JSON
func (c *cancelReason) UnmarshalJSON(b []byte) error {
	v := &environment{}
	err := v.value.UnmarshalJSON(b)
	if err != nil {
		return fmt.Errorf("failed to unmarshal the value of cancel_reason: %w", err)
	}
	if !v.value.Valid {
		return errors.New("cancel_reason is a required field")
	}
	_c, err := newCancelReason(v.value.ValueOrZero())
	if err != nil {
		return fmt.Errorf("failed to unmarshal the value of cancel_reason: %w", err)
	}
	c.value = _c.value

	return nil
}
