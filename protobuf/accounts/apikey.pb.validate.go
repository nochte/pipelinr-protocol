// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: apikey.proto

package accounts

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on APIKey with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *APIKey) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on APIKey with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in APIKeyMultiError, or nil if none found.
func (m *APIKey) ValidateAll() error {
	return m.validate(true)
}

func (m *APIKey) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetId()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, APIKeyValidationError{
					field:  "Id",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, APIKeyValidationError{
					field:  "Id",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return APIKeyValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for UserId

	// no validation rules for Key

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, APIKeyValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, APIKeyValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return APIKeyValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, APIKeyValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, APIKeyValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return APIKeyValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ActivationState

	// no validation rules for Label

	if all {
		switch v := interface{}(m.GetLastUsed()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, APIKeyValidationError{
					field:  "LastUsed",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, APIKeyValidationError{
					field:  "LastUsed",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLastUsed()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return APIKeyValidationError{
				field:  "LastUsed",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return APIKeyMultiError(errors)
	}

	return nil
}

// APIKeyMultiError is an error wrapping multiple validation errors returned by
// APIKey.ValidateAll() if the designated constraints aren't met.
type APIKeyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m APIKeyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m APIKeyMultiError) AllErrors() []error { return m }

// APIKeyValidationError is the validation error returned by APIKey.Validate if
// the designated constraints aren't met.
type APIKeyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e APIKeyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e APIKeyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e APIKeyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e APIKeyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e APIKeyValidationError) ErrorName() string { return "APIKeyValidationError" }

// Error satisfies the builtin error interface
func (e APIKeyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAPIKey.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = APIKeyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = APIKeyValidationError{}

// Validate checks the field values on CreateAPIKeyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateAPIKeyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateAPIKeyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateAPIKeyRequestMultiError, or nil if none found.
func (m *CreateAPIKeyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateAPIKeyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	// no validation rules for Label

	if len(errors) > 0 {
		return CreateAPIKeyRequestMultiError(errors)
	}

	return nil
}

// CreateAPIKeyRequestMultiError is an error wrapping multiple validation
// errors returned by CreateAPIKeyRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateAPIKeyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateAPIKeyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateAPIKeyRequestMultiError) AllErrors() []error { return m }

// CreateAPIKeyRequestValidationError is the validation error returned by
// CreateAPIKeyRequest.Validate if the designated constraints aren't met.
type CreateAPIKeyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateAPIKeyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateAPIKeyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateAPIKeyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateAPIKeyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateAPIKeyRequestValidationError) ErrorName() string {
	return "CreateAPIKeyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateAPIKeyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateAPIKeyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateAPIKeyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateAPIKeyRequestValidationError{}

// Validate checks the field values on CreateAPIKeyResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateAPIKeyResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateAPIKeyResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateAPIKeyResponseMultiError, or nil if none found.
func (m *CreateAPIKeyResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateAPIKeyResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for APIKey

	// no validation rules for Message

	if len(errors) > 0 {
		return CreateAPIKeyResponseMultiError(errors)
	}

	return nil
}

// CreateAPIKeyResponseMultiError is an error wrapping multiple validation
// errors returned by CreateAPIKeyResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateAPIKeyResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateAPIKeyResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateAPIKeyResponseMultiError) AllErrors() []error { return m }

// CreateAPIKeyResponseValidationError is the validation error returned by
// CreateAPIKeyResponse.Validate if the designated constraints aren't met.
type CreateAPIKeyResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateAPIKeyResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateAPIKeyResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateAPIKeyResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateAPIKeyResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateAPIKeyResponseValidationError) ErrorName() string {
	return "CreateAPIKeyResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateAPIKeyResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateAPIKeyResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateAPIKeyResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateAPIKeyResponseValidationError{}

// Validate checks the field values on ListAPIKeyResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListAPIKeyResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListAPIKeyResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListAPIKeyResponseMultiError, or nil if none found.
func (m *ListAPIKeyResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListAPIKeyResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetApiKeys() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListAPIKeyResponseValidationError{
						field:  fmt.Sprintf("ApiKeys[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListAPIKeyResponseValidationError{
						field:  fmt.Sprintf("ApiKeys[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListAPIKeyResponseValidationError{
					field:  fmt.Sprintf("ApiKeys[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListAPIKeyResponseMultiError(errors)
	}

	return nil
}

// ListAPIKeyResponseMultiError is an error wrapping multiple validation errors
// returned by ListAPIKeyResponse.ValidateAll() if the designated constraints
// aren't met.
type ListAPIKeyResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListAPIKeyResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListAPIKeyResponseMultiError) AllErrors() []error { return m }

// ListAPIKeyResponseValidationError is the validation error returned by
// ListAPIKeyResponse.Validate if the designated constraints aren't met.
type ListAPIKeyResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListAPIKeyResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListAPIKeyResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListAPIKeyResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListAPIKeyResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListAPIKeyResponseValidationError) ErrorName() string {
	return "ListAPIKeyResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListAPIKeyResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListAPIKeyResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListAPIKeyResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListAPIKeyResponseValidationError{}
