package parsinglogfiles

import "testing"

func TestIsValidLine(t *testing.T) {
	tests := []struct {
		description string
		text        string
		expected    bool
	}{
		{
			description: "Valid ERR message",
			text:        "[ERR] A good error here",
			expected:    true,
		},
		{
			description: "Valid INF message",
			text:        "[INF] The latest information",
			expected:    true,
		},
		{
			description: "Invalid ERR message",
			text:        "Any old [ERR] text",
			expected:    false,
		},
		{
			description: "Invalid INF message",
			text:        "bad start to [INF] Message",
			expected:    false,
		},
		{
			description: "Invalid tag",
			text:        "[BOB] Any old text",
			expected:    false,
		},
		{
			description: "Line with less characters than 5",
			text:        "text",
			expected:    false,
		},
		{
			description: "Empty line",
			text:        "",
			expected:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := IsValidLine(tt.text)
			want := tt.expected
			if want != got {
				t.Fatalf("expected: %v, got: %v", want, got)
			}
		})
	}
}

func TestSplitLogLine(t *testing.T) {
	tests := []struct {
		description string
		text        string
		expected    []string
	}{
		{
			description: "three sections",
			text:        "section 1<*>section 2<~~~>section 3",
			expected:    []string{"section 1", "section 2", "section 3"},
		},
		{
			description: "three sections with different symbols inside angular brackets",
			text:        "section 1<=>section 2<-*~*->section 3",
			expected:    []string{"section 1", "section 2", "section 3"},
		},
		{
			description: "two sections with no characters between angular brackets",
			text:        "section 1<>section 2",
			expected:    []string{"section 1", "section 2"},
		},
		{
			description: "single section with some angular brackets",
			text:        "<start> <end>",
			expected:    []string{"<start> <end>"},
		},
		{
			description: "empty text",
			text:        "",
			expected:    []string{""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := SplitLogLine(tt.text)
			want := tt.expected
			if !stringSliceEqual(want, got) {
				t.Fatalf("expected: %v, got: %v", want, got)
			}
		})
	}
}

func TestCountQuotedPasswords(t *testing.T) {
	tests := []struct {
		description string
		lines       []string
		expected    int
	}{
		{
			description: "text with two matches",
			lines: []string{
				``,
				`[INF] passWord`,
				`"passWord"`,
				`[INF] User saw error message "Unexpected Error" on page load.`,
				`[INF] The message "Please reset your password" was ignored by the user`,
			},
			expected: 2,
		},
		{
			description: "text with no matches",
			lines: []string{
				`passWord"passWor"`,
			},
			expected: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := CountQuotedPasswords(tt.lines)
			want := tt.expected
			if want != got {
				t.Fatalf("expected: %v, got: %v", want, got)
			}
		})
	}
}

func TestRemoveEndOfLineText(t *testing.T) {
	tests := []struct {
		description string
		text        string
		expected    string
	}{
		{
			description: "INF message",
			text:        "[INF] end-of-line23033 Network Failure end-of-line27",
			expected:    "[INF]  Network Failure ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := RemoveEndOfLineText(tt.text)
			want := tt.expected
			if want != got {
				t.Fatalf("expected: %v, got: %v", want, got)
			}
		})
	}
}

func TestTagWithUserName(t *testing.T) {
	tests := []struct {
		description string
		lines       []string
		expected    []string
	}{
		{
			description: "INF message",
			lines: []string{
				"[WRN] User James123 has exceeded storage space.",
				"[WRN] Host down. User   Michelle4 lost connection.",
				"[INF] Users can login again after 23:00.",
				"[DBG] We need to check that user names are at least 6 chars long.",
			},
			expected: []string{
				"[USR] James123 [WRN] User James123 has exceeded storage space.",
				"[USR] Michelle4 [WRN] Host down. User   Michelle4 lost connection.",
				"[INF] Users can login again after 23:00.",
				"[DBG] We need to check that user names are at least 6 chars long.",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := TagWithUserName(tt.lines)
			want := tt.expected
			if !stringSliceEqual(want, got) {
				for i := range got {
					if want[i] != got[i] {
						t.Fatalf("expected: %v, got: %v", want[i], got[i])
					}
				}
				t.Fatalf("expected: %v, got: %v", want, got)
			}
		})
	}
}

func stringSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 {
		return true
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
