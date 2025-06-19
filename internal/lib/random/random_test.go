package random

import (
	"testing"
)

	// валидатор для проверки случайных строк (true если строка содержит разрешенные символы)
	func isValidRandomString(s string) bool {
    for _, c := range s {
        if !(('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || ('0' <= c && c <= '9')) {
            return false
        }
    }
    return true
}

func TestNewRandomStrings(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{
		name: "Test with size 1",
		size: 1,
		},

		{
			name: "Test with size 5",
			size: 5,
		},

		{
			name: "Test with size 10",
			size: 10,
		},

		{
			name: "Test with size 13",
			size: 13,
		},
		{
			name: "Test with size 18",
			size: 18,
		},
		{
			name: "Test with size 25",
			size: 25,
		},
		{
			name: "Test with size 43",
			size: 43,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewRandomString(tt.size)
			if len(got) != tt.size {
				t.Errorf("NewRandomString() = %v, want %v", len(got), tt.size)
			}
			if !isValidRandomString(got) {
				t.Errorf("NewRandomString() = %v, contains invalid characters", got)
			}
		})
	}
}
