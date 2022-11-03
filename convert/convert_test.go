package convert

import (
	"fmt"
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
