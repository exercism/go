package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)
const lineFmt    = "%-31s| %2d | %2d | %2d | %2d | %2d\n"
const headerFmt  = "%-31s| %2s | %2s | %2s | %2s | %2s\n"

type team struct {
	name   string
	won    int
	lost   int
	drawn  int
}

type teams map[string]team
type ranking []team

func (t teams) getTeam(n string) team {
	if _, ok := t[n]; !ok {
		t[n] = team{name: n}
	}
	return t[n]
}

func (t teams) parse(reader io.Reader) error {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || line[0] == '#' {
			continue
		}
		chunks := strings.Split(line, ";")
		if len(chunks) != 3 {
			return fmt.Errorf("wrong field count for line: %s (got %d fields)", line, len(chunks))
		}
		home, away, result := t.getTeam(chunks[0]), t.getTeam(chunks[1]), chunks[2]
		switch result {
		case "win":
			home.win()
			away.lose()
		case "loss":
			home.lose()
			away.win()
		case "draw":
			home.draw()
			away.draw()
		default:
			return fmt.Errorf("Invalid or missing result missing: %s", line)
		}
		// uh... yeah, this is a smell
		t[home.name], t[away.name]= home, away
	}
	return nil
}

func (t teams) Rank() []team {
	ranking := make(ranking, len(t))
	i := 0
	for _, v := range t {
		ranking[i] = v
		i++
	}
	sort.Sort(sort.Reverse(ranking))
	return ranking
}

func (l ranking) Len() int { return len(l) }
func (l ranking) Less(i, j int) bool {
	delta := int(l[i].points()) - int(l[j].points())
	if delta == 0 {
		return l[i].name > l[j].name
	}
	return delta < 0
}
func (l ranking) Swap(i, j int) { l[i], l[j] = l[j], l[i] }
func (t *team) win() {
	t.won++
}

func (t *team) lose() {
	t.lost++
}

func (t *team) draw() {
	t.drawn++
}

func (t *team) played() int {
	return t.won + t.drawn + t.lost
}

func (t *team) points() int {
	return 3*t.won + t.drawn
}

func Tally(r io.Reader, w io.Writer) error {
	t := make(teams)
	err := t.parse(r)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(w, headerFmt, "Team", "MP", "W", "D", "L", "P")
	if err != nil {
		return err
	}

	for _, team := range t.Rank() {
		_, err := fmt.Fprintf(w, lineFmt, team.name, team.played(), team.won, team.drawn, team.lost, team.points())
		if err != nil {
			return err
		}
	}
	return nil
}
