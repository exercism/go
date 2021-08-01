package zero

import (
	"reflect"
	"testing"
)

func TestEmptyInterface(t *testing.T) {
	test := struct {
		name string
		want interface{}
	}{
		"EmptyInterface",
		nil,
	}
	t.Run(test.name, func(t *testing.T) {
		if got := EmptyInterface(); got != test.want {
			t.Errorf("EmptyInterface() = %v, want %v", got, test.want)
		}
	})
}

func TestEmptyMap(t *testing.T) {
	test := struct {
		name string
		want map[int]int
	}{
		"EmptyMap",
		nil,
	}
	t.Run(test.name, func(t *testing.T) {
		if got := EmptyMap(); got != nil {
			t.Errorf("EmptyMap() = %v, want %v", got, test.want)
		}
	})
}

func TestEmptySlice(t *testing.T) {
	test := struct {
		name string
		want []int
	}{
		"EmptySlice",
		nil,
	}
	t.Run(test.name, func(t *testing.T) {
		if got := EmptySlice(); got != nil {
			t.Errorf("EmptySlice() = %p, want %p", got, test.want)
		}
	})
}

func TestEmptyString(t *testing.T) {
	test := struct {
		name string
		want string
	}{
		"EmptyString",
		"",
	}
	t.Run(test.name, func(t *testing.T) {
		if got := EmptyString(); got != test.want {
			t.Errorf("EmptyString() = %s, want %s", got, test.want)
		}
	})
}

func TestEmptyChannel(t *testing.T) {
	test := struct {
		name string
		want chan int
	}{
		"EmptyChannel",
		nil,
	}
	t.Run(test.name, func(t *testing.T) {
		if got := EmptyChannel(); got != test.want {
			t.Errorf("EmptyChannel() = %v, want %v", got, test.want)
		}
	})
}

func TestEmptyPointer(t *testing.T) {
	test := struct {
		name string
		want *int
	}{
		"EmptyPointer",
		nil,
	}
	t.Run(test.name, func(t *testing.T) {
		if got := EmptyPointer(); got != test.want {
			t.Errorf("EmptyPointer() = %v, want %v", got, test.want)
		}
	})
}

func TestEmptyBool(t *testing.T) {
	test := struct {
		name string
		want bool
	}{
		"EmptyBool",
		false,
	}
	t.Run(test.name, func(t *testing.T) {
		if got := EmptyBool(); got != test.want {
			t.Errorf("EmptyBool() = %t, want %t", got, test.want)
		}
	})
}

func TestEmptyFunc(t *testing.T) {
	test := struct {
		name string
		want func()
	}{
		"EmptyFunc",
		nil,
	}
	t.Run(test.name, func(t *testing.T) {
		if got := EmptyFunc(); !reflect.DeepEqual(got, test.want) {
			t.Errorf("EmptyFunc() = %p, want nil", got)
		}
	})
}

func TestEmptyInt(t *testing.T) {
	test := struct {
		name string
		want int
	}{
		"EmptyInt",
		0,
	}
	t.Run(test.name, func(t *testing.T) {
		if got := EmptyInt(); got != test.want {
			t.Errorf("EmptyInt() = %d, want %d", got, test.want)
		}
	})
}
