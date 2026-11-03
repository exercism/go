package genericset

import (
	"maps"
	"slices"
	"testing"
)

// testRunnerTaskID=1
func TestNewMembers(t *testing.T) {
	if got := NewSet[string](); got == nil {
		t.Errorf("NewSet[string]() = nil, want a Set.")
	}
}

// testRunnerTaskID=1
func TestNewAges(t *testing.T) {
	if got := NewSet[int](); got == nil {
		t.Errorf("NewSet[int]() = nil, want a Set.")
	}
}

// testRunnerTaskID=1
func TestAddMembers(t *testing.T) {
	members := NewSet[string]()
	want := map[string]struct{}{}
	log := []string{"Alice", "Bob", "Alice", "Cara", "Bob"}
	for i := range log {
		want[log[i]] = struct{}{}
		members.Add(log[i])
		if !maps.Equal(members, want) {
			t.Fatalf("Added %v. Got %v but expected %v.", log[0:i], members, want)
		}
	}
}

// testRunnerTaskID=1
func TestAddAges(t *testing.T) {
	ages := NewSet[int]()
	want := map[int]struct{}{}
	log := []int{5, 5, 5, 6, 4, 3, 4, 3, 2}
	for i := range log {
		want[log[i]] = struct{}{}
		ages.Add(log[i])
		if !maps.Equal(ages, want) {
			t.Fatalf("Added %v. Got %v but expected %v.", log[0:i], ages, want)
		}
	}
}

// testRunnerTaskID=2
func TestContainsMember(t *testing.T) {
	members := NewSet[string]()
	log := []string{"Alice", "Bob", "Alice", "Cara", "Bob"}
	for i := range log {
		members.Add(log[i])
		for _, member := range []string{"Alice", "Bob", "Cara", "Dave", "Ellie"} {
			want := slices.Contains(log[0:i+1], member)
			if got := members.Contains(member); got != want {
				t.Fatalf("Added %v. members.Contains(%#v) = %t but want %t.", log[0:i], member, got, want)
			}
		}
	}
}

// testRunnerTaskID=3
func TestEquals(t *testing.T) {
	cases := []struct {
		one  []string
		two  []string
		want bool
	}{
		{
			[]string{},
			[]string{},
			true,
		},
		{
			[]string{"Alice", "Bob"},
			[]string{"Bob", "Alice"},
			true,
		},
		{
			[]string{"Alice"},
			[]string{"Bob", "Alice"},
			false,
		},
		{
			[]string{"Alice", "Bob"},
			[]string{"Alice"},
			false,
		},
	}
	for _, tc := range cases {
		lastMonth := NewSet[string]()
		for _, member := range tc.one {
			lastMonth.Add(member)
		}
		thisMonth := NewSet[string]()
		for _, member := range tc.two {
			thisMonth.Add(member)
		}
		if got := lastMonth.Equals(thisMonth); got != tc.want {
			t.Fatalf("Last month: %v\nThis month: %v\nlastMonth.Equals(thisMonth) = %t, want %t.", tc.one, tc.two, got, tc.want)
		}
	}
}
