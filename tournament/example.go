package tournament

import (
	"encoding/csv"
	"fmt"
	"io"
	"sort"
)

const TestVersion = 1

type outcome int

const (
	LOSS outcome = iota
	DRAW
	WIN
)

type inputEntry struct {
	teams    [2]string
	outcomes [2]outcome
}

type teamResult struct {
	team   string
	played int
	wins   int
	draws  int
	losses int
	points int
}

type TeamResultSlice []teamResult

// sort.Interface implementation, sorts on points, descending
func (s TeamResultSlice) Len() int {
	return len(s)
}

func (s TeamResultSlice) Less(i, j int) bool {
	switch {
	case s[i].points != s[j].points:
		return s[i].points > s[j].points
	case s[i].wins != s[j].wins:
		return s[i].wins > s[j].wins
	default:
		return s[i].team < s[j].team
	}
}

func (s TeamResultSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func readInput(reader io.Reader) ([]inputEntry, error) {
	var entries []inputEntry
	csvReader := csv.NewReader(reader)
	csvReader.Comma = ';'
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if len(record) == 3 {
			t1, t2 := record[0], record[1]
			teams := [2]string{t1, t2}
			var outcomes [2]outcome
			switch record[2] {
			case "win":
				outcomes = [2]outcome{WIN, LOSS}
			case "loss":
				outcomes = [2]outcome{LOSS, WIN}
			case "draw":
				outcomes = [2]outcome{DRAW, DRAW}
			default:
				return nil, fmt.Errorf("Invalid outcome %q", record[2])
			}
			entries = append(entries, inputEntry{teams: teams, outcomes: outcomes})
		} else {
			return nil, fmt.Errorf("Invalid record: %v", record)
		}
	}
	return entries, nil
}

func tallyEntries(entries []inputEntry) map[string]teamResult {
	var results = make(map[string]teamResult)
	for _, entry := range entries {
		for i := 0; i < 2; i++ {
			team := entry.teams[i]
			outcome := entry.outcomes[i]
			result, present := results[team]
			if !present {
				result.team = team
				// The rest is 0, which is correct
			}
			switch outcome {
			case WIN:
				result.wins += 1
				result.points += 3
			case DRAW:
				result.draws += 1
				result.points += 1
			case LOSS:
				result.losses += 1
			}
			result.played += 1
			results[team] = result
		}
	}
	return results
}

func report(writer io.Writer, resultMap map[string]teamResult) {
	var entries TeamResultSlice = make([]teamResult, 0, len(resultMap))
	for _, entry := range resultMap {
		entries = append(entries, entry)
	}
	sort.Sort(entries)
	fmt.Fprintf(writer, "Team                           | MP |  W |  D |  L |  P\n")
	for _, entry := range entries {
		fmt.Fprintf(writer, "%-30s | %2d | %2d | %2d | %2d | %2d\n",
			entry.team, entry.played, entry.wins, entry.draws, entry.losses, entry.points)
	}
}

func Tally(reader io.Reader, writer io.Writer) error {
	entries, err := readInput(reader)
	if err != nil {
		return err
	}
	report(writer, tallyEntries(entries))
	return nil
}
