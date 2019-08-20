// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttipb

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

// ValidateFields checks the field values on
// GetTenantIdentifiersForEndDeviceEUIsRequest with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *GetTenantIdentifiersForEndDeviceEUIsRequest) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = GetTenantIdentifiersForEndDeviceEUIsRequestFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "join_eui":
			// no validation rules for JoinEUI
		case "dev_eui":
			// no validation rules for DevEUI
		default:
			return GetTenantIdentifiersForEndDeviceEUIsRequestValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// GetTenantIdentifiersForEndDeviceEUIsRequestValidationError is the validation
// error returned by
// GetTenantIdentifiersForEndDeviceEUIsRequest.ValidateFields if the
// designated constraints aren't met.
type GetTenantIdentifiersForEndDeviceEUIsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTenantIdentifiersForEndDeviceEUIsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTenantIdentifiersForEndDeviceEUIsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTenantIdentifiersForEndDeviceEUIsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTenantIdentifiersForEndDeviceEUIsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTenantIdentifiersForEndDeviceEUIsRequestValidationError) ErrorName() string {
	return "GetTenantIdentifiersForEndDeviceEUIsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetTenantIdentifiersForEndDeviceEUIsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTenantIdentifiersForEndDeviceEUIsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTenantIdentifiersForEndDeviceEUIsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTenantIdentifiersForEndDeviceEUIsRequestValidationError{}

// ValidateFields checks the field values on
// GetTenantIdentifiersForGatewayEUIRequest with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *GetTenantIdentifiersForGatewayEUIRequest) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = GetTenantIdentifiersForGatewayEUIRequestFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "eui":
			// no validation rules for EUI
		default:
			return GetTenantIdentifiersForGatewayEUIRequestValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// GetTenantIdentifiersForGatewayEUIRequestValidationError is the validation
// error returned by GetTenantIdentifiersForGatewayEUIRequest.ValidateFields
// if the designated constraints aren't met.
type GetTenantIdentifiersForGatewayEUIRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTenantIdentifiersForGatewayEUIRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTenantIdentifiersForGatewayEUIRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTenantIdentifiersForGatewayEUIRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTenantIdentifiersForGatewayEUIRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTenantIdentifiersForGatewayEUIRequestValidationError) ErrorName() string {
	return "GetTenantIdentifiersForGatewayEUIRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetTenantIdentifiersForGatewayEUIRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTenantIdentifiersForGatewayEUIRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTenantIdentifiersForGatewayEUIRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTenantIdentifiersForGatewayEUIRequestValidationError{}
