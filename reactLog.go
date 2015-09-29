/*
Package reactLog is reaction middleware for standard log.

Basic usage:
	reactLogger := reactLog.New(os.Stderr)

    copyBuf := &bytes.Buffer{}
	reactLogger.AddReaction("user ID 85", &reactLog.Copy{copyBuf})

	log.SetOutput(reactLogger)

	log.PrintLn("This is regular log message")
	log.PrintLn("This error message concers user ID 85 and will be copied to copyBuf.")

reactLog concept is to filter and add additional functionality to log
based on log message content.
If used in main package it enhance log globally with the use of log.SetOutput method.
Any number of trigger words can be registered using AddReaction
method each with it's own Reactor.

Reactor is the interface that wraps the Reaction method.
reactLog comes with few types that already implements Reactor interface:
 Discard for discarding log messages.
 Redirect to redirect log messages to other io.Writer.
 Copy to write log message both to underlying io.Writer and additional io.Writer.
Feel free to create Reactors for you specific use case by implementing Reactor interface.

See Examples for more info.
*/
package reactLog

import (
	"bytes"
	"io"
)

// A Logger is logging object that passes writes to given
// io.Writer if no appropriate reaction is found or reaction
// returns true.
type Logger struct {
	out      io.Writer
	reactors map[string]Reactor
}

// New creates a new Logger. Pass-through io.Writer
// must be given, os.Stderr in most cases.
func New(out io.Writer) *Logger {
	return &Logger{out: out, reactors: make(map[string]Reactor)}
}

func (l *Logger) Write(p []byte) (n int, err error) {
allTriggers:
	for key := range l.reactors {
		trigger := []byte(key)
		if len(trigger) > len(p) {
			continue allTriggers
		}

		// find start
		maxPos := len(p) - len(trigger) + 2
		for i := 0; i < maxPos; i = i + 1 {
			if p[i] == trigger[0] {
				if bytes.Equal(p[i:i+len(trigger)], trigger) {
					output, err := l.reactors[key].Reaction(p)
					if err != nil {
						return 0, err
					}
					if output {
						n, err = l.out.Write(p)
					}
					return n, nil
				}
			}
		}
	}

	// default behaviour
	n, err = l.out.Write(p)

	return
}

// Reactor is the interface that wraps the Reaction method.
//
// Reaction decides if logLine is to be written to underlying
// io.Writer object by returning true.
type Reactor interface {
	Reaction(logLine []byte) (passOut bool, err error)
}

// AddReaction adds reaction to be executed when
// trigger is encountered in log line.
func (l *Logger) AddReaction(trigger string, reaction Reactor) {
	l.reactors[trigger] = reaction
}

// Discard type implements Reactor with discard functionality.
type Discard struct{}

// Reaction is to disacard log if trigger is found
func (d *Discard) Reaction(logLine []byte) (passOut bool, err error) {
	return false, nil
}

// Redirect type implements Reactor with redirection functionality.
// It will redirect log output to given io.Writer
type Redirect struct {
	Out io.Writer
}

// Reaction is to redirect log to other writer if trigger is found
func (r *Redirect) Reaction(logLine []byte) (passOut bool, err error) {
	_, err = r.Out.Write(logLine)
	return false, err
}

// Copy type implements Reactor with copy functionality.
// It will copy log to given io.Writer.
type Copy struct {
	Out io.Writer
}

// Reaction is to copy log to other writer if trigger is found
func (c *Copy) Reaction(logLine []byte) (passOut bool, err error) {
	_, err = c.Out.Write(logLine)
	return true, err
}
