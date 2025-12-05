package day

import (
	"fmt"
	"log/slog"

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
	availableIDs := make([]int, 0)

	parsingRange := true
	for l := range input.AugmentedLineStream() {
		if len(l.T) == 0 {
			parsingRange = false
			continue
		}
		if parsingRange {
			var start, end int
			_, err := fmt.Sscanf(l.T, "%d-%d", &start, &end)
			if err != nil {
				log.Error("Failed to parse range", "line", l.T, "error", err)
				continue
			}
			freshRanges = append(freshRanges, FreshRange{Start: start, End: end})
		} else {
			var id int
			_, err := fmt.Sscanf(l.T, "%d", &id)
			if err != nil {
				log.Error("Failed to parse ID", "line", l.T, "error", err)
				continue
			}
			availableIDs = append(availableIDs, id)
		}

	}
	slog.Debug("Parsed input", "freshRanges", freshRanges, "availableIDs", availableIDs)

	for _, id := range availableIDs {
		for _, fr := range freshRanges {
			if fr.IsInside(id) {
				slog.Debug("ID is inside fresh range", "id", id, "range", fr)
				result++
				break
			}
		}
	}
	return result
}

type FreshRange struct {
	Start int
	End   int
}

func (fr FreshRange) IsInside(id int) bool {
	return id >= fr.Start && id <= fr.End
}
