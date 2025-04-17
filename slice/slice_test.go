package slice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestContains(t *testing.T) {
	tests := []struct {
		name      string
		checkFunc func() bool
	}{
		{name: "contains string", checkFunc: func() bool { return Contains([]string{"a", "b", "c"}, "b") }},
		{name: "contains not string", checkFunc: func() bool { return !Contains([]string{"a", "b", "c"}, "e") }},
		{name: "contains int", checkFunc: func() bool { return Contains([]int{123, 456, 789}, 789) }},
		{name: "contains not int", checkFunc: func() bool { return !Contains([]int{123, 456, 789}, 777) }},
	}
	for idx, tt := range tests {
		t.Run(fmt.Sprintf("[No.%d] %s", idx+1, tt.name), func(t *testing.T) {
			if !tt.checkFunc() {
				t.Fail()
			}
		})
	}
}

func TestExtract(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	// Define Person test data once to be reused
	people := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
	}

	tests := []struct {
		name     string
		input    interface{}
		selector interface{}
		expected interface{}
	}{
		{
			name:     "extract string from string slice",
			input:    []string{"apple", "banana", "cherry"},
			selector: func(s string) string { return s + "!" },
			expected: []string{"apple!", "banana!", "cherry!"},
		},
		{
			name:     "extract int from int slice",
			input:    []int{1, 2, 3},
			selector: func(i int) int { return i * 2 },
			expected: []int{2, 4, 6},
		},
		{
			name:     "extract field from struct slice",
			input:    people,
			selector: func(p Person) string { return p.Name },
			expected: []string{"Alice", "Bob", "Charlie"},
		},
		{
			name:     "extract and transform field from struct slice",
			input:    people,
			selector: func(p Person) int { return p.Age + 10 },
			expected: []int{40, 35, 45},
		},
	}

	for idx, tt := range tests {
		t.Run(fmt.Sprintf("[No.%d] %s", idx+1, tt.name), func(t *testing.T) {
			switch input := tt.input.(type) {
			case []string:
				selector := tt.selector.(func(string) string)
				result := Extract(input, selector)
				expected := tt.expected.([]string)
				if !reflect.DeepEqual(result, expected) {
					t.Errorf("Expected %v, got %v", expected, result)
				}
			case []int:
				selector := tt.selector.(func(int) int)
				result := Extract(input, selector)
				expected := tt.expected.([]int)
				if !reflect.DeepEqual(result, expected) {
					t.Errorf("Expected %v, got %v", expected, result)
				}
			case []Person:
				switch selector := tt.selector.(type) {
				case func(Person) string:
					result := Extract(input, selector)
					expected := tt.expected.([]string)
					if !reflect.DeepEqual(result, expected) {
						t.Errorf("Expected %v, got %v", expected, result)
					}
				case func(Person) int:
					result := Extract(input, selector)
					expected := tt.expected.([]int)
					if !reflect.DeepEqual(result, expected) {
						t.Errorf("Expected %v, got %v", expected, result)
					}
				}
			}
		})
	}
}
