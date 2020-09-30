package stringutil

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestStringBuilderAppend(t *testing.T) {
	const caseName1 = "指定した文字列が結合されること"
	tests := []struct {
		name  string
		input []string
		want  string
	}{
		{name: caseName1, input: []string{}, want: ""},
		{name: caseName1, input: []string{"test"}, want: "test"},
		{name: caseName1, input: []string{"ABC", "DEF", "GHI"}, want: "ABCDEFGHI"},
	}
	for idx, tt := range tests {
		t.Run(fmt.Sprintf("[No.%d] %s", idx+1, tt.name), func(t *testing.T) {
			builder := NewStringBuilder()
			for _, in := range tt.input {
				builder.Append(in)
			}
			got := builder.ToString()
			if tt.want != got {
				t.Errorf("Append[%v] = %v, want %v", strings.Join(tt.input, ","), got, tt.want)
			}
		})
	}
}

func TestStringBuilderAppendInt(t *testing.T) {
	const caseName1 = "指定した数値が結合されること"
	tests := []struct {
		name  string
		input []int
		want  string
	}{
		{name: caseName1, input: []int{}, want: ""},
		{name: caseName1, input: []int{1}, want: "1"},
		{name: caseName1, input: []int{123, 456, 789}, want: "123456789"},
	}
	for idx, tt := range tests {
		t.Run(fmt.Sprintf("[No.%d] %s", idx+1, tt.name), func(t *testing.T) {
			builder := NewStringBuilder()
			for _, in := range tt.input {
				builder.AppendInt(in)
			}
			got := builder.ToString()
			if tt.want != got {
				var sa []string
				for _, in := range tt.input {
					sa = append(sa, strconv.Itoa(in))
				}
				t.Errorf("Append[%v] = %v, want %v", strings.Join(sa, ","), got, tt.want)
			}
		})
	}
}