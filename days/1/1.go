package day

import (
	"fmt"
	"log/slog"

	"github.com/dikkadev/aoc25/days"
	"github.com/dikkadev/aoc25/input"
)

const DAY = 1

func init() {
	days.RegisterDay(DAY, Solve)
}

func Solve(input *input.Input, log *slog.Logger) int {
	result := 0
	lock := Lock{State: 50}
	for l := range input.AugmentedLineStream() {
		if len(l.T) == 0 {
			continue
		}
		line := ParseLine(l.T)
		lock.Turn(line.Dir, line.Steps)
		if lock.State == 0 {
			result++
		}
	}

	return result
}

type Line struct {
	Dir   Direction
	Steps int
}

func ParseLine(inp string) Line {
	firstChar := inp[0]
	var dir Direction
	if firstChar == 'L' {
		dir = LEFT
	} else {
		dir = RIGHT
	}

	var steps int
	fmt.Sscanf(inp[1:], "%d", &steps)
	return Line{Dir: dir, Steps: steps}
}

type Lock struct {
	State int
}

type Direction int

const (
	LEFT Direction = iota
	RIGHT
)

func (d Direction) String() string {
	switch d {
	case LEFT:
		return "LEFT"
	case RIGHT:
		return "RIGHT"
	default:
		return "UNKNOWN"
	}
}

func (l *Lock) Turn(dir Direction, steps int) {
	currState := l.State
	switch dir {
	case LEFT:
		for range steps {
			if currState == 0 {
				currState = 99
			} else {
				currState--
			}
		}
	case RIGHT:
		for range steps {
			if currState == 99 {
				currState = 0
			} else {
				currState++
			}
		}
	}

	l.State = currState
	slog.Debug("Turned", "direction", dir, "steps", steps, "newState", l.State)
}
