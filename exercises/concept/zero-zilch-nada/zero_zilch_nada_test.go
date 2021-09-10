package zero

import (
	"testing"
)

func TestIsZeroBool(t *testing.T) {
	tests := []struct {
		name  string
		value bool
		want  bool
	}{
		{"ZeroBool", false, true},
		{"NonZeroBool", true, false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := IsZeroBool(test.value); got != test.want {
				t.Errorf("IsZeroBool() = %t, want %t", got, test.want)
			}
		})
	}
}

func TestIsZeroInt(t *testing.T) {
	tests := []struct {
		name  string
		value int
		want  bool
	}{
		{"ZeroInt", 0, true},
		{"NonZeroInt", 42, false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := IsZeroInt(test.value); got != test.want {
				t.Errorf("IsZeroInt() = %t, want %t", got, test.want)
			}
		})
	}
}

func TestIsZeroString(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  bool
	}{
		{"ZeroString", "", true},
		{"NonZeroString", "exercism", false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := IsZeroString(test.value); got != test.want {
				t.Errorf("IsZeroString() = %t, want %t", got, test.want)
			}
		})
	}
}

func TestIsZeroPointer(t *testing.T) {
	i := 42

	tests := []struct {
		name  string
		value *int
		want  bool
	}{
		{"ZeroPointer", nil, true},
		{"NonZeroPointer", &i, false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := IsZeroPointer(test.value); got != test.want {
				t.Errorf("IsZeroPointer() = %t, want %t", got, test.want)
			}
		})
	}
}

func TestIsZeroFunc(t *testing.T) {
	tests := []struct {
		name  string
		value func()
		want  bool
	}{
		{"ZeroFunc", nil, true},
		{"NonZeroFunc", func() {}, false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := IsZeroFunc(test.value); got != test.want {
				t.Errorf("IsZeroFunc() = %t, want %t", got, test.want)
			}
		})
	}
}

func TestIsZeroInterface(t *testing.T) {
	tests := []struct {
		name  string
		value interface{}
		want  bool
	}{
		{"ZeroInterface", nil, true},
		{"NonZeroInterface", "exercism", false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := IsZeroInterface(test.value); got != test.want {
				t.Errorf("IsZeroInterface() = %t, want %t", got, test.want)
			}
		})
	}
}

func TestIsZeroSlice(t *testing.T) {
	tests := []struct {
		name  string
		value []int
		want  bool
	}{
		{"ZeroSlice", nil, true},
		{"NonZeroSlice", []int{}, false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := IsZeroSlice(test.value); got != test.want {
				t.Errorf("IsZeroSlice() = %t, want %t", got, test.want)
			}
		})
	}
}

func TestIsZeroChannel(t *testing.T) {
	tests := []struct {
		name  string
		value chan int
		want  bool
	}{
		{"ZeroChannel", nil, true},
		{"NonZeroChannel", make(chan int), false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := IsZeroChannel(test.value); got != test.want {
				t.Errorf("IsZeroChannel() = %t, want %t", got, test.want)
			}
		})
	}
}

func TestIsZeroMap(t *testing.T) {
	tests := []struct {
		name  string
		value map[string]int
		want  bool
	}{
		{"ZeroMap", nil, true},
		{"NonZeroMap", map[string]int{}, false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := IsZeroMap(test.value); got != test.want {
				t.Errorf("IsZeroMap() = %t, want %t", got, test.want)
			}
		})
	}
}
