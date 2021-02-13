package counter

// NOTE: This file is called example.go and not example_test.go because if it's
// called example_test.go it will get picked up by the test suite and fail
// because COUNTER_IMPL isn't set (see maker.go).

import "testing"

func TestNoAdd(t *testing.T) {
	counter := makeCounter()
	if counter.Lines() != 0 {
		t.Errorf("Lines mismatch: got %d, expected %d", counter.Lines(), 0)
	}
	if counter.Letters() != 0 {
		t.Errorf("Letters mismatch: got %d, expected %d", counter.Letters(), 0)
	}
	if counter.Characters() != 0 {
		t.Errorf("Characters mismatch: got %d, expected %d", counter.Characters(), 0)
	}
}

func TestEmptyString(t *testing.T) {
	counter := makeCounter()
	counter.AddString("")
	if counter.Lines() != 0 {
		t.Errorf("Lines mismatch: got %d, expected %d", counter.Lines(), 0)
	}
	if counter.Letters() != 0 {
		t.Errorf("Letters mismatch: got %d, expected %d", counter.Letters(), 0)
	}
	if counter.Characters() != 0 {
		t.Errorf("Characters mismatch: got %d, expected %d", counter.Characters(), 0)
	}
}

func TestASCIIString(t *testing.T) {
	counter := makeCounter()
	counter.AddString("Hello\nworld!")
	if counter.Lines() != 2 {
		t.Errorf("Lines mismatch: got %d, expected %d", counter.Lines(), 2)
	}
	if counter.Letters() != 10 {
		t.Errorf("Letters mismatch: got %d, expected %d", counter.Letters(), 10)
	}
	if counter.Characters() != 12 {
		t.Errorf("Characters mismatch: got %d, expected %d", counter.Characters(), 12)
	}
}

func TestRussianString(t *testing.T) {
	counter := makeCounter()
	// Lifted this translation from the ru.po file of GNU hello
	counter.AddString("здравствуй, мир\n")
	if counter.Lines() != 1 {
		t.Errorf("Lines mismatch: got %d, expected %d", counter.Lines(), 2)
	}
	if counter.Letters() != 13 {
		t.Errorf("Letters mismatch: got %d, expected %d", counter.Letters(), 10)
	}
	if counter.Characters() != 16 {
		t.Errorf("Characters mismatch: got %d, expected %d", counter.Characters(), 12)
	}
}
