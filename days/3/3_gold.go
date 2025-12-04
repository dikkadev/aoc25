package day

import (
	"fmt"
	"log/slog"
	"sync"

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

	// for _, bank := range banks {
	// 	maxPair := bank.MaxJoltagePair()
	// 	result += int(maxPair)
	// 	slog.Debug("Bank max joltage pair", "bank", bank, "max_pair", maxPair)
	// }
	// parallelize above

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	for _, bank := range banks {
		wg.Add(1)
		go func(bank Bank) {
			defer wg.Done()
			slog.Debug("Processing bank", "bank", bank)
			maxPair := bank.MaxJoltagePair()
			mu.Lock()
			result += int(maxPair)
			mu.Unlock()
			slog.Debug("Bank max joltage pair", "bank", bank, "max_pair", maxPair)
		}(bank)
	}
	wg.Wait()

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
	biggest := Joltage(-1)
	//12 deep
	for i := range len(b.Batteries) {
		first := b.Batteries[i]
		for second := i + 1; second < len(b.Batteries); second++ {
			second := b.Batteries[second]
			for third := second + 1; third < Joltage(len(b.Batteries)); third++ {
				third := b.Batteries[third]
				for fourth := third + 1; fourth < Joltage(len(b.Batteries)); fourth++ {
					fourth := b.Batteries[fourth]
					for fifth := fourth + 1; fifth < Joltage(len(b.Batteries)); fifth++ {
						fifth := b.Batteries[fifth]
						for sixth := fifth + 1; sixth < Joltage(len(b.Batteries)); sixth++ {
							sixth := b.Batteries[sixth]
							for seventh := sixth + 1; seventh < Joltage(len(b.Batteries)); seventh++ {
								seventh := b.Batteries[seventh]
								for eighth := seventh + 1; eighth < Joltage(len(b.Batteries)); eighth++ {
									eighth := b.Batteries[eighth]
									for ninth := eighth + 1; ninth < Joltage(len(b.Batteries)); ninth++ {
										ninth := b.Batteries[ninth]
										for tenth := ninth + 1; tenth < Joltage(len(b.Batteries)); tenth++ {
											tenth := b.Batteries[tenth]
											for eleventh := tenth + 1; eleventh < Joltage(len(b.Batteries)); eleventh++ {
												eleventh := b.Batteries[eleventh]
												for twelfth := eleventh + 1; twelfth < Joltage(len(b.Batteries)); twelfth++ {
													twelfth := b.Batteries[twelfth]
													// combined := CombineJoltages(first, CombineJoltages(second, CombineJoltages(third, CombineJoltages(fourth, CombineJoltages(fifth, CombineJoltages(sixth, CombineJoltages(seventh, CombineJoltages(eighth, CombineJoltages(ninth, CombineJoltages(tenth, CombineJoltages(eleventh, twelfth)))))))))))
													combined := CombineNJoltages(first, second, third, fourth, fifth, sixth, seventh, eighth, ninth, tenth, eleventh, twelfth)
													// slog.Debug("Checking combination", "combination", combined)
													if combined > biggest {
														biggest = combined
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	// combined := CombineJoltages(first, second)
	// if combined > biggest {
	// 	biggest = combined
	// }
	return biggest
}
