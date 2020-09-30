package stringutil

import (
	"fmt"
	"strings"
)

// NewStringBuilder ...
func NewStringBuilder() StringBuilder {
	return &stringBuilder{sb: strings.Builder{}, anyErrors: []error{}}
}

// StringBuilder ...
type StringBuilder interface {
	Append(str string) StringBuilder
	AppendInt(num int) StringBuilder
	ToString() string
	HasError() bool
	Errors() []error
	ErrorMsg() string
}

type stringBuilder struct {
	sb           strings.Builder
	anyErrors    []error
	anyErrorMsgs []string
}

// Append ...
func (b *stringBuilder) Append(str string) StringBuilder {
	_, err := b.sb.WriteString(str)
	if err != nil {
		b.anyErrors = append(b.anyErrors, err)
		b.anyErrorMsgs = append(b.anyErrorMsgs, err.Error())
	}
	return b
}

// AppendInt ...
func (b *stringBuilder) AppendInt(num int) StringBuilder {
	_, err := b.sb.WriteString(fmt.Sprintf("%d", num))
	if err != nil {
		b.anyErrors = append(b.anyErrors, err)
		b.anyErrorMsgs = append(b.anyErrorMsgs, err.Error())
	}
	return b
}

// ToString ...
func (b *stringBuilder) ToString() string {
	if len(b.anyErrors) > 0 {
		return ""
	}
	return b.sb.String()
}

// HasError ...
func (b *stringBuilder) HasError() bool {
	return b.anyErrors != nil && len(b.anyErrors) > 0
}

// Errors ...
func (b *stringBuilder) Errors() []error {
	return b.anyErrors
}

// ErrorMsg ...
func (b *stringBuilder) ErrorMsg() string {
	return strings.Join(b.anyErrorMsgs, "\n")
}

// IsBlank ...
func IsBlank(s *string) bool {
	if s == nil {
		return true
	}
	return *s == ""
}

// IsNotBlank ...
func IsNotBlank(s *string) bool {
	return !IsBlank(s)
}
