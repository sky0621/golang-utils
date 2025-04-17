package convert

import (
	"fmt"
	"reflect"
	"testing"
)

func TestToPtrAndToVal(t *testing.T) {
	var bVal = true
	var sVal = "str"
	var iVal = 123
	var i8val int8 = -8
	var i16val int16 = -16
	var i32val int32 = -32
	var i64val int64 = -64
	var u8val uint8 = 8
	var u16val uint16 = 16
	var u32val uint32 = 32
	var u64val uint64 = 64
	var byteVal byte = 8
	var runeVal rune
	s := "あ"
	for _, r := range s {
		runeVal = r
	}
	var f32val float32 = 3.2
	var f64val = 6.4

	tests := []struct {
		name string
		val  any
	}{
		{name: "bool <-> *bool", val: bVal},

		{name: "string <-> *string", val: sVal},

		{name: "int <-> *int", val: iVal},
		{name: "int8 <-> *int8", val: i8val},
		{name: "int16 <-> *int16", val: i16val},
		{name: "int32 <-> *int32", val: i32val},
		{name: "int64 <-> *int64", val: i64val},

		{name: "uint8 <-> *uint8", val: u8val},
		{name: "uint16 <-> *uint16", val: u16val},
		{name: "uint32 <-> *uint32", val: u32val},
		{name: "uint64 <-> *uint64", val: u64val},

		{name: "byte <-> *byte", val: byteVal},

		{name: "rune <-> *rune", val: runeVal},

		{name: "float32 <-> *float32", val: f32val},
		{name: "float64 <-> *float64", val: f64val},

		// complex は端折る。
	}
	for idx, tt := range tests {
		t.Run(fmt.Sprintf("[No.%d] %s", idx+1, tt.name), func(t *testing.T) {
			got := ToVal(ToPtr(tt.val))
			if got != tt.val {
				t.Errorf("expected: %v, but got %v", tt.val, got)
			}
		})
	}
}

func TestToMap(t *testing.T) {
	type Person struct {
		ID   int
		Name string
	}

	tests := []struct {
		name     string
		slice    interface{}
		keyFunc  interface{}
		expected interface{}
	}{
		{
			name: "slice of structs to map with ID as key",
			slice: []Person{
				{ID: 1, Name: "Alice"},
				{ID: 2, Name: "Bob"},
				{ID: 3, Name: "Charlie"},
			},
			keyFunc: func(p Person) int { return p.ID },
			expected: map[int]Person{
				1: {ID: 1, Name: "Alice"},
				2: {ID: 2, Name: "Bob"},
				3: {ID: 3, Name: "Charlie"},
			},
		},
		{
			name:     "slice of strings to map with first character as key",
			slice:    []string{"apple", "banana", "cherry"},
			keyFunc:  func(s string) byte { return s[0] },
			expected: map[byte]string{'a': "apple", 'b': "banana", 'c': "cherry"},
		},
		{
			name:     "empty slice",
			slice:    []int{},
			keyFunc:  func(i int) int { return i },
			expected: map[int]int{},
		},
		{
			name:     "slice of ints to map with doubled value as key",
			slice:    []int{1, 2, 3},
			keyFunc:  func(i int) int { return i * 2 },
			expected: map[int]int{2: 1, 4: 2, 6: 3},
		},
	}

	for idx, tt := range tests {
		t.Run(fmt.Sprintf("[No.%d] %s", idx+1, tt.name), func(t *testing.T) {
			switch slice := tt.slice.(type) {
			case []Person:
				keyFunc := tt.keyFunc.(func(Person) int)
				expected := tt.expected.(map[int]Person)
				result := ToMap(slice, keyFunc)
				if !reflect.DeepEqual(result, expected) {
					t.Errorf("expected: %v, but got %v", expected, result)
				}
			case []string:
				keyFunc := tt.keyFunc.(func(string) byte)
				expected := tt.expected.(map[byte]string)
				result := ToMap(slice, keyFunc)
				if !reflect.DeepEqual(result, expected) {
					t.Errorf("expected: %v, but got %v", expected, result)
				}
			case []int:
				keyFunc := tt.keyFunc.(func(int) int)
				expected := tt.expected.(map[int]int)
				result := ToMap(slice, keyFunc)
				if !reflect.DeepEqual(result, expected) {
					t.Errorf("expected: %v, but got %v", expected, result)
				}
			default:
				t.Fatalf("unsupported test case type: %T", tt.slice)
			}
		})
	}
}
