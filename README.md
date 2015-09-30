# reactLog
Package reactLog is reaction middle-ware for standard golang log.

[![Build Status](https://travis-ci.org/MatejB/reactLog.svg)](https://travis-ci.org/MatejB/reactLog) [![Code Coverage](http://gocover.io/_badge/github.com/MatejB/reactLog)](http://gocover.io/github.com/MatejB/reactLog) [![Documentation](https://godoc.org/github.com/MatejB/reactLog?status.svg)](https://godoc.org/github.com/MatejB/reactLog)

Basic usage:
<pre>
reactLogger := reactLog.New(os.Stderr)

copyBuf := &bytes.Buffer{}
reactLogger.AddReaction("user ID 85", &reactLog.Copy{copyBuf})

log.SetOutput(reactLogger)

log.Println("This is regular log message")
log.Println("This error message concers user ID 85 and will be copied to copyBuf.")
</pre>

reactLog concept is to filter and add additional functionality to log
based on log message content.

If used in main package it enhance log globally with the use of log.SetOutput method.

Any number of trigger words can be registered using AddReaction
method each with it's own Reactor.

Reactor is the interface that wraps the Reaction method.
reactLog comes with few types that already implements Reactor interface:
* Discard for discarding log messages.
* Redirect to redirect log messages to other io.Writer.
* Copy to write log message both to underlying io.Writer and additional io.Writer.

Feel free to create Reactors for you specific use case by implementing Reactor interface.

See [Documentation](https://godoc.org/github.com/MatejB/reactLog) for more info.
