package airportrobot

import "testing"

// testRunnerTaskID=2
func TestSayHello_Italien(t *testing.T) {
	tests := []struct {
		testName string
		name     string
		want     string
	}{
		{
			testName: "name without spaces",
			name:     "Flora",
			want:     "I can speak Italian: Ciao Flora!",
		},
		{
			testName: "full name",
			name:     "Tomaso Giulio Micheli",
			want:     "I can speak Italian: Ciao Tomaso Giulio Micheli!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			if got := SayHello(tt.name, Italian{}); got != tt.want {
				t.Errorf("SayHello(%q, \"Italian{}\") = %q, want %q", tt.name, got, tt.want)
			}
		})
	}
}

// testRunnerTaskID=3
func TestSayHello_Portuguese(t *testing.T) {
	tests := []struct {
		testName string
		name     string
		want     string
	}{
		{
			testName: "name without spaces",
			name:     "Fabrício",
			want:     "I can speak Portuguese: Olá Fabrício!",
		},
		{
			testName: "full name",
			name:     "Manuela Alberto",
			want:     "I can speak Portuguese: Olá Manuela Alberto!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			if got := SayHello(tt.name, Portuguese{}); got != tt.want {
				t.Errorf("SayHello(%q, \"Portuguese{}\") = %q, want %q", tt.name, got, tt.want)
			}
		})
	}
}
