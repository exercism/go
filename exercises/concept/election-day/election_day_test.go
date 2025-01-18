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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewVoteCounter(tt.votes)
			if got == nil {
				t.Errorf("NewVoteCounter(%d) = %s, &%d", tt.votes, intPtrRepresentation(got), tt.votes)
			}
			if got != nil && *got != tt.votes {
				t.Errorf("NewVoteCounter(%d) = %s, &%d", tt.votes, intPtrRepresentation(got), tt.votes)
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VoteCount(tt.counter); got != tt.expected {
				t.Fatalf("VoteCount(%v) = %d, want %d", intPtrRepresentation(tt.counter), got, tt.expected)
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			before := intPtrRepresentation(tt.counter)
			IncrementVoteCount(tt.counter, tt.increment)
			after := intPtrRepresentation(tt.counter)

			if tt.counter == nil {
				t.Errorf("counter before: %s | counter after: %v | wanted: &%d", before, after, tt.expected)
			}

			if tt.counter != nil && *tt.counter != tt.expected {
				t.Errorf("counter before: %s | counter after: %v | wanted: &%d", before, after, tt.expected)
			}
		})
	}
}

func TestNewElectionResult(t *testing.T) {
	tests := []struct {
		name          string
		candidateName string
		votes         int
		wanted        ElectionResult
	}{
		{
			name:          "Call to NewElectionResult for Peter with 2 votes",
			candidateName: "Peter",
			votes:         2,
			wanted: ElectionResult{
				Name:  "Peter",
				Votes: 2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewElectionResult(tt.candidateName, tt.votes)

			if result == nil || result.Name != tt.wanted.Name || result.Votes != tt.wanted.Votes {
				t.Errorf("NewElectionResult(\"%s\", %d) = %#v, wanted %#v",
					tt.candidateName, tt.votes, result, tt.wanted)
			}
		})
	}
}

func TestDisplayResult(t *testing.T) {
	tests := []struct {
		name   string
		result *ElectionResult
		wanted string
	}{
		{
			name: "Call to DisplayResult for John with 5 votes",
			result: &ElectionResult{
				Name:  "John",
				Votes: 5,
			},
			wanted: "John (5)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := DisplayResult(tt.result); result != tt.wanted {
				t.Errorf("DisplayResult(%#v) = %s, wanted %s", *tt.result, result, tt.wanted)
			}
		})
	}
}

func TestDecrementVotesOfCandidate(t *testing.T) {
	tests := []struct {
		name      string
		candidate string
		results   map[string]int
		wanted    int
	}{
		{
			name:      "Call to DecrementVotesOfCandidate for John with 3 votes",
			candidate: "John",
			results: map[string]int{
				"John": 3,
			},
			wanted: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DecrementVotesOfCandidate(tt.results, tt.candidate)
			if votes, ok := tt.results[tt.candidate]; !ok || votes != tt.wanted {
				t.Errorf("DecrementVotesOfCandidate(%v) | wanted %d, got %d",
					tt.results, tt.wanted, votes)
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
