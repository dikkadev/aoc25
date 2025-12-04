package day

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/dikkadev/aoc25/days"
	"github.com/dikkadev/aoc25/input"
)

const DAY = 2

func init() {
	days.RegisterDay(DAY, Solve)
}

func Solve(input *input.Input, log *slog.Logger) int {
	result := 0
	ranges := make([]IdRange, 0)
	for l := range input.AugmentedLineStream() {
		if len(l.T) == 0 {
			continue
		}
		rangesInps := strings.Split(l.T, ",")
		for _, inp := range rangesInps {
			r := ParseRange(inp)
			ranges = append(ranges, r)
			// slog.Debug("Parsed range", "range", r)
		}
	}
	for _, r := range ranges {
		slog.Debug("Processing range", "range", r)
		for i := r.Start; i <= r.End; i++ {
			if IsNumberStupid(i) {
				result = result + i
				slog.Debug("Found stupid number", "number", i)
			}
		}
	}
	return result
}

type IdRange struct {
	Start int
	End   int
}

func ParseRange(inp string) IdRange {
	var start, end int
	fmt.Sscanf(inp, "%d-%d", &start, &end)
	return IdRange{Start: start, End: end}
}

func IsNumberStupid(nInt int) bool {
	n := fmt.Sprintf("%d", nInt)
	if len(n)%2 != 0 {
		return false
	}
	halfSize := len(n) / 2

	// for length := startLength; length <= halfSize; length++ {

	length := halfSize
	if len(n)%length != 0 {
		// continue
		return false
	}
	subs := make([]string, 0)
	for i := 0; i < len(n); i += length {
		subs = append(subs, n[i:i+length])
	}
	slog.Debug("Checking substrings", "number", n, "length", length, "subs", subs)
	allSame := true
	firstSub := subs[0]
	for _, sub := range subs[1:] {
		if sub != firstSub {
			allSame = false
			break
		}
	}
	if allSame {
		return true
	}
	// }

	return false
}
