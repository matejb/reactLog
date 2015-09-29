package reactLog_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/MatejB/reactLog"
)

func ExampleDiscard() {
	reactLogger := reactLog.New(os.Stdout) // use os.Stderr for default log functionality
	reactLogger.AddReaction("INFO", &reactLog.Discard{})

	logger := log.New(reactLogger, "", 0)
	logger.Println("INFO this will NOT be written")
	logger.Println("ERROR this will be written")
	// Output: ERROR this will be written
}

func ExampleRedirect() {
	logContainerForUser107 := &bytes.Buffer{} // it can be any io.Writer eq. file

	reactLogger := reactLog.New(ioutil.Discard) // use os.Stderr for default log functionality
	reactLogger.AddReaction("user ID 107", &reactLog.Redirect{logContainerForUser107})

	logger := log.New(reactLogger, "", 0)

	logger.Println("INFO dummy log 1")
	logger.Println("ERROR dummy log 2")

	logger.Println("Log concerning user ID 107 with extra data")

	logger.Println("INFO dummy log 3")
	logger.Println("ERROR dummy log 4")

	fmt.Println(logContainerForUser107)
	// Output: Log concerning user ID 107 with extra data
}

func ExampleCopy() {
	logContainerForUser107 := &bytes.Buffer{} // it can be any io.Writer eq. file

	reactLogger := reactLog.New(os.Stdout) // use os.Stderr for default log functionality
	reactLogger.AddReaction("user ID 107", &reactLog.Copy{logContainerForUser107})

	logger := log.New(reactLogger, "", 0)

	logger.Println("Log concerning user ID 107 with extra data")
	logger.Println("INFO dummy log")

	fmt.Println(logContainerForUser107) // in logContainerForUser107 is copy of log just for USER_107
	// Output:
	// Log concerning user ID 107 with extra data
	// INFO dummy log
	// Log concerning user ID 107 with extra data
}
