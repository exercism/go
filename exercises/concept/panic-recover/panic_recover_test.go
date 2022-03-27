package go/exercises/concept/panic-recover

import (
	"testing"
)

func AddPanicTest(t *testing.T) {

	tests := []struct {
		name []string
	}{
		{
			name: []string{"firstName", "secondName", "thirdName"},
		},
	}

	for index, tt := range tests {
		t.Run(tt.name[index], func(*testing.T) {
			defer func() {
				if err := recover(); err != "Index Out Of Bounds" {
					t.Errorf("Panic Recovered, Error: %s", err)
				}
			}()
			AddPanic(tt.name, len(tt.name))
		})
	}
}

func RecoverPanicTest(t *testing.T) {
	tests := []struct {
		name []string
	}{
		{
			name: []string{"firstName", "secondName", "thirdName"},
		},
	}
	for _, tt := range tests {
		t.Run("Recover", func(*testing.T) {
			RecoverPanic()
		})
		val := CreatePanic(tt.name, len(tt.name))
	}
}
