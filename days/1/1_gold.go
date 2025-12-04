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
		timesAtZero := lock.Turn(line.Dir, line.Steps)
		result += timesAtZero
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

func (l *Lock) Turn(dir Direction, steps int) int {
	currState := l.State
	timesAtZero := 0
	switch dir {
	case LEFT:
		for range steps {
			currState--
			if currState < 0 {
				currState = 99
			}
			if currState == 0 {
				timesAtZero++
			}
		}
	case RIGHT:
		for range steps {
			currState++
			if currState > 99 {
				currState = 0
			}
			if currState == 0 {
				timesAtZero++
			}
		}
	}

	l.State = currState
	slog.Debug("Turned", "direction", dir, "steps", steps, "newState", l.State, "timesAtZero", timesAtZero)

	return timesAtZero
}
