package day

import (
	"fmt"
	"log/slog"

	"github.com/dikkadev/aoc25/days"
	"github.com/dikkadev/aoc25/input"
)

const DAY = 3

func init() {
	days.RegisterDay(DAY, Solve)
}

func Solve(input *input.Input, log *slog.Logger) int {
	result := 0
	banks := make([]Bank, 0)
	for l := range input.AugmentedLineStream() {
		if len(l.T) == 0 {
			continue
		}
		bank := ParseBank(l.T)
		banks = append(banks, bank)
		// slog.Debug("Parsed bank", "bank", bank)
	}

	for _, bank := range banks {
		maxPair := bank.MaxJoltagePair()
		result += int(maxPair)
		slog.Debug("Bank max joltage pair", "bank", bank, "max_pair", maxPair)
	}

	return result
}

type Joltage int

func Combine2Joltages(first, second Joltage) Joltage {
	firstStr := fmt.Sprintf("%d", first)
	secondStr := fmt.Sprintf("%d", second)
	combinedStr := firstStr + secondStr
	var combined Joltage
	fmt.Sscanf(combinedStr, "%d", &combined)
	return combined
}

func digits(n Joltage) int {
	if n == 0 {
		return 1
	}
	d := 0
	for n > 0 {
		n /= 10
		d++
	}
	return d
}

func CombineNJoltages(joltages ...Joltage) Joltage {
	var out Joltage
	for _, j := range joltages {
		mul := Joltage(1)
		for i := 0; i < digits(j); i++ {
			mul *= 10
		}
		out = out*mul + j
	}
	return out
}

type Bank struct {
	Batteries []Joltage
}

func ParseBank(inp string) Bank {
	batteries := make([]Joltage, 0)
	for _, c := range inp {
		batteries = append(batteries, Joltage(c-'0'))
	}
	return Bank{Batteries: batteries}
}

func (b *Bank) MaxJoltagePair() Joltage {
	// Greedy approach: for each of the 12 positions, pick the largest digit
	// available such that enough digits remain to fill the remaining positions.
	//
	// If we need to pick digit at position `pos` (0-11), and our current index
	// in the array is `start`, we can pick any index from `start` to
	// `n - (12 - pos - 1) - 1 = n - 12 + pos` (inclusive), because we need
	// (12 - pos - 1) more digits after this one.

	const toSelect = 12
	n := len(b.Batteries)

	selected := make([]Joltage, 0, toSelect)
	start := 0

	for pos := 0; pos < toSelect; pos++ {
		remaining := toSelect - pos - 1 // digits we still need after this one
		maxIdx := n - remaining - 1     // last valid index we can pick from

		// Find the maximum digit in range [start, maxIdx]
		bestIdx := start
		bestVal := b.Batteries[start]
		for i := start + 1; i <= maxIdx; i++ {
			if b.Batteries[i] > bestVal {
				bestVal = b.Batteries[i]
				bestIdx = i
			}
		}

		selected = append(selected, bestVal)
		start = bestIdx + 1 // next digit must come after this one
	}

	return CombineNJoltages(selected...)
}
