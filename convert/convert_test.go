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

	tests := map[string]func() bool{
		"bool <-> *bool": func() bool { return ToVal(ToPtr(bVal)) != bVal },

		"string <-> *string": func() bool { return ToVal(ToPtr(sVal)) != sVal },

		"int <-> *int":     func() bool { return ToVal(ToPtr(iVal)) != iVal },
		"int8 <-> *int8":   func() bool { return ToVal(ToPtr(i8val)) != i8val },
		"int16 <-> *int16": func() bool { return ToVal(ToPtr(i16val)) != i16val },
		"int32 <-> *int32": func() bool { return ToVal(ToPtr(i32val)) != i32val },
		"int64 <-> *int64": func() bool { return ToVal(ToPtr(i64val)) != i64val },

		"uint8 <-> *uint8":   func() bool { return ToVal(ToPtr(u8val)) != u8val },
		"uint16 <-> *uint16": func() bool { return ToVal(ToPtr(u16val)) != u16val },
		"uint32 <-> *uint32": func() bool { return ToVal(ToPtr(u32val)) != u32val },
		"uint64 <-> *uint64": func() bool { return ToVal(ToPtr(u64val)) != u64val },

		"byte <-> *byte": func() bool { return ToVal(ToPtr(byteVal)) != byteVal },

		"rune <-> *rune": func() bool { return ToVal(ToPtr(runeVal)) != runeVal },

		"float32 <-> *float32": func() bool { return ToVal(ToPtr(f32val)) != f32val },
		"float64 <-> *float64": func() bool { return ToVal(ToPtr(f64val)) != f64val },

		// complex は端折る。
	}
	for name, test := range tests {
		t.Run(fmt.Sprintf(name), func(t *testing.T) {
			t.Parallel()
			test()
		})
	}
}
