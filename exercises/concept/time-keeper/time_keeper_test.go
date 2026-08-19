package time_keeper

import (
	"testing"
	"time"
)

func TestGetDuration(t *testing.T) {
	start := time.Date(2026, 4, 20, 9, 0, 0, 0, time.UTC)
	finish := time.Date(2026, 4, 20, 10, 5, 19, 234000000, time.UTC)

	got := GetDuration(start, finish)
	want := 1*time.Hour + 5*time.Minute + 19*time.Second + 234*time.Millisecond

	if got != want {
		t.Errorf("GetDuration() = %v, want %v", got, want)
	}
}

func TestGetDurationInDifferentTimezones(t *testing.T) {
	newYork, err := time.LoadLocation("America/New_York")
	if err != nil {
		t.Fatalf("failed to load New York location: %v", err)
	}

	regina, err := time.LoadLocation("America/Regina")
	if err != nil {
		t.Fatalf("failed to load Regina location: %v", err)
	}

	start := time.Date(2026, 4, 20, 8, 0, 0, 0, newYork)
	finish := time.Date(2026, 4, 20, 8, 30, 0, 0, regina)

	got := GetDurationInDifferentTimezones(start, finish)

	startUTC := start.UTC()
	finishUTC := finish.UTC()
	want := finishUTC.Sub(startUTC)

	if got != want {
		t.Errorf("GetDurationInDifferentTimezones() = %v, want %v", got, want)
	}
}

func TestParseRun(t *testing.T) {
	got, err := ParseRun("10h5m19.234s")
	if err != nil {
		t.Fatalf("ParseRun() returned error: %v", err)
	}

	want := 10*time.Hour + 5*time.Minute + 19*time.Second + 234*time.Millisecond

	if got != want {
		t.Errorf("ParseRun() = %v, want %v", got, want)
	}
}

func TestParseRunInvalid(t *testing.T) {
	_, err := ParseRun("10hours5minutes")

	if err == nil {
		t.Error("ParseRun() expected error for invalid input, got nil")
	}
}

func TestFixStartTime(t *testing.T) {
	incorrectStart := time.Date(2026, 4, 20, 9, 0, 15, 0, time.UTC)
	offset := 15 * time.Second

	got := FixStartTime(incorrectStart, offset)
	want := time.Date(2026, 4, 20, 9, 0, 0, 0, time.UTC)

	if !got.Equal(want) {
		t.Errorf("FixStartTime() = %v, want %v", got, want)
	}
}

func TestIsValidDetection(t *testing.T) {
	start := time.Date(2026, 4, 20, 9, 0, 0, 0, time.UTC)

	tests := []struct {
		name      string
		detection time.Time
		want      bool
	}{
		{
			name:      "before lower bound",
			detection: start.Add(9 * time.Minute),
			want:      false,
		},
		{
			name:      "exactly at lower bound",
			detection: start.Add(10 * time.Minute),
			want:      true,
		},
		{
			name:      "inside range",
			detection: start.Add(20 * time.Minute),
			want:      true,
		},
		{
			name:      "exactly at upper bound",
			detection: start.Add(30 * time.Minute),
			want:      true,
		},
		{
			name:      "after upper bound",
			detection: start.Add(31 * time.Minute),
			want:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValidDetection(start, 10*time.Minute, 30*time.Minute, tt.detection)
			if got != tt.want {
				t.Errorf("IsValidDetection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatResult(t *testing.T) {
	result := 1*time.Hour + 5*time.Minute + 19*time.Second + 234*time.Millisecond
	leader := 1*time.Hour + 5*time.Minute + 14*time.Second + 391*time.Millisecond

	got := FormatResult(result, leader)
	want := "1h5m19.234s +4.843"

	if got != want {
		t.Errorf("FormatResult() = %q, want %q", got, want)
	}
}

func TestFormatResultLeaderHasSameTime(t *testing.T) {
	result := 42*time.Minute + 12*time.Second
	leader := 42*time.Minute + 12*time.Second

	got := FormatResult(result, leader)
	want := "42m12s +0.000"

	if got != want {
		t.Errorf("FormatResult() = %q, want %q", got, want)
	}
}
