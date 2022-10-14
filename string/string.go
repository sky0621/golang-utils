package string

import (
	"fmt"
	"strings"
)

// NewBuilder ...
func NewBuilder() Builder {
	return &builder{sb: strings.Builder{}, anyErrors: []error{}}
}

// Builder ...
type Builder interface {
	Append(str string) Builder
	AppendInt(num int) Builder
	ToString() string
	HasError() bool
	Errors() []error
	ErrorMsg() string
}

type builder struct {
	sb               strings.Builder
	anyErrors        []error
	anyErrorMessages []string
}

// Append ...
func (b *builder) Append(str string) Builder {
	_, err := b.sb.WriteString(str)
	if err != nil {
		b.anyErrors = append(b.anyErrors, err)
		b.anyErrorMessages = append(b.anyErrorMessages, err.Error())
	}
	return b
}

// AppendInt ...
func (b *builder) AppendInt(num int) Builder {
	_, err := b.sb.WriteString(fmt.Sprintf("%d", num))
	if err != nil {
		b.anyErrors = append(b.anyErrors, err)
		b.anyErrorMessages = append(b.anyErrorMessages, err.Error())
	}
	return b
}

// ToString ...
func (b *builder) ToString() string {
	if len(b.anyErrors) > 0 {
		return ""
	}
	return b.sb.String()
}

// HasError ...
func (b *builder) HasError() bool {
	return b.anyErrors != nil && len(b.anyErrors) > 0
}

// Errors ...
func (b *builder) Errors() []error {
	return b.anyErrors
}

// ErrorMsg ...
func (b *builder) ErrorMsg() string {
	return strings.Join(b.anyErrorMessages, "\n")
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
