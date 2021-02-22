package main

import "testing"

func TestMain(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "No Panic",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
