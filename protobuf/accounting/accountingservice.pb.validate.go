// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: accountingservice.proto

package accounting

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

// Validate checks the field values on GetReportRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetReportRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetReportRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetReportRequestMultiError, or nil if none found.
func (m *GetReportRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetReportRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	if all {
		switch v := interface{}(m.GetStartTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetReportRequestValidationError{
					field:  "StartTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetReportRequestValidationError{
					field:  "StartTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStartTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetReportRequestValidationError{
				field:  "StartTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetEndTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetReportRequestValidationError{
					field:  "EndTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetReportRequestValidationError{
					field:  "EndTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEndTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetReportRequestValidationError{
				field:  "EndTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetReportRequestMultiError(errors)
	}

	return nil
}

// GetReportRequestMultiError is an error wrapping multiple validation errors
// returned by GetReportRequest.ValidateAll() if the designated constraints
// aren't met.
type GetReportRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetReportRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetReportRequestMultiError) AllErrors() []error { return m }

// GetReportRequestValidationError is the validation error returned by
// GetReportRequest.Validate if the designated constraints aren't met.
type GetReportRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetReportRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetReportRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetReportRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetReportRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetReportRequestValidationError) ErrorName() string { return "GetReportRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetReportRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetReportRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetReportRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetReportRequestValidationError{}

// Validate checks the field values on GetUserResponseElement with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetUserResponseElement) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUserResponseElement with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetUserResponseElementMultiError, or nil if none found.
func (m *GetUserResponseElement) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUserResponseElement) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Type

	// no validation rules for Count

	if all {
		switch v := interface{}(m.GetTimestamp()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetUserResponseElementValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetUserResponseElementValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimestamp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetUserResponseElementValidationError{
				field:  "Timestamp",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for UserId

	if len(errors) > 0 {
		return GetUserResponseElementMultiError(errors)
	}

	return nil
}

// GetUserResponseElementMultiError is an error wrapping multiple validation
// errors returned by GetUserResponseElement.ValidateAll() if the designated
// constraints aren't met.
type GetUserResponseElementMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUserResponseElementMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUserResponseElementMultiError) AllErrors() []error { return m }

// GetUserResponseElementValidationError is the validation error returned by
// GetUserResponseElement.Validate if the designated constraints aren't met.
type GetUserResponseElementValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserResponseElementValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserResponseElementValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserResponseElementValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserResponseElementValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserResponseElementValidationError) ErrorName() string {
	return "GetUserResponseElementValidationError"
}

// Error satisfies the builtin error interface
func (e GetUserResponseElementValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserResponseElement.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserResponseElementValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserResponseElementValidationError{}

// Validate checks the field values on GetReportResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetReportResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetReportResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetReportResponseMultiError, or nil if none found.
func (m *GetReportResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetReportResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetReport() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetReportResponseValidationError{
						field:  fmt.Sprintf("Report[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetReportResponseValidationError{
						field:  fmt.Sprintf("Report[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetReportResponseValidationError{
					field:  fmt.Sprintf("Report[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetReportResponseMultiError(errors)
	}

	return nil
}

// GetReportResponseMultiError is an error wrapping multiple validation errors
// returned by GetReportResponse.ValidateAll() if the designated constraints
// aren't met.
type GetReportResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetReportResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetReportResponseMultiError) AllErrors() []error { return m }

// GetReportResponseValidationError is the validation error returned by
// GetReportResponse.Validate if the designated constraints aren't met.
type GetReportResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetReportResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetReportResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetReportResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetReportResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetReportResponseValidationError) ErrorName() string {
	return "GetReportResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetReportResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetReportResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetReportResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetReportResponseValidationError{}
