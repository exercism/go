package partyrobot

import (
	"strings"
	"testing"
)

func TestWelcome(t *testing.T) {
	tests := []struct {
		description string
		name        string
		want        string
	}{
		{
			description: "Greet Chihiro with a welcoming message",
			name:        "Chihiro",
			want:        "Welcome to my party, Chihiro!",
		},
		{
			description: "Greet Xuân Jing with a welcoming message",
			name:        "Xuân Jing",
			want:        "Welcome to my party, Xuân Jing!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			if got := Welcome(tt.name); got != tt.want {
				t.Errorf("Welcome(%s) = %s, want %s", tt.name, got, tt.want)
			}
		})
	}
}

func TestHappyBirthday(t *testing.T) {
	tests := []struct {
		description string
		name        string
		age         int
		want        string
	}{
		{
			description: "Wish Chihiro Happy Birthday with name and age",
			name:        "Chihiro",
			age:         61,
			want:        "Happy birthday Chihiro! You are now 61 years old!",
		},
		{
			description: "Wish Xuân Jing Happy Birthday with name and age",
			name:        "Xuân Jing",
			age:         17,
			want:        "Happy birthday Xuân Jing! You are now 17 years old!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			if got := HappyBirthday(tt.name, tt.age); got != tt.want {
				t.Errorf("HappyBirthday(%s, %d) = %s, want %s", tt.name, tt.age, got, tt.want)
			}
		})
	}
}

func TestAssignTable(t *testing.T) {
	tests := []struct {
		description string
		name        string
		direction   string
		tableNumber int
		distance    float64
		seatmate    string
		want        string
	}{
		{
			description: "Greet Chihiro and give them directions to their seat",
			name:        "Chihiro",
			direction:   "straight ahead",
			tableNumber: 22,
			distance:    9.2394381,
			seatmate:    "Akachi Chikondi",
			want:        "Welcome to my party, Chihiro!\nYou have been assigned to table 022. Your table is straight ahead, exactly 9.2 meters from here.\nYou will be sitting next to Akachi Chikondi.",
		},
		{
			description: "Greet Xuân Jing and give them directions to their seat",
			name:        "Xuân Jing",
			direction:   "by the façade",
			tableNumber: 4,
			distance:    23.470103,
			seatmate:    "Renée",
			want:        "Welcome to my party, Xuân Jing!\nYou have been assigned to table 004. Your table is by the façade, exactly 23.5 meters from here.\nYou will be sitting next to Renée.",
		},
		{
			description: "Greet Paula and give them directions to their seat",
			name:        "Paula",
			direction:   "on the right",
			tableNumber: 101,
			distance:    100.0,
			seatmate:    "Chioma",
			want:        "Welcome to my party, Paula!\nYou have been assigned to table 101. Your table is on the right, exactly 100.0 meters from here.\nYou will be sitting next to Chioma.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			if got := AssignTable(tt.name, tt.tableNumber, tt.seatmate, tt.direction, tt.distance); got != tt.want {
				t.Errorf(`AssignTable(%s,%d,%s,%s,%f)
got:
%q
want:
%q`,
					tt.name, tt.tableNumber, tt.seatmate, tt.direction, tt.distance, got, tt.want)

				wantLen := len(tt.want)
				gotLen := len(got)
				wantNewlines := strings.Count(tt.want, "\n")
				gotNewlines := strings.Count(got, "\n")
				wantSpaces := strings.Count(tt.want, " ")
				gotSpaces := strings.Count(got, " ")

				if wantLen != gotLen {
					t.Errorf("String lengths do not match - got: %d, want: %d", gotLen, wantLen)
				}

				if wantNewlines != gotNewlines {
					t.Errorf("Check your newline characters - got: %d, want: %d", gotNewlines, wantNewlines)
				}

				if wantSpaces != gotSpaces {
					t.Errorf("Check your space characters - got: %d, want: %d", gotSpaces, wantSpaces)
				}

			}
		})
	}
}
