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

func CombineJoltages(first, second Joltage) Joltage {
	firstStr := fmt.Sprintf("%d", first)
	secondStr := fmt.Sprintf("%d", second)
	combinedStr := firstStr + secondStr
	var combined Joltage
	fmt.Sscanf(combinedStr, "%d", &combined)
	return combined
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
	biggest := Joltage(-1)
	for i := range len(b.Batteries) {
		first := b.Batteries[i]
		for j := i + 1; j < len(b.Batteries); j++ {
			second := b.Batteries[j]
			combined := CombineJoltages(first, second)
			if combined > biggest {
				biggest = combined
			}
		}
	}
	return biggest
}
