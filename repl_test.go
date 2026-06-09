package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	type test struct {
		input string
		want []string
	}

	tests := []test{
		{input: " Hello World ", want: []string{"hello", "world"}},
		{input: "MY DUMMY COMMAND", want: []string{"my", "dummy", "command"}},
		{input: "   this is   text", want: []string{"this", "is", "text"}},
	}

	for _, tc := range tests {
		got := cleanInput(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got %v", tc.want, got)
		}
	}
}
