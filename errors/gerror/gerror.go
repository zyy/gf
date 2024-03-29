// Copyright GoFrame gf Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package gerror provides simple functions to manipulate errors.
//
// Very note that, this package is quite a basic package, which SHOULD NOT import extra packages
// except standard packages and internal packages, to avoid cycle imports.
package gerror

import (
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/errors/gcode"
)

// New creates and returns an error which is formatted from given text.
func New(text string) error {
	return &Error{
		stack: callers(),
		text:  text,
		code:  gcode.CodeNil,
	}
}

// Newf returns an error that formats as the given format and args.
func Newf(format string, args ...interface{}) error {
	return &Error{
		stack: callers(),
		text:  fmt.Sprintf(format, args...),
		code:  gcode.CodeNil,
	}
}

// NewSkip creates and returns an error which is formatted from given text.
// The parameter `skip` specifies the stack callers skipped amount.
func NewSkip(skip int, text string) error {
	return &Error{
		stack: callers(skip),
		text:  text,
		code:  gcode.CodeNil,
	}
}

// NewSkipf returns an error that formats as the given format and args.
// The parameter `skip` specifies the stack callers skipped amount.
func NewSkipf(skip int, format string, args ...interface{}) error {
	return &Error{
		stack: callers(skip),
		text:  fmt.Sprintf(format, args...),
		code:  gcode.CodeNil,
	}
}

// Wrap wraps error with text. It returns nil if given err is nil.
// Note that it does not lose the error code of wrapped error, as it inherits the error code from it.
func Wrap(err error, text string) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(),
		text:  text,
		code:  Code(err),
	}
}

// Wrapf returns an error annotating err with a stack trace at the point Wrapf is called, and the format specifier.
// It returns nil if given `err` is nil.
// Note that it does not lose the error code of wrapped error, as it inherits the error code from it.
func Wrapf(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(),
		text:  fmt.Sprintf(format, args...),
		code:  Code(err),
	}
}

// WrapSkip wraps error with text. It returns nil if given err is nil.
// The parameter `skip` specifies the stack callers skipped amount.
// Note that it does not lose the error code of wrapped error, as it inherits the error code from it.
func WrapSkip(skip int, err error, text string) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(skip),
		text:  text,
		code:  Code(err),
	}
}

// WrapSkipf wraps error with text that is formatted with given format and args. It returns nil if given err is nil.
// The parameter `skip` specifies the stack callers skipped amount.
// Note that it does not lose the error code of wrapped error, as it inherits the error code from it.
func WrapSkipf(skip int, err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(skip),
		text:  fmt.Sprintf(format, args...),
		code:  Code(err),
	}
}

// NewCode creates and returns an error that has error code and given text.
func NewCode(code gcode.Code, text ...string) error {
	return &Error{
		stack: callers(),
		text:  strings.Join(text, ", "),
		code:  code,
	}
}

// NewCodef returns an error that has error code and formats as the given format and args.
func NewCodef(code gcode.Code, format string, args ...interface{}) error {
	return &Error{
		stack: callers(),
		text:  fmt.Sprintf(format, args...),
		code:  code,
	}
}

// NewCodeSkip creates and returns an error which has error code and is formatted from given text.
// The parameter `skip` specifies the stack callers skipped amount.
func NewCodeSkip(code gcode.Code, skip int, text ...string) error {
	return &Error{
		stack: callers(skip),
		text:  strings.Join(text, ", "),
		code:  code,
	}
}

// NewCodeSkipf returns an error that has error code and formats as the given format and args.
// The parameter `skip` specifies the stack callers skipped amount.
func NewCodeSkipf(code gcode.Code, skip int, format string, args ...interface{}) error {
	return &Error{
		stack: callers(skip),
		text:  fmt.Sprintf(format, args...),
		code:  code,
	}
}

// WrapCode wraps error with code and text.
// It returns nil if given err is nil.
func WrapCode(code gcode.Code, err error, text ...string) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(),
		text:  strings.Join(text, ", "),
		code:  code,
	}
}

// WrapCodef wraps error with code and format specifier.
// It returns nil if given `err` is nil.
func WrapCodef(code gcode.Code, err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(),
		text:  fmt.Sprintf(format, args...),
		code:  code,
	}
}

// WrapCodeSkip wraps error with code and text.
// It returns nil if given err is nil.
// The parameter `skip` specifies the stack callers skipped amount.
func WrapCodeSkip(code gcode.Code, skip int, err error, text ...string) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(skip),
		text:  strings.Join(text, ", "),
		code:  code,
	}
}

// WrapCodeSkipf wraps error with code and text that is formatted with given format and args.
// It returns nil if given err is nil.
// The parameter `skip` specifies the stack callers skipped amount.
func WrapCodeSkipf(code gcode.Code, skip int, err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(skip),
		text:  fmt.Sprintf(format, args...),
		code:  code,
	}
}

// Code returns the error code of current error.
// It returns CodeNil if it has no error code neither it does not implement interface Code.
func Code(err error) gcode.Code {
	if err == nil {
		return gcode.CodeNil
	}
	if e, ok := err.(iCode); ok {
		return e.Code()
	}
	if e, ok := err.(iNext); ok {
		return Code(e.Next())
	}
	if e, ok := err.(iUnwrap); ok {
		return Code(e.Unwrap())
	}
	return gcode.CodeNil
}

// Cause returns the root cause error of `err`.
func Cause(err error) error {
	if err == nil {
		return nil
	}
	if e, ok := err.(iCause); ok {
		return e.Cause()
	}
	if e, ok := err.(iNext); ok {
		return Cause(e.Next())
	}
	if e, ok := err.(iUnwrap); ok {
		return Cause(e.Unwrap())
	}
	return err
}

// Stack returns the stack callers as string.
// It returns the error string directly if the `err` does not support stacks.
func Stack(err error) string {
	if err == nil {
		return ""
	}
	if e, ok := err.(iStack); ok {
		return e.Stack()
	}
	return err.Error()
}

// Current creates and returns the current level error.
// It returns nil if current level error is nil.
func Current(err error) error {
	if err == nil {
		return nil
	}
	if e, ok := err.(iCurrent); ok {
		return e.Current()
	}
	return err
}

// Next returns the next level error.
// It returns nil if current level error or the next level error is nil.
func Next(err error) error {
	if err == nil {
		return nil
	}
	if e, ok := err.(iNext); ok {
		return e.Next()
	}
	return nil
}

// Unwrap is alias of function `Next`.
// It is just for implements for stdlib errors.Unwrap from Go version 1.17.
func Unwrap(err error) error {
	return Next(err)
}

// HasStack checks and returns whether `err` implemented interface `iStack`.
func HasStack(err error) bool {
	_, ok := err.(iStack)
	return ok
}

// Equal reports whether current error `err` equals to error `target`.
// Please note that, in default comparison for `Error`,
// the errors are considered the same if both the `code` and `text` of them are the same.
func Equal(err, target error) bool {
	if err == target {
		return true
	}
	if e, ok := err.(iEqual); ok {
		return e.Equal(target)
	}
	if e, ok := target.(iEqual); ok {
		return e.Equal(err)
	}
	return false
}

// Is reports whether current error `err` has error `target` in its chaining errors.
// It is just for implements for stdlib errors.Unwrap from Go version 1.17.
func Is(err, target error) bool {
	if e, ok := err.(iIs); ok {
		return e.Is(target)
	}
	return false
}
