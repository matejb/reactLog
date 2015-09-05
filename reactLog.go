/*
Package reactLog is reaction middleware for standard log.

Typical usage:
	reactLogger := reactLog.New(os.Stderr)
	reactLogger.AddReaction("INFO", &reactLog.Discard{})

	log.SetOutput(reactLogger)

	log.PrintLn("INFO this will not be written")
	log.PrintLn("ERROR this will be written")

*/
package reactLog

import (
	"bufio"
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
	analyseBuf := &bytes.Buffer{}
	n, err = analyseBuf.Write(p)
	if err != nil {
		return
	}

	scanner := bufio.NewScanner(analyseBuf)

	scanner.Split(bufio.ScanWords)

	word := ""
	for scanner.Scan() {
		word = scanner.Text()
		if reactor, ok := l.reactors[word]; ok {
			output, err := reactor.Reaction(p)
			if err != nil {
				return 0, err
			}
			if output {
				n, err = l.out.Write(p)
			}
			return n, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return n, err
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

// AddReaction add's reaction to be executed when trigger word
// is encounterd in log line.
func (l *Logger) AddReaction(triggerWord string, reaction Reactor) {
	l.reactors[triggerWord] = reaction
}

// Discard type implements Reactor with discard functionality.
type Discard struct{}

func (d *Discard) Reaction(logLine []byte) (passOut bool, err error) {
	return false, nil
}

// Redirect type implements Reactor with redirection functionality.
// It will redirect log output to given io.Writer
type Redirect struct {
	Out io.Writer
}

func (r *Redirect) Reaction(logLine []byte) (passOut bool, err error) {
	_, err = r.Out.Write(logLine)
	return false, err
}

// Copy type implements Reactor with copy functionality.
// It will copy log to given io.Writer.
type Copy struct {
	Out io.Writer
}

func (c *Copy) Reaction(logLine []byte) (passOut bool, err error) {
	_, err = c.Out.Write(logLine)
	return true, err
}
