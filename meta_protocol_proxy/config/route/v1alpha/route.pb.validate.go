// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/meta_protocol_proxy/config/route/v1alpha/route.proto

package aeraki_meta_protocol_proxy_config_route_v1alpha

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

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on RouteConfiguration with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RouteConfiguration) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	for idx, item := range m.GetRoutes() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RouteConfigurationValidationError{
					field:  fmt.Sprintf("Routes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// RouteConfigurationValidationError is the validation error returned by
// RouteConfiguration.Validate if the designated constraints aren't met.
type RouteConfigurationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RouteConfigurationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RouteConfigurationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RouteConfigurationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RouteConfigurationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RouteConfigurationValidationError) ErrorName() string {
	return "RouteConfigurationValidationError"
}

// Error satisfies the builtin error interface
func (e RouteConfigurationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRouteConfiguration.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RouteConfigurationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RouteConfigurationValidationError{}

// Validate checks the field values on Route with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Route) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	if v, ok := interface{}(m.GetMatch()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RouteValidationError{
				field:  "Match",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetRoute() == nil {
		return RouteValidationError{
			field:  "Route",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetRoute()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RouteValidationError{
				field:  "Route",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(m.GetRequestMutation()) > 1000 {
		return RouteValidationError{
			field:  "RequestMutation",
			reason: "value must contain no more than 1000 item(s)",
		}
	}

	for idx, item := range m.GetRequestMutation() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RouteValidationError{
					field:  fmt.Sprintf("RequestMutation[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(m.GetResponseMutation()) > 1000 {
		return RouteValidationError{
			field:  "ResponseMutation",
			reason: "value must contain no more than 1000 item(s)",
		}
	}

	for idx, item := range m.GetResponseMutation() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RouteValidationError{
					field:  fmt.Sprintf("ResponseMutation[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// RouteValidationError is the validation error returned by Route.Validate if
// the designated constraints aren't met.
type RouteValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RouteValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RouteValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RouteValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RouteValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RouteValidationError) ErrorName() string { return "RouteValidationError" }

// Error satisfies the builtin error interface
func (e RouteValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRoute.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RouteValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RouteValidationError{}

// Validate checks the field values on RouteMatch with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *RouteMatch) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetMetadata() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RouteMatchValidationError{
					field:  fmt.Sprintf("Metadata[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// RouteMatchValidationError is the validation error returned by
// RouteMatch.Validate if the designated constraints aren't met.
type RouteMatchValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RouteMatchValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RouteMatchValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RouteMatchValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RouteMatchValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RouteMatchValidationError) ErrorName() string { return "RouteMatchValidationError" }

// Error satisfies the builtin error interface
func (e RouteMatchValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRouteMatch.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RouteMatchValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RouteMatchValidationError{}

// Validate checks the field values on RouteAction with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *RouteAction) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetHashPolicy()) > 100 {
		return RouteActionValidationError{
			field:  "HashPolicy",
			reason: "value must contain no more than 100 item(s)",
		}
	}

	switch m.ClusterSpecifier.(type) {

	case *RouteAction_Cluster:
		// no validation rules for Cluster

	case *RouteAction_WeightedClusters:

		if v, ok := interface{}(m.GetWeightedClusters()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RouteActionValidationError{
					field:  "WeightedClusters",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return RouteActionValidationError{
			field:  "ClusterSpecifier",
			reason: "value is required",
		}

	}

	return nil
}

// RouteActionValidationError is the validation error returned by
// RouteAction.Validate if the designated constraints aren't met.
type RouteActionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RouteActionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RouteActionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RouteActionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RouteActionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RouteActionValidationError) ErrorName() string { return "RouteActionValidationError" }

// Error satisfies the builtin error interface
func (e RouteActionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRouteAction.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RouteActionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RouteActionValidationError{}

// Validate checks the field values on KeyValue with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *KeyValue) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetKey()) < 1 {
		return KeyValueValidationError{
			field:  "Key",
			reason: "value length must be at least 1 runes",
		}
	}

	if len(m.GetKey()) > 16384 {
		return KeyValueValidationError{
			field:  "Key",
			reason: "value length must be at most 16384 bytes",
		}
	}

	if len(m.GetValue()) > 16384 {
		return KeyValueValidationError{
			field:  "Value",
			reason: "value length must be at most 16384 bytes",
		}
	}

	return nil
}

// KeyValueValidationError is the validation error returned by
// KeyValue.Validate if the designated constraints aren't met.
type KeyValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e KeyValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e KeyValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e KeyValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e KeyValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e KeyValueValidationError) ErrorName() string { return "KeyValueValidationError" }

// Error satisfies the builtin error interface
func (e KeyValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sKeyValue.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = KeyValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = KeyValueValidationError{}
