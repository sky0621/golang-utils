package slice

import (
	"fmt"
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
