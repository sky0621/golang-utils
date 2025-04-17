package string

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestStringBuilderAppend(t *testing.T) {
	type testCase struct {
		name  string
		input any
		want  string
	}

	tests := []testCase{
		// String test cases
		{name: "指定した文字列が結合されること", input: []string{}, want: ""},
		{name: "指定した文字列が結合されること", input: []string{"test"}, want: "test"},
		{name: "指定した文字列が結合されること", input: []string{"ABC", "DEF", "GHI"}, want: "ABCDEFGHI"},

		// Integer test cases
		{name: "指定した数値が結合されること", input: []int{}, want: ""},
		{name: "指定した数値が結合されること", input: []int{1}, want: "1"},
		{name: "指定した数値が結合されること", input: []int{123, 456, 789}, want: "123456789"},
	}

	for idx, tt := range tests {
		t.Run(fmt.Sprintf("[No.%d] %s", idx+1, tt.name), func(t *testing.T) {
			builder := NewBuilder()

			switch input := tt.input.(type) {
			case []string:
				for _, in := range input {
					builder.Append(in)
				}
			case []int:
				for _, in := range input {
					builder.Append(in)
				}
			}

			got := builder.ToString()
			if tt.want != got {
				var inputStr string

				switch input := tt.input.(type) {
				case []string:
					inputStr = strings.Join(input, ",")
				case []int:
					var sa []string
					for _, in := range input {
						sa = append(sa, strconv.Itoa(in))
					}
					inputStr = strings.Join(sa, ",")
				}

				t.Errorf("Append[%v] = %v, want %v", inputStr, got, tt.want)
			}
		})
	}
}
