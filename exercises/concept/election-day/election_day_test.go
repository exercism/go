package electionday

import (
	"strconv"
	"testing"
)

func TestNewVoteCounter(t *testing.T) {
	tests := []struct {
		name  string
		votes int
	}{
		{
			name:  "Simple vote counter with 2 votes",
			votes: 2,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := NewVoteCounter(tc.votes)
			if got == nil {
				t.Errorf("NewVoteCounter(%d) = nil, want &%d", tc.votes, tc.votes)
			} else if *got != tc.votes {
				t.Errorf("NewVoteCounter(%d) = &%d, want &%d", tc.votes, *got, tc.votes)
			}
		})
	}
}

func TestVoteCount(t *testing.T) {
	twoVotes := 2

	tests := []struct {
		name     string
		counter  *int
		expected int
	}{
		{
			name:     "Call to VoteCount with a nil argument",
			counter:  nil,
			expected: 0,
		},
		{
			name:     "Call to VoteCount with a pointer to an int with a value of 2",
			counter:  &twoVotes,
			expected: 2,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := VoteCount(tc.counter); got != tc.expected {
				t.Fatalf("VoteCount(%s) = %d, want %d", intPtrRepresentation(tc.counter), got, tc.expected)
			}
		})
	}
}

func TestIncrementVoteCount(t *testing.T) {
	twoVotes := 2
	fiveVotes := 5
	noVotes := 0

	tests := []struct {
		name      string
		counter   *int
		increment int
		expected  int
	}{
		{
			name:      "Call to IncrementVoteCount with a pointer to an int with a value of 0 and increment of 1",
			counter:   &noVotes,
			increment: 1,
			expected:  1,
		},
		{
			name:      "Call to IncrementVoteCount with a pointer to an int with a value of 2 and increment of 2",
			counter:   &twoVotes,
			increment: 2,
			expected:  4,
		},
		{
			name:      "Call to IncrementVoteCount with a pointer to an int with a value of 5 and increment of 7",
			counter:   &fiveVotes,
			increment: 7,
			expected:  12,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			before := intPtrRepresentation(tc.counter)
			IncrementVoteCount(tc.counter, tc.increment)
			after := intPtrRepresentation(tc.counter)

			if tc.counter == nil {
				t.Errorf("counter before: %s | counter after: %v | want: &%d", before, after, tc.expected)
			} else if *tc.counter != tc.expected {
				t.Errorf("counter before: %s | counter after: %v | want: &%d", before, after, tc.expected)
			}
		})
	}
}

func TestNewElectionResult(t *testing.T) {
	tests := []struct {
		name          string
		candidateName string
		votes         int
		want          ElectionResult
	}{
		{
			name:          "Call to NewElectionResult for Peter with 2 votes",
			candidateName: "Peter",
			votes:         2,
			want: ElectionResult{
				Name:  "Peter",
				Votes: 2,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := NewElectionResult(tc.candidateName, tc.votes)

			if result == nil || result.Name != tc.want.Name || result.Votes != tc.want.Votes {
				t.Errorf("NewElectionResult(%q, %d) = %#v, want %#v",
					tc.candidateName, tc.votes, result, tc.want)
			}
		})
	}
}

func TestDisplayResult(t *testing.T) {
	tests := []struct {
		name   string
		result *ElectionResult
		want   string
	}{
		{
			name: "Call to DisplayResult for John with 5 votes",
			result: &ElectionResult{
				Name:  "John",
				Votes: 5,
			},
			want: "John (5)",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if result := DisplayResult(tc.result); result != tc.want {
				t.Errorf("DisplayResult(%#v) = %q, want %q", *tc.result, result, tc.want)
			}
		})
	}
}

func TestDecrementVotesOfCandidate(t *testing.T) {
	tests := []struct {
		name      string
		candidate string
		results   map[string]int
		want      int
	}{
		{
			name:      "Call to DecrementVotesOfCandidate for John with 3 votes",
			candidate: "John",
			results: map[string]int{
				"John": 3,
			},
			want: 2,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			DecrementVotesOfCandidate(tc.results, tc.candidate)
			if votes, ok := tc.results[tc.candidate]; !ok || votes != tc.want {
				t.Errorf("DecrementVotesOfCandidate(%#v, %q) = %d, want %d",
					tc.results, tc.candidate, votes, tc.want)
			}
		})
	}
}

func intPtrRepresentation(p *int) string {
	if p == nil {
		return "nil"
	}
	return "&" + strconv.Itoa(*p)
}
