// +build integration

package cmd

import (
	"testing"
)

func Test_check(t *testing.T) {

	tests := []struct {
		name string
		args string
		want bool
	}{
		{"zero length string", "", true},
		{"1 char", "1", true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := check(tt.args); got != tt.want {
				t.Errorf("check() = %v, want %v", got, tt.want)
			}
		})
	}
}
