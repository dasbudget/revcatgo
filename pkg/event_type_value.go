package revcatgo

import (
	"errors"
	"fmt"
	"strings"

	"gopkg.in/guregu/null.v4"
)

const (
	eventTypeTest                = "TEST"
	eventTypeInitialPurchase     = "INITIAL_PURCHASE"
	eventTypeNonRenewingPurchase = "NON_RENEWING_PURCHASE"
	eventTypeRenewal             = "RENEWAL"
	eventTypeProductChange       = "PRODUCT_CHANGE"
	eventTypeChancellation       = "CANCELLATION"
	eventTypeBillingIssue        = "BILLING_ISSUE"
	eventTypeSubscriberAlias     = "SUBSCRIBER_ALIAS"
	eventTypeSubscriptionPaused  = "SUBSCRIPTION_PAUSED"
)

var validEventTypeValues = []string{
	eventTypeTest,
	eventTypeInitialPurchase,
	eventTypeNonRenewingPurchase,
	eventTypeRenewal,
	eventTypeProductChange,
	eventTypeChancellation,
	eventTypeBillingIssue,
	eventTypeSubscriberAlias,
	eventTypeSubscriptionPaused,
}

type eventType struct {
	value null.String
}

func newEventType(s string) (*eventType, error) {
	if !contains(validEventTypeValues, s) {
		return &eventType{}, errors.New("eventType value should be one of the following: " + strings.Join(validEventTypeValues, ", "))
	}
	return &eventType{value: null.StringFrom(s)}, nil
}

func (e eventType) String() string {
	return e.value.ValueOrZero()
}

// MarshalJSON serializes a store to JSON.
func (e eventType) MarshalJSON() ([]byte, error) {
	return e.value.MarshalJSON()
}

// UnmarshalJSON deserializes a store from JSON
func (e *eventType) UnmarshalJSON(b []byte) error {
	v := &eventType{}
	err := v.value.UnmarshalJSON(b)
	if err != nil {
		return fmt.Errorf("failed to unmarshal the value of type: %w", err)
	}
	if !v.value.Valid {
		return errors.New("type is a required field")
	}
	_e, err := newEventType(v.value.ValueOrZero())
	if err != nil {
		return fmt.Errorf("failed to unmarshal the value of type: %w", err)
	}
	e.value = _e.value

	return nil
}
