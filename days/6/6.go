package day

import (
	"log/slog"
	"strconv"
	"strings"

	"github.com/dikkadev/aoc25/days"
	"github.com/dikkadev/aoc25/input"
)

const DAY = 6

func init() {
	days.RegisterDay(DAY, Solve)
}

func Solve(input *input.Input, log *slog.Logger) int {
	result := 0
	rows := make([][]int, 0)
	ops := make([]Operation, 0)
	for l := range input.AugmentedLineStream() {
		if len(l.T) == 0 {
			continue
		}
		split := strings.Split(l.T, " ")
		if split[0] == "+" || split[0] == "*" {
			for _, s := range split {
				if s != "" {
					switch s {
					case "+":
						ops = append(ops, Add)
					case "*":
						ops = append(ops, Mult)
					}
				}
			}
		} else {
			row := make([]int, 0)
			for _, s := range split {
				if s != "" {
					asNumber, err := strconv.Atoi(s)
					if err != nil {
						panic("invalid number")
					}
					row = append(row, asNumber)
				}
			}
			rows = append(rows, row)
		}
	}
	slog.Debug("Parsed input", "rows", rows, "ops", ops)

	rowLength := len(rows[0])
	for _, row := range rows {
		if len(row) != rowLength {
			panic("inconsistent row lengths")
		}
	}
	if len(ops) != rowLength {
		panic("inconsistent row lengths")
	}

	problems := make([]Problem, 0)
	for i := range rowLength {
		values := make([]int, 0)
		for _, row := range rows {
			values = append(values, row[i])
		}
		problems = append(problems, Problem{
			Values: values,
			Op:     ops[i],
		})
	}

	slog.Debug("Constructed problems", "problems", problems)

	for _, p := range problems {
		r := p.Solve()
		slog.Debug("Solved problem", "problem", p, "result", r)
		result += r
	}

	return result
}

type Operation int

const (
	Add Operation = iota
	Mult
)

func (op Operation) String() string {
	switch op {
	case Add:
		return "+"
	case Mult:
		return "*"
	default:
		return "?"
	}
}

type Problem struct {
	Values []int
	Op     Operation
}

func (p *Problem) Solve() int {
	result := 0
	if p.Op == Mult {
		result = 1
	}
	for _, v := range p.Values {
		switch p.Op {
		case Add:
			result += v
		case Mult:
			result *= v
		}
	}

	return result
}
