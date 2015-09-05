# reactLog
Package reactLog is reaction middleware for standard golang log.

[![Build Status](https://travis-ci.org/MatejB/reactLog.svg)](https://travis-ci.org/MatejB/reactLog) [![Code Coverage](http://gocover.io/_badge/github.com/MatejB/reactLog)](http://gocover.io/github.com/MatejB/reactLog) [![Documentation](https://godoc.org/github.com/MatejB/reactLog?status.svg)](https://godoc.org/github.com/MatejB/reactLog)

Basic usage:
<pre>
reactLogger := reactLog.New(os.Stderr)
reactLogger.AddReaction("INFO", &reactLog.Discard{})

log.SetOutput(reactLogger)

log.PrintLn("INFO this will not be written")
log.PrintLn("ERROR this will be written")
</pre>

reactLog concept is to filter and add additional functionality
to log messages based on trigger words.
If used in main package it enchance log globally with the use of log.SetOutput method.
Any number of trigger words can be registered using AddReaction
method each with it's own Reactor.

Reactor is the interface that wraps the Reaction method.
reactLog comes with few types that already implements Reactor interface:
* Discard for discarding log messages.
* Redirect to redirect log messages to other io.Writer.
* Copy to write log message both to underlying io.Writer and additional io.Writer.

See [Documentation](https://godoc.org/github.com/MatejB/reactLog) for more info.
