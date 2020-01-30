package basic_strings

import (
	"strings"
	"testing"
)

func TestMessage(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Extract message from error message",
			args: args{
				line: "[ERROR]: Stack overflow",
			},
			want: "Stack overflow",
		},
		{
			name: "Extract message from warning message",
			args: args{
				line: "[WARNING]: Disk almost full",
			},
			want: "Disk almost full",
		},
		{
			name: "Extract message from info message",
			args: args{
				line: "[INFO]: File moved",
			},
			want: "File moved",
		},
		{
			name: "Extract message without extra whitespace",
			args: args{
				line: "[WARNING]:   \tTimezone not set  \r\n",
			},
			want: "Timezone not set",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Message(tt.args.line); got != tt.want {
				t.Errorf("Message(\"%s\") = \"%s\", want \"%s\"", escapeWhiteSpace(tt.args.line), got, tt.want)
			}
		})
	}
}

func TestMessageLen(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Message length from error message",
			args: args{
				line: "[ERROR]: Stack overflow",
			},
			want: 14,
		},
		{
			name: "Message length from warning message",
			args: args{
				line: "[WARNING]: Disk almost full",
			},
			want: 16,
		},
		{
			name: "Message length from info message",
			args: args{
				line: "[INFO]: File moved",
			},
			want: 10,
		},
		{
			name: "Message length without extra whitespace",
			args: args{
				line: "[WARNING]:   \tTimezone not set  \r\n",
			},
			want: 16,
		},
		{
			name: "Message length with special characters",
			args: args{
				line: "[INFO]: Hello, 世界!",
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MessageLen(tt.args.line); got != tt.want {
				t.Errorf("MessageLen(\"%v\") = \"%v\", want \"%v\"", escapeWhiteSpace(tt.args.line), got, tt.want)
			}
		})
	}
}

func TestLogLevel(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Extract log level from error",
			args: args{
				line: "[ERROR]: Disk full",
			},
			want: "error",
		},
		{
			name: "Extract log level from warning",
			args: args{
				line: "[WARNING]: Unsafe password",
			},
			want: "warning",
		},
		{
			name: "Extract log level from info",
			args: args{
				line: "[INFO]: Timezone changed",
			},
			want: "info",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LogLevel(tt.args.line); got != tt.want {
				t.Errorf("LogLevel(\"%s\") = \"%s\", want \"%s\"", escapeWhiteSpace(tt.args.line), got, tt.want)
			}
		})
	}
}

func TestReformat(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Reformat error message",
			args: args{
				line: "[ERROR]: Segmentation fault",
			},
			want: "Segmentation fault (error)",
		},
		{
			name: "Reformat warning message",
			args: args{
				line: "[WARNING]: Decreased performance",
			},
			want: "Decreased performance (warning)",
		},
		{
			name: "Reformat info message",
			args: args{
				line: "[INFO]: Disk defragmented",
			},
			want: "Disk defragmented (info)",
		},
		{
			name: "Reformat message with extra whitespace",
			args: args{
				line: "[ERROR]: \t Corrupt disk\t \t \r\n",
			},
			want: "Corrupt disk (error)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reformat(tt.args.line); got != tt.want {
				t.Errorf("Reformat(\"%s\") = \"%s\", want \"%s\"", escapeWhiteSpace(tt.args.line), got, tt.want)
			}
		})
	}
}

func escapeWhiteSpace(s string) string {
	s = strings.ReplaceAll(s, "\\", "\\\\")
	s = strings.ReplaceAll(s, "\n", "\\n")
	s = strings.ReplaceAll(s, "\r", "\\r")
	return strings.ReplaceAll(s, "\t", "\\t")
}
