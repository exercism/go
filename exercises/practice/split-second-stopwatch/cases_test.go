package splitsecondstopwatch

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: f8c8521 Add `split-second-stopwatch` exercise (#2547)

type operation struct {
	command          string
	by               string
	expected         string
	expectedPrevLaps []string
	expectedErr      string
}

var testCases = []struct {
	description string
	commands    []operation
}{
	{
		description: "new stopwatch starts in ready state",
		commands: []operation{

			{
				command: "New",
			},
			{
				command:  "State",
				expected: "ready",
			},
		},
	},
	{
		description: "new stopwatch's current lap has no elapsed time",
		commands: []operation{

			{
				command: "New",
			},
			{
				command:  "CurrentLap",
				expected: "00:00:00",
			},
		},
	},
	{
		description: "new stopwatch's total has no elapsed time",
		commands: []operation{

			{
				command: "New",
			},
			{
				command:  "Total",
				expected: "00:00:00",
			},
		},
	},
	{
		description: "new stopwatch does not have previous laps",
		commands: []operation{

			{
				command: "New",
			},
			{
				command:          "PreviousLaps",
				expectedPrevLaps: []string{},
			},
		},
	},
	{
		description: "start from ready state changes state to running",
		commands: []operation{

			{
				command: "New",
			},
			{
				command: "Start",
			},
			{
				command:  "State",
				expected: "running",
			},
		},
	},
	{
		description: "start does not change previous laps",
		commands: []operation{

			{
				command: "New",
			},
			{
				command: "Start",
			},
			{
				command:          "PreviousLaps",
				expectedPrevLaps: []string{},
			},
		},
	},
	{
		description: "start initiates time tracking for current lap",
		commands: []operation{

			{
				command: "New",
			},
			{
				command: "Start",
			},
			{
				command: "AdvanceTime",
				by:      "00:00:05",
			},
			{
				command:  "CurrentLap",
				expected: "00:00:05",
			},
		},
	},
	{
		description: "start initiates time tracking for total",
		commands: []operation{

			{
				command: "New",
			},
			{
				command: "Start",
			},
			{
				command: "AdvanceTime",
				by:      "00:00:23",
			},
			{
				command:  "Total",
				expected: "00:00:23",
			},
		},
	},
	{
		description: "start cannot be called from running state",
		commands: []operation{

			{
				command: "New",
			},
			{
				command: "Start",
			},
			{
				command:     "Start",
				expectedErr: "cannot start an already running stopwatch",
			},
		},
	},
	{
		description: "stop from running state changes state to stopped",
		commands: []operation{

			{
				command: "New",
			},
			{
				command: "Start",
			},
			{
				command: "Stop",
			},
			{
				command:  "State",
				expected: "stopped",
			},
		},
	},
	{
		description: "stop pauses time tracking for current lap",
		commands: []operation{

			{
				command: "New",
			},
			{
				command: "Start",
			},
			{
				command: "AdvanceTime",
				by:      "00:00:05",
			},
			{
				command: "Stop",
			},
			{
				command: "AdvanceTime",
				by:      "00:00:08",
			},
			{
				command:  "CurrentLap",
				expected: "00:00:05",
			},
		},
	},
	{
		description: "stop pauses time tracking for total",
		commands: []operation{

			{
				command: "New",
			},
			{
				command: "Start",
			},
			{
				command: "AdvanceTime",
				by:      "00:00:13",
			},
			{
				command: "Stop",
			},
			{
				command: "AdvanceTime",
				by:      "00:00:44",
			},
			{
				command:  "Total",
				expected: "00:00:13",
			},
		},
	},
	{
		description: "stop cannot be called from ready state",
		commands: []operation{

			{
				command: "New",
			},
			{
				command:     "Stop",
				expectedErr: "cannot stop a stopwatch that is not running",
			},
		},
	},
	{
		description: "stop cannot be called from stopped state",
		commands: []operation{

			{
				command: "New",
			},
			{
				command: "Start",
			},
			{
				command: "Stop",
			},
			{
				command:     "Stop",
				expectedErr: "cannot stop a stopwatch that is not running",
			},
		},
	},
	{
		description: "start from stopped state changes state to running",
		commands: []operation{

			{
				command: "New",
			},
			{
				command: "Start",
			},
			{
				command: "Stop",
			},
			{
				command: "Start",
			},
			{
				command:  "State",
				expected: "running",
			},
		},
	},
	{
		description: "start from stopped state resumes time tracking for current lap",
		commands: []operation{

			{
				command: "New",
			},
			{
				command: "Start",
			},
			{
				command: "AdvanceTime",
				by:      "00:01:20",
			},
			{
				command: "Stop",
			},
			{
				command: "AdvanceTime",
				by:      "00:00:20",
			},
			{
				command: "Start",
			},
			{
				command: "AdvanceTime",
				by:      "00:00:08",
			},
			{
				command:  "CurrentLap",
				expected: "00:01:28",
			},
		},
	},
	{
		description: "start from stopped state resumes time tracking for total",
		commands: []operation{

			{
				command: "New",
			},
			{
				command: "Start",
			},
			{
				command: "AdvanceTime",
				by:      "00:00:23",
			},
			{
				command: "Stop",
			},
			{
				command: "AdvanceTime",
				by:      "00:00:44",
			},
			{
				command: "Start",
			},
			{
				command: "AdvanceTime",
				by:      "00:00:09",
			},
			{
				command:  "Total",
				expected: "00:00:32",
			},
		},
	},
	{
		description: "lap adds current lap to previous laps",
		commands: []operation{

			{
				command: "New",
			},
			{
				command: "Start",
			},
			{
				command: "AdvanceTime",
				by:      "00:01:38",
			},
			{
				command: "Lap",
			},
			{
				command:          "PreviousLaps",
				expectedPrevLaps: []string{"00:01:38"},
			},
			{
				command: "AdvanceTime",
				by:      "00:00:44",
			},
			{
				command: "Lap",
			},
			{
				command:          "PreviousLaps",
				expectedPrevLaps: []string{"00:01:38", "00:00:44"},
			},
		},
	},
	{
		description: "lap resets current lap and resumes time tracking",
		commands: []operation{

			{
				command: "New",
			},
			{
				command: "Start",
			},
			{
				command: "AdvanceTime",
				by:      "00:08:22",
			},
			{
				command: "Lap",
			},
			{
				command:  "CurrentLap",
				expected: "00:00:00",
			},
			{
				command: "AdvanceTime",
				by:      "00:00:15",
			},
			{
				command:  "CurrentLap",
				expected: "00:00:15",
			},
		},
	},
	{
		description: "lap continues time tracking for total",
		commands: []operation{

			{
				command: "New",
			},
			{
				command: "Start",
			},
			{
				command: "AdvanceTime",
				by:      "00:00:22",
			},
			{
				command: "Lap",
			},
			{
				command: "AdvanceTime",
				by:      "00:00:33",
			},
			{
				command:  "Total",
				expected: "00:00:55",
			},
		},
	},
	{
		description: "lap cannot be called from ready state",
		commands: []operation{

			{
				command: "New",
			},
			{
				command:     "Lap",
				expectedErr: "cannot lap a stopwatch that is not running",
			},
		},
	},
	{
		description: "lap cannot be called from stopped state",
		commands: []operation{

			{
				command: "New",
			},
			{
				command: "Start",
			},
			{
				command: "Stop",
			},
			{
				command:     "Lap",
				expectedErr: "cannot lap a stopwatch that is not running",
			},
		},
	},
	{
		description: "stop does not change previous laps",
		commands: []operation{

			{
				command: "New",
			},
			{
				command: "Start",
			},
			{
				command: "AdvanceTime",
				by:      "00:11:22",
			},
			{
				command: "Lap",
			},
			{
				command:          "PreviousLaps",
				expectedPrevLaps: []string{"00:11:22"},
			},
			{
				command: "Stop",
			},
			{
				command:          "PreviousLaps",
				expectedPrevLaps: []string{"00:11:22"},
			},
		},
	},
	{
		description: "reset from stopped state changes state to ready",
		commands: []operation{

			{
				command: "New",
			},
			{
				command: "Start",
			},
			{
				command: "Stop",
			},
			{
				command: "Reset",
			},
			{
				command:  "State",
				expected: "ready",
			},
		},
	},
	{
		description: "reset resets current lap",
		commands: []operation{

			{
				command: "New",
			},
			{
				command: "Start",
			},
			{
				command: "AdvanceTime",
				by:      "00:00:10",
			},
			{
				command: "Stop",
			},
			{
				command: "Reset",
			},
			{
				command:  "CurrentLap",
				expected: "00:00:00",
			},
		},
	},
	{
		description: "reset clears previous laps",
		commands: []operation{

			{
				command: "New",
			},
			{
				command: "Start",
			},
			{
				command: "AdvanceTime",
				by:      "00:00:10",
			},
			{
				command: "Lap",
			},
			{
				command: "AdvanceTime",
				by:      "00:00:20",
			},
			{
				command: "Lap",
			},
			{
				command:          "PreviousLaps",
				expectedPrevLaps: []string{"00:00:10", "00:00:20"},
			},
			{
				command: "Stop",
			},
			{
				command: "Reset",
			},
			{
				command:          "PreviousLaps",
				expectedPrevLaps: []string{},
			},
		},
	},
	{
		description: "reset cannot be called from ready state",
		commands: []operation{

			{
				command: "New",
			},
			{
				command:     "Reset",
				expectedErr: "cannot reset a stopwatch that is not stopped",
			},
		},
	},
	{
		description: "reset cannot be called from running state",
		commands: []operation{

			{
				command: "New",
			},
			{
				command: "Start",
			},
			{
				command:     "Reset",
				expectedErr: "cannot reset a stopwatch that is not stopped",
			},
		},
	},
	{
		description: "supports very long laps",
		commands: []operation{

			{
				command: "New",
			},
			{
				command: "Start",
			},
			{
				command: "AdvanceTime",
				by:      "01:23:45",
			},
			{
				command:  "CurrentLap",
				expected: "01:23:45",
			},
			{
				command: "Lap",
			},
			{
				command:          "PreviousLaps",
				expectedPrevLaps: []string{"01:23:45"},
			},
			{
				command: "AdvanceTime",
				by:      "04:01:40",
			},
			{
				command:  "CurrentLap",
				expected: "04:01:40",
			},
			{
				command:  "Total",
				expected: "05:25:25",
			},
			{
				command: "Lap",
			},
			{
				command:          "PreviousLaps",
				expectedPrevLaps: []string{"01:23:45", "04:01:40"},
			},
			{
				command: "AdvanceTime",
				by:      "08:43:05",
			},
			{
				command:  "CurrentLap",
				expected: "08:43:05",
			},
			{
				command:  "Total",
				expected: "14:08:30",
			},
			{
				command: "Lap",
			},
			{
				command:          "PreviousLaps",
				expectedPrevLaps: []string{"01:23:45", "04:01:40", "08:43:05"},
			},
		},
	},
}
