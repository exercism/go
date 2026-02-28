package splitsecondstopwatch

import (
	"fmt"
	"slices"
	"strings"
	"testing"
)

type StopwatchTest struct {
	stopwatch *SplitSecondStopwatch
	commands []string
}

func (st *StopwatchTest) commandHistory() string {
	return strings.Join(st.commands, "; ")
}

func (st *StopwatchTest) Run(op operation) (string, []string, error) {
	if op.command == "New" {
		st.commands = append(st.commands, "NewSplitSecondStopwatch()")
	} else if op.by != "" {
		st.commands = append(st.commands, fmt.Sprintf("stopwatch.%s(%s)", op.command, op.by))
	} else {
		st.commands = append(st.commands, fmt.Sprintf("stopwatch.%s()", op.command))
	}

	switch op.command {
	case "New":
		st.stopwatch = NewSplitSecondStopwatch()
		return "", nil, nil
	case "Start":
		return "", nil, st.stopwatch.Start()
	case "Reset":
		return "", nil, st.stopwatch.Reset()
	case "Lap":
		return "", nil, st.stopwatch.Lap()
	case "Stop":
		return "", nil, st.stopwatch.Stop()
	case "AdvanceTime":
		st.stopwatch.AdvanceTime(op.by)
		return "", nil, nil
	case "State":
		return st.stopwatch.State(), nil, nil
	case "CurrentLap":
		return st.stopwatch.CurrentLap(), nil, nil
	case "Total":
		return st.stopwatch.Total(), nil, nil
	case "PreviousLaps":
		return "", st.stopwatch.PreviousLaps(), nil
	default:
		panic(fmt.Sprintf("unknown operation, %q", op.command))
	}
}

func TestTime(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			st := StopwatchTest{}
			for _, operation := range tc.commands {
				actual, previousLaps, err := st.Run(operation)
				if operation.expectedErr == "" {
					if err != nil {
						t.Fatalf("%s, got unexpected error %q", st.commandHistory(), err)
					}
				} else {
					if err == nil {
						t.Fatalf("%s, did not error, expected error %q", st.commandHistory(), operation.expectedErr)
					} else if operation.expectedErr != err.Error() {
						t.Fatalf("%s, expected error %q, got error %q", st.commandHistory(), operation.expectedErr, err)
					}
				}

				if operation.expected != "" && actual != operation.expected {
					t.Fatalf("%s = %q, want %q", st.commandHistory(), actual, operation.expected)
				}

				if len(operation.expectedPrevLaps) != 0 && !slices.Equal(previousLaps, operation.expectedPrevLaps) {
					t.Fatalf("%s = %q, want %q", st.commandHistory(), previousLaps, operation.expectedPrevLaps)
				}
			}
		})
	}
}

func BenchmarkTime(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			st := StopwatchTest{}
			for _, operation := range tc.commands {
				st.Run(operation)
			}
		}
	}
}
