package annalyn

import "testing"

type charactersState struct {
	desc            string
	knightIsAwake   bool
	archerIsAwake   bool
	prisonerIsAwake bool
	dogIsPresent    bool
	expected        bool
}

func TestCanFastAttack(t *testing.T) {
	tests := []charactersState{
		{
			desc:          "Knight is awake",
			knightIsAwake: true,
			expected:      false,
		},
		{
			desc:          "Knight is sleeping",
			knightIsAwake: false,
			expected:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			if got := CanFastAttack(tt.knightIsAwake); got != tt.expected {
				t.Errorf("CanFastAttack(%v) = %v; want %v", tt.knightIsAwake, tt.knightIsAwake, tt.expected)
			}
		})
	}
}

func TestCanSpy(t *testing.T) {
	tests := []charactersState{
		{
			desc:            "All characters are sleeping",
			knightIsAwake:   false,
			archerIsAwake:   false,
			prisonerIsAwake: false,
			expected:        false,
		},
		{
			desc:            "Knight is awake, archer and prisoner are sleeping",
			knightIsAwake:   true,
			archerIsAwake:   false,
			prisonerIsAwake: false,
			expected:        true,
		},
		{
			desc:            "Knight and archer are awake, prisoner is sleeping",
			knightIsAwake:   true,
			archerIsAwake:   true,
			prisonerIsAwake: false,
			expected:        true,
		},
		{
			desc:            "Knight and prisoner are awake, archer is sleeping",
			knightIsAwake:   true,
			archerIsAwake:   false,
			prisonerIsAwake: true,
			expected:        true,
		},
		{
			desc:            "Archer is awake, knight and prisoner are sleeping",
			knightIsAwake:   false,
			archerIsAwake:   true,
			prisonerIsAwake: false,
			expected:        true,
		},
		{
			desc:            "Archer and prisoner are awake, knight is sleeping",
			knightIsAwake:   false,
			archerIsAwake:   true,
			prisonerIsAwake: true,
			expected:        true,
		},
		{
			desc:            "Prisoner is awake, knight and archer are sleeping",
			knightIsAwake:   false,
			archerIsAwake:   false,
			prisonerIsAwake: true,
			expected:        true,
		},
		{
			desc:            "All characters are awake",
			knightIsAwake:   true,
			archerIsAwake:   true,
			prisonerIsAwake: true,
			expected:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			if got := CanSpy(tt.knightIsAwake, tt.archerIsAwake, tt.prisonerIsAwake); got != tt.expected {
				t.Errorf("CanSpy(%v, %v, %v) = %v; want %v", tt.knightIsAwake, tt.archerIsAwake, tt.prisonerIsAwake, got, tt.expected)
			}
		})
	}
}

func TestCanSignalPrisoner(t *testing.T) {
	tests := []charactersState{
		{
			desc:            "All characters are sleeping",
			archerIsAwake:   false,
			prisonerIsAwake: false,
			expected:        false,
		},
		{
			desc:            "Archer is sleeping, prisoner is awake",
			archerIsAwake:   false,
			prisonerIsAwake: true,
			expected:        true,
		},
		{
			desc:            "Archer is awake, prisoner is sleeping",
			archerIsAwake:   true,
			prisonerIsAwake: false,
			expected:        false,
		},
		{
			desc:            "All characters are awake",
			archerIsAwake:   true,
			prisonerIsAwake: true,
			expected:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			if got := CanSignalPrisoner(tt.archerIsAwake, tt.prisonerIsAwake); got != tt.expected {
				t.Errorf("CanSignalPrisoner(%v, %v) = %v; want %v", tt.archerIsAwake, tt.prisonerIsAwake, got, tt.expected)
			}
		})
	}
}

