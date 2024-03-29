package main

import "testing"

func Test_hello(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"hello, world", "hello, world"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hello(); got != tt.want {
				t.Errorf("hello() = %v, want %v", got, tt.want)
			}
		})
	}
}
