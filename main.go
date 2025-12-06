package main

import (
	"flag"
	"log/slog"
	"os"
	"os/signal"
	"runtime/pprof"
	"syscall"

	_ "net/http/pprof"

	"github.com/dikkadev/aoc25/days"
	_ "github.com/dikkadev/aoc25/days/6"
	_ "github.com/dikkadev/aoc25/days/5"
	_ "github.com/dikkadev/aoc25/days/6"
	_ "github.com/dikkadev/aoc25/days/4"
	_ "github.com/dikkadev/aoc25/days/6"
	_ "github.com/dikkadev/aoc25/days/5"
	_ "github.com/dikkadev/aoc25/days/6"
	_ "github.com/dikkadev/aoc25/days/1"
	_ "github.com/dikkadev/aoc25/days/6"
	_ "github.com/dikkadev/aoc25/days/5"
	_ "github.com/dikkadev/aoc25/days/6"
	_ "github.com/dikkadev/aoc25/days/4"
	_ "github.com/dikkadev/aoc25/days/6"
	_ "github.com/dikkadev/aoc25/days/5"
	_ "github.com/dikkadev/aoc25/days/6"
	_ "github.com/dikkadev/aoc25/days/2"
	_ "github.com/dikkadev/aoc25/days/6"
	_ "github.com/dikkadev/aoc25/days/5"
	_ "github.com/dikkadev/aoc25/days/6"
	_ "github.com/dikkadev/aoc25/days/4"
	_ "github.com/dikkadev/aoc25/days/6"
	_ "github.com/dikkadev/aoc25/days/5"
	_ "github.com/dikkadev/aoc25/days/6"
	_ "github.com/dikkadev/aoc25/days/3"
	_ "github.com/dikkadev/aoc25/days/6"
	_ "github.com/dikkadev/aoc25/days/5"
	_ "github.com/dikkadev/aoc25/days/6"
	_ "github.com/dikkadev/aoc25/days/4"
	_ "github.com/dikkadev/aoc25/days/6"
	_ "github.com/dikkadev/aoc25/days/5"
	_ "github.com/dikkadev/aoc25/days/6"
	"github.com/dikkadev/prettyslog"
)

var (
	small     bool
	dayNumber uint
	verbose   bool
)

func main() {
	flag.BoolVar(&small, "s", false, "Use small input")
	flag.UintVar(&dayNumber, "d", 0, "Day to run")
	flag.BoolVar(&verbose, "v", false, "Enable debug level logging")
	flag.Parse()

	logLevel := slog.LevelInfo
	if verbose {
		logLevel = slog.LevelDebug
	}

	handler := prettyslog.NewPrettyslogHandler("AOC", prettyslog.WithLevel(logLevel))
	slog.SetDefault(slog.New(handler))

	if verbose {
		for _, d := range days.Days {
			if d != nil {
				d.SetLogger(slog.New(handler))
			}
		}
	}

	slog.Info("Starting")

	f, err := os.Create("cpu.prof")
	if err != nil {
		slog.Error("Could not create CPU profile", "error", err)
		return
	}
	defer f.Close()
	if err := pprof.StartCPUProfile(f); err != nil {
		slog.Error("Could not start CPU profile", "error", err)
		return
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		slog.Info("Caught interrupt, stopping CPU profile...")
		pprof.StopCPUProfile()
		f.Close()
		os.Exit(0)
	}()
	defer pprof.StopCPUProfile()

	day := days.Days[dayNumber]
	if day == nil {
		slog.Error("Day not found", "day", dayNumber)
		return
	}

	err = day.PrepeareInputs()
	if err != nil {
		slog.Error("Failed to prepare inputs", "error", err)
		return
	}

	slog.Info("Solving", "day", dayNumber, "small", small)

	result := day.Solve(small)

	slog.Info("Result", "result", result)

}