func TestCanFreePrisoner(t *testing.T) {
	tests := []charactersState{
		{
			desc:            "All characters are sleeping. Dog is not present.",
			knightIsAwake:   false,
			archerIsAwake:   false,
			prisonerIsAwake: false,
			dogIsPresent:    false,
			expected:        false,
		},
		{
			desc:            "All characters are sleeping. Dog is present.",
			knightIsAwake:   false,
			archerIsAwake:   false,
			prisonerIsAwake: false,
			dogIsPresent:    true,
			expected:        true,
		},
		{
			desc:            "Knight and archer are sleeping. Prisoner is awake. Dog is not present.",
			knightIsAwake:   false,
			archerIsAwake:   false,
			prisonerIsAwake: true,
			dogIsPresent:    false,
			expected:        true,
		},
		{
			desc:            "Knight and archer are sleeping. Prisoner is awake. Dog is present.",
			knightIsAwake:   false,
			archerIsAwake:   false,
			prisonerIsAwake: true,
			dogIsPresent:    true,
			expected:        true,
		},
		{
			desc:            "Knight is sleeping. Archer is awake. Prisoner is sleeping. Dog is not present.",
			knightIsAwake:   false,
			archerIsAwake:   true,
			prisonerIsAwake: false,
			dogIsPresent:    false,
			expected:        false,
		},
		{
			desc:            "Knight is sleeping. Archer is awake. Prisoner is sleeping. Dog is present.",
			knightIsAwake:   false,
			archerIsAwake:   true,
			prisonerIsAwake: false,
			dogIsPresent:    true,
			expected:        false,
		},
		{
			desc:            "Knight is sleeping. Archer is awake. Prisoner is awake. Dog is not present.",
			knightIsAwake:   false,
			archerIsAwake:   true,
			prisonerIsAwake: true,
			dogIsPresent:    false,
			expected:        false,
		},
		{
			desc:            "Knight is sleeping. Archer is awake. Prisoner is awake. Dog is present.",
			knightIsAwake:   false,
			archerIsAwake:   true,
			prisonerIsAwake: true,
			dogIsPresent:    true,
			expected:        false,
		},
		{
			desc:            "Knight is awake. Archer is sleeping. Prisoner is sleeping. Dog is not present.",
			knightIsAwake:   true,
			archerIsAwake:   false,
			prisonerIsAwake: false,
			dogIsPresent:    false,
			expected:        false,
		},
		{
			desc:            "Knight is awake. Archer is sleeping. Prisoner is sleeping. Dog is present.",
			knightIsAwake:   true,
			archerIsAwake:   false,
			prisonerIsAwake: false,
			dogIsPresent:    true,
			expected:        true,
		},
		{
			desc:            "Knight is awake. Archer is sleeping. Prisoner is awake. Dog is not present",
			knightIsAwake:   true,
			archerIsAwake:   false,
			prisonerIsAwake: true,
			dogIsPresent:    false,
			expected:        false,
		},
		{
			desc:            "Knight is awake. Archer is sleeping. Prisoner is awake. Dog is present",
			knightIsAwake:   true,
			archerIsAwake:   false,
			prisonerIsAwake: true,
			dogIsPresent:    true,
			expected:        true,
		},
		{
			desc:            "Knight and archer are awake. Prisoner is sleeping. Dog is not present",
			knightIsAwake:   true,
			archerIsAwake:   true,
			prisonerIsAwake: false,
			dogIsPresent:    false,
			expected:        false,
		},
		{
			desc:            "Knight and archer are awake. Prisoner is sleeping. Dog is present",
			knightIsAwake:   true,
			archerIsAwake:   true,
			prisonerIsAwake: false,
			dogIsPresent:    true,
			expected:        false,
		},
		{
			desc:            "Knight and archer are awake. Prisoner is awake. Dog is not present",
			knightIsAwake:   true,
			archerIsAwake:   true,
			prisonerIsAwake: true,
			dogIsPresent:    false,
			expected:        false,
		},
		{
			desc:            "Knight and archer are awake. Prisoner is awake. Dog is present",
			knightIsAwake:   true,
			archerIsAwake:   true,
			prisonerIsAwake: true,
			dogIsPresent:    true,
			expected:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			if got := CanFreePrisoner(tt.knightIsAwake, tt.archerIsAwake, tt.prisonerIsAwake, tt.dogIsPresent); got != tt.expected {
				t.Errorf("CanFreePrisoner(%v,%v,%v,%v) = %v; want %v", tt.knightIsAwake, tt.archerIsAwake, tt.prisonerIsAwake, tt.dogIsPresent, got, tt.expected)
			}
		})
	}
}
