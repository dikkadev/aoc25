package day

import (
	"log/slog"
	"strings"

	"github.com/dikkadev/aoc25/days"
	"github.com/dikkadev/aoc25/input"
)

const DAY = 4

func init() {
	days.RegisterDay(DAY, Solve)
}

func Solve(input *input.Input, log *slog.Logger) int {
	result := 0
	fromMap := ParseMap(input)
	toMap := &Map{Width: fromMap.Width, Places: make([]Place, len(fromMap.Places))}
	slog.Debug("Parsed map", "map", fromMap)
	papersToRemove := true
	for papersToRemove {
		reachablePapers := 0
		for i, p := range fromMap.Places {
			toMap.Places[i] = p
			if p&JustRemoved != 0 {
				toMap.Places[i] = Empty
			}
			if p&Paper != 0 && fromMap.IsReachable(i) {
				toMap.Places[i] &^= Paper
				toMap.Places[i] |= JustRemoved
				reachablePapers++
			}
		}
		if reachablePapers == 0 {
			papersToRemove = false
		} else {
			slog.Debug("Removed reachable papers", "count", reachablePapers, "map", toMap)
			result += reachablePapers
			// Swap maps
			fromMap, toMap = toMap, fromMap
		}
	}
	slog.Debug("Final map", "map", fromMap)
	return result
}

type Place uint64

const (
	Empty       Place = 0
	JustRemoved Place = 1 << 62
	Paper       Place = 1 << 0
	Reachable   Place = 1 << 63
)

type Map struct {
	Places []Place
	Width  int
}

func ParseMap(input *input.Input) *Map {
	m := &Map{
		Places: make([]Place, 0),
		Width:  len(input.Lines()[0]),
	}
	for _, line := range input.Lines() {
		for _, char := range line {
			switch char {
			case '.':
				m.Places = append(m.Places, Place(Empty))
			case '@':
				m.Places = append(m.Places, Place(Paper))
			default:
				panic("unknown char")
			}
		}
	}
	return m
}

func (m *Map) IsReachable(idx int) bool {
	row, col := idx/m.Width, idx%m.Width
	h := len(m.Places) / m.Width

	count := 0
	check := func(r, c int) {
		if r < 0 || c < 0 || r >= h || c >= m.Width {
			return
		}
		if m.Places[r*m.Width+c]&Paper != 0 {
			count++
		}
	}

	check(row, col-1)
	check(row, col+1)
	check(row-1, col)
	check(row+1, col)
	check(row-1, col-1)
	check(row-1, col+1)
	check(row+1, col-1)
	check(row+1, col+1)

	// slog.Debug("Checking reachability", "index", idx, "row", row, "col", col, "paperNeighbors", count)

	return count < 4
}

func (m *Map) String() string {
	var sb strings.Builder
	sb.WriteRune('\n')
	for i, p := range m.Places {
		if i > 0 && i%m.Width == 0 {
			sb.WriteRune('\n')
		}
		switch {
		case p == Empty:
			sb.WriteRune('.')
		case p&JustRemoved != 0:
			sb.WriteRune('R')
		case p&Paper != 0 && p&Reachable != 0:
			sb.WriteRune('x')
		case p&Paper != 0:
			sb.WriteRune('@')
		default:
			sb.WriteRune('?')
		}
	}
	return sb.String()
}
