package zero

import (
	"reflect"
	"testing"
)

func TestZeroInterface(t *testing.T) {
	test := struct {
		name string
		want interface{}
	}{
		"ZeroInterface",
		nil,
	}
	t.Run(test.name, func(t *testing.T) {
		if got := ZeroInterface(); got != test.want {
			t.Errorf("ZeroInterface() = %v, want %v", got, test.want)
		}
	})
}

func TestZeroMap(t *testing.T) {
	test := struct {
		name string
		want map[int]int
	}{
		"ZeroMap",
		nil,
	}
	t.Run(test.name, func(t *testing.T) {
		if got := ZeroMap(); got != nil {
			t.Errorf("ZeroMap() = %v, want %v", got, test.want)
		}
	})
}

func TestZeroSlice(t *testing.T) {
	test := struct {
		name string
		want []int
	}{
		"ZeroSlice",
		nil,
	}
	t.Run(test.name, func(t *testing.T) {
		if got := ZeroSlice(); got != nil {
			t.Errorf("ZeroSlice() = %p, want %p", got, test.want)
		}
	})
}

func TestZeroString(t *testing.T) {
	test := struct {
		name string
		want string
	}{
		"ZeroString",
		"",
	}
	t.Run(test.name, func(t *testing.T) {
		if got := ZeroString(); got != test.want {
			t.Errorf("ZeroString() = %s, want %s", got, test.want)
		}
	})
}

func TestZeroChannel(t *testing.T) {
	test := struct {
		name string
		want chan int
	}{
		"ZeroChannel",
		nil,
	}
	t.Run(test.name, func(t *testing.T) {
		if got := ZeroChannel(); got != test.want {
			t.Errorf("ZeroChannel() = %v, want %v", got, test.want)
		}
	})
}

func TestZeroPointer(t *testing.T) {
	test := struct {
		name string
		want *int
	}{
		"ZeroPointer",
		nil,
	}
	t.Run(test.name, func(t *testing.T) {
		if got := ZeroPointer(); got != test.want {
			t.Errorf("ZeroPointer() = %v, want %v", got, test.want)
		}
	})
}

func TestZeroBool(t *testing.T) {
	test := struct {
		name string
		want bool
	}{
		"ZeroBool",
		false,
	}
	t.Run(test.name, func(t *testing.T) {
		if got := ZeroBool(); got != test.want {
			t.Errorf("ZeroBool() = %t, want %t", got, test.want)
		}
	})
}

func TestZeroFunc(t *testing.T) {
	test := struct {
		name string
		want func()
	}{
		"ZeroFunc",
		nil,
	}
	t.Run(test.name, func(t *testing.T) {
		if got := ZeroFunc(); !reflect.DeepEqual(got, test.want) {
			t.Errorf("ZeroFunc() = %p, want nil", got)
		}
	})
}

func TestZeroInt(t *testing.T) {
	test := struct {
		name string
		want int
	}{
		"ZeroInt",
		0,
	}
	t.Run(test.name, func(t *testing.T) {
		if got := ZeroInt(); got != test.want {
			t.Errorf("ZeroInt() = %d, want %d", got, test.want)
		}
	})
}
