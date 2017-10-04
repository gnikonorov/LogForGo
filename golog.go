/*

	Log4Go Like logger for the go programming language

	Structure initialization:
		Appender_Mode: Either DEBUG, SCREEN, or BOTH
		Appender_File:

	Author: Gleb Nikonorov
*/

package golog

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

const (
	// Logging modes for the logger
	fatal = "FATAL"   // Non-recoverable error.
	err   = "ERROR"   // A recoverable error
	warn  = "WARNING" // Indicator of potential problems
	info  = "INFO"    // Non severe log information. Should be used for things like user input
	debug = "DEBUG"   // This mode should be used debug information

	// File indicates information will be outputted to a log file
	File = 1

	// Screen indicates information will be outputted to the screen
	Screen = 2

	// Both indicates information will be outted to both file and screen
	Both = 3

	// These are constants for terminal colors
	colorFatal = "\x1B[34" // This is blue
	colorErr   = "\x1B[31" // This is red
	colorWarn  = "\x1B[33" // This is yellow
	colorDebug = "\x1B[32" // This is green
	// Info is left in the native terminal color
)

// Logger is representative of the logger for use in other go programs
// Contains the following fields:
//	loggingMode: Must be either File, Screen, or Both
//	loggingDirectory: The directory to store log files in. Must be a valid directory if loggingMode is Screen or Both
//	loggingFile: The name of the log file to output to. Must be a valid file name of loggingMode is Screen or Both
//
// The following methods are exposed by this structure:
//	Debug(logText string): Log debug output to log destination
//	Info(logText string): Log info output to log destination
//	Warning(logText string): Log warning output to log destination
//	Err(logText string): Log error output to log destination
//	Fatal(logText string): Log fatal output to log destination
//	Is_Uninitialized: Returns true if this structure has not been allocated
type Logger struct {
	loggingMode      int    // The mode of the logger (Should be FILE, SCREEN, or BOTH)
	loggingDirectory string // The directory to store logs in
	loggingFile      string // The file to store logs in
}

// Debug Outputs debug log information to the logging destination
func (logger *Logger) Debug(logText string) {
	if logger.loggingMode == Screen || logger.loggingMode == Both {
		fmt.Printf("[%s] %s: %s\n", time.Now().String(), debug, logText)
	}

	if logger.loggingMode == Both || logger.loggingMode == File {
		var fileName = logger.loggingDirectory + "/" + logger.loggingFile
		var writeBytes = []byte("[" + time.Now().String() + "] " + debug + ":" + logText)
		ioutil.WriteFile(fileName, writeBytes, 0644)
	}
}

// Info Outputs info log information to the logging destination
func (logger *Logger) Info(logText string) {
	if logger.loggingMode == Screen || logger.loggingMode == Both {
		fmt.Printf("[%s] %s: %s\n", time.Now().String(), info, logText)
	}

	if logger.loggingMode == Both || logger.loggingMode == File {
		var fileName = logger.loggingDirectory + "/" + logger.loggingFile
		var writeBytes = []byte("[" + time.Now().String() + "] " + info + ":" + logText)
		ioutil.WriteFile(fileName, writeBytes, 0644)
	}
}

// Warning Outputs warning information to the logging destination
func (logger *Logger) Warning(logText string) {
	if logger.loggingMode == Screen || logger.loggingMode == Both {
		fmt.Printf("[%s] %s: %s\n", time.Now().String(), warn, logText)
	}

	if logger.loggingMode == Both || logger.loggingMode == File {
		var fileName = logger.loggingDirectory + "/" + logger.loggingFile
		var writeBytes = []byte("[" + time.Now().String() + "] " + warn + ":" + logText)
		ioutil.WriteFile(fileName, writeBytes, 0644)
	}
}

// Err Outputs error information to the logging destination
func (logger *Logger) Err(logText string) {
	if logger.loggingMode == Screen || logger.loggingMode == Both {
		fmt.Printf("[%s] %s: %s\n", time.Now().String(), err, logText)
	}

	if logger.loggingMode == Both || logger.loggingMode == File {
		var fileName = logger.loggingDirectory + "/" + logger.loggingFile
		var writeBytes = []byte("[" + time.Now().String() + "] " + err + ":" + logText)
		ioutil.WriteFile(fileName, writeBytes, 0644)
	}
}

// Fatal Outputs fatal information to the logging desination
func (logger *Logger) Fatal(logText string) {
	if logger.loggingMode == Screen || logger.loggingMode == Both {
		fmt.Printf("[%s] %s: %s\n", time.Now().String(), fatal, logText)
	}

	if logger.loggingMode == Both || logger.loggingMode == File {
		var fileName = logger.loggingDirectory + "/" + logger.loggingFile
		var writeBytes = []byte("[" + time.Now().String() + "] " + fatal + ":" + logText)
		ioutil.WriteFile(fileName, writeBytes, 0644)
	}
}

// IsUninitialized Returns true if this structure has not yet been allocated
func (logger *Logger) IsUninitialized() bool {
	return logger.loggingMode == 0
}

// SetupLoggerWithFilename Sets up and returns a logger instance.
func SetupLoggerWithFilename(logMode int, logDirectory string, logFile string) Logger {
	// Validate parmaters
	if logMode != File && logMode != Screen && logMode != Both {
		panic("Log mode must either be File, Screen, or Both. Goodbye")
	}

	if logMode == File || logMode == Both {
		// We're logging to a file, make sure that the directory given to us was valid
		fileInfo, err := os.Stat(logDirectory)
		if err != nil {
			if os.IsNotExist(err) {
				// The file does not exist
				panic("Please provide a valid log directory. Goodbye.")
			}
		}

		// Check to make sure we actually gave a directory
		if !fileInfo.IsDir() {
			panic("You must give a directory! Not a file!")
		}
	}

	logger := Logger{loggingMode: logMode, loggingDirectory: logDirectory, loggingFile: logFile}
	return logger
}
