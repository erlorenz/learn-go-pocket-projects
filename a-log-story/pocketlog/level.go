package pocketlog

// Level represents an available logging level.
type Level byte

const (
	// LevelDebug reprents the lowest level of log, mostly for debugging purpses.
	LevelDebug Level = iota
	// LevelInfo represents a logging level that contains information deemed valuable.
	LevelInfo
	// LevelError represents the highest logging level, only to be used to trace errors.
	LevelError
)
