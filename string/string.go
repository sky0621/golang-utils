package string

import (
	"fmt"
	"strings"
)

// NewBuilder ...
func NewBuilder() Builder[any] {
	return &builder[any]{sb: strings.Builder{}, anyErrors: []error{}}
}

// Builder ...
type Builder[T any] interface {
	Append(v T) Builder[T]
	ToString() string
	HasError() bool
	Errors() []error
	ErrorMsg() string
}

type builder[T any] struct {
	sb               strings.Builder
	anyErrors        []error
	anyErrorMessages []string
}

// Append ...
func (b *builder[T]) Append(v T) Builder[T] {
	var str string

	switch val := any(v).(type) {
	case string:
		str = val
	case int:
		str = fmt.Sprintf("%d", val)
	default:
		str = fmt.Sprintf("%v", val)
	}

	_, err := b.sb.WriteString(str)
	if err != nil {
		b.anyErrors = append(b.anyErrors, err)
		b.anyErrorMessages = append(b.anyErrorMessages, err.Error())
	}
	return b
}

// ToString ...
func (b *builder[T]) ToString() string {
	if len(b.anyErrors) > 0 {
		return ""
	}
	return b.sb.String()
}

// HasError ...
func (b *builder[T]) HasError() bool {
	return b.anyErrors != nil && len(b.anyErrors) > 0
}

// Errors ...
func (b *builder[T]) Errors() []error {
	return b.anyErrors
}

// ErrorMsg ...
func (b *builder[T]) ErrorMsg() string {
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
