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
	CancelReasonUnsubscribe        = "UNSUBSCRIBE"
	CancelReasonBillingError       = "BILLING_ERROR"
	CancelReasonDeveloperInitiated = "DEVELOPER_INITIATED"
	CancelReasonPriceIncrease      = "PRICE_INCREASE"
	CancelReasonCustomerSupport    = "CUSTOMER_SUPPORT"
	CancelReasonUnknown            = "UNKNOWN"
)

var validCancelReasonValues = []string{
	CancelReasonUnsubscribe,
	CancelReasonBillingError,
	CancelReasonDeveloperInitiated,
	CancelReasonPriceIncrease,
	CancelReasonCustomerSupport,
	CancelReasonUnknown,
}

func newCancelReason(v string) (*cancelReason, error) {
	if v != "" && !contains(validCancelReasonValues, v) {
		return &cancelReason{}, errors.New("cancelReason value should be one of the following:" + strings.Join(validCancelReasonValues, ","))
	}
	return &cancelReason{value: null.NewString(v, v != "")}, nil
}

func (c cancelReason) String() string {
	return c.value.ValueOrZero()
}

func (c cancelReason) NullString() null.String {
	return c.value
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
	_c, err := newCancelReason(v.value.ValueOrZero())
	if err != nil {
		return fmt.Errorf("failed to unmarshal the value of cancel_reason: %w", err)
	}
	c.value = _c.value

	return nil
}
