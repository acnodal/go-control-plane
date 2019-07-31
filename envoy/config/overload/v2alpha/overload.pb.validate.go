// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/overload/v2alpha/overload.proto

package v2alpha

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogo/protobuf/types"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = types.DynamicAny{}
)

// Validate checks the field values on ResourceMonitor with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ResourceMonitor) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetName()) < 1 {
		return ResourceMonitorValidationError{
			field:  "Name",
			reason: "value length must be at least 1 bytes",
		}
	}

	switch m.ConfigType.(type) {

	case *ResourceMonitor_Config:

		if v, ok := interface{}(m.GetConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ResourceMonitorValidationError{
					field:  "Config",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ResourceMonitor_TypedConfig:

		if v, ok := interface{}(m.GetTypedConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ResourceMonitorValidationError{
					field:  "TypedConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ResourceMonitorValidationError is the validation error returned by
// ResourceMonitor.Validate if the designated constraints aren't met.
type ResourceMonitorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResourceMonitorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResourceMonitorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResourceMonitorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResourceMonitorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResourceMonitorValidationError) ErrorName() string { return "ResourceMonitorValidationError" }

// Error satisfies the builtin error interface
func (e ResourceMonitorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResourceMonitor.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResourceMonitorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResourceMonitorValidationError{}

// Validate checks the field values on ThresholdTrigger with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ThresholdTrigger) Validate() error {
	if m == nil {
		return nil
	}

	if val := m.GetValue(); val < 0 || val > 1 {
		return ThresholdTriggerValidationError{
			field:  "Value",
			reason: "value must be inside range [0, 1]",
		}
	}

	return nil
}

// ThresholdTriggerValidationError is the validation error returned by
// ThresholdTrigger.Validate if the designated constraints aren't met.
type ThresholdTriggerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ThresholdTriggerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ThresholdTriggerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ThresholdTriggerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ThresholdTriggerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ThresholdTriggerValidationError) ErrorName() string { return "ThresholdTriggerValidationError" }

// Error satisfies the builtin error interface
func (e ThresholdTriggerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sThresholdTrigger.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ThresholdTriggerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ThresholdTriggerValidationError{}

// Validate checks the field values on Trigger with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Trigger) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetName()) < 1 {
		return TriggerValidationError{
			field:  "Name",
			reason: "value length must be at least 1 bytes",
		}
	}

	switch m.TriggerOneof.(type) {

	case *Trigger_Threshold:

		if v, ok := interface{}(m.GetThreshold()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TriggerValidationError{
					field:  "Threshold",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return TriggerValidationError{
			field:  "TriggerOneof",
			reason: "value is required",
		}

	}

	return nil
}

// TriggerValidationError is the validation error returned by Trigger.Validate
// if the designated constraints aren't met.
type TriggerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TriggerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TriggerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TriggerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TriggerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TriggerValidationError) ErrorName() string { return "TriggerValidationError" }

// Error satisfies the builtin error interface
func (e TriggerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTrigger.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TriggerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TriggerValidationError{}

// Validate checks the field values on OverloadAction with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *OverloadAction) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetName()) < 1 {
		return OverloadActionValidationError{
			field:  "Name",
			reason: "value length must be at least 1 bytes",
		}
	}

	if len(m.GetTriggers()) < 1 {
		return OverloadActionValidationError{
			field:  "Triggers",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetTriggers() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OverloadActionValidationError{
					field:  fmt.Sprintf("Triggers[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// OverloadActionValidationError is the validation error returned by
// OverloadAction.Validate if the designated constraints aren't met.
type OverloadActionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OverloadActionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OverloadActionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OverloadActionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OverloadActionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OverloadActionValidationError) ErrorName() string { return "OverloadActionValidationError" }

// Error satisfies the builtin error interface
func (e OverloadActionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOverloadAction.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OverloadActionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OverloadActionValidationError{}

// Validate checks the field values on OverloadManager with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *OverloadManager) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetRefreshInterval()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OverloadManagerValidationError{
				field:  "RefreshInterval",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(m.GetResourceMonitors()) < 1 {
		return OverloadManagerValidationError{
			field:  "ResourceMonitors",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetResourceMonitors() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OverloadManagerValidationError{
					field:  fmt.Sprintf("ResourceMonitors[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetActions() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OverloadManagerValidationError{
					field:  fmt.Sprintf("Actions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// OverloadManagerValidationError is the validation error returned by
// OverloadManager.Validate if the designated constraints aren't met.
type OverloadManagerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OverloadManagerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OverloadManagerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OverloadManagerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OverloadManagerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OverloadManagerValidationError) ErrorName() string { return "OverloadManagerValidationError" }

// Error satisfies the builtin error interface
func (e OverloadManagerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOverloadManager.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OverloadManagerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OverloadManagerValidationError{}
