package day

import (
	"fmt"
	"log/slog"
	"sort"

	"github.com/dikkadev/aoc25/days"
	"github.com/dikkadev/aoc25/input"
)

const DAY = 5

func init() {
	days.RegisterDay(DAY, Solve)
}

func Solve(input *input.Input, log *slog.Logger) int {
	result := 0
	freshRanges := []FreshRange{}

	for l := range input.AugmentedLineStream() {
		if len(l.T) == 0 {
			break
		}
		var start, end int
		_, err := fmt.Sscanf(l.T, "%d-%d", &start, &end)
		if err != nil {
			log.Error("Failed to parse range", "line", l.T, "error", err)
			continue
		}
		freshRanges = append(freshRanges, FreshRange{Start: start, End: end})

	}
	slog.Debug("Parsed input", "freshRanges", freshRanges)

	// freshIds := make(map[int]struct{})
	// for _, fr := range freshRanges {
	// 	slog.Debug("Adding fresh IDs from range", "range", fr)
	// 	for id := fr.Start; id <= fr.End; id++ {
	// 		freshIds[id] = struct{}{}
	// 	}
	// }
	// result = len(freshIds)

	changes := make([]Change, 0)
	for _, fr := range freshRanges {
		changes = append(changes, Change{At: fr.Start, Up: true})
		changes = append(changes, Change{At: fr.End + 1, Up: false})
	}
	//sort
	sort.Slice(changes, func(i, j int) bool {
		if changes[i].At == changes[j].At {
			return changes[i].Up && !changes[j].Up
		}
		return changes[i].At < changes[j].At
	})
	slog.Debug("Computed changes", "changes", changes)

	x := changes[0].At
	currentlyInsideFresh := 1

	for _, ch := range changes[1:] {
		if currentlyInsideFresh > 0 {
			result += ch.At - x
		}
		x = ch.At

		if ch.Up {
			currentlyInsideFresh++
		} else {
			currentlyInsideFresh--
		}

		slog.Debug("Processed change", "change", ch, "currentlyInsideFresh", currentlyInsideFresh, "result", result)
	}

	return result
}

type FreshRange struct {
	Start int
	End   int
}

type Change struct {
	At int
	Up bool
}

func (c Change) String() string {
	if c.Up {
		return fmt.Sprintf("+@%d", c.At)
	} else {
		return fmt.Sprintf("-@%d", c.At)
	}
}
