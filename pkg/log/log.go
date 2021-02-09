/*Package log is a package which meant to be used for logging the different levels of information
* to different output like (console, file, elastic, .. etc) this library could be  extended in the future
* to add more outputs to this library. this library meant to isolate our code from the logging place
* this package also provides the levels of log such as (debug, warning, info, error)
*
* Author: MNP
* Creation by: 8/12/2020
 */
package log

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

//Log is the interface which will be exposed to any Program that want to print logs
// any structure which provides new output method to this package should implement this interface
type log interface {
	debug(string, string, string)
	info(string, string, string)
	warning(string, string, string)
	error(string, string, string)
}

const (
	// CONSOLE is an enum for console output
	CONSOLE = iota
	// FILE is an enum for file output
	FILE
	// ELASTIC is an enum for elastic search output
	ELASTIC
)

var logger log = &console{}

// SetLogger is the package factory method to create a file output logging
// @arg filepath: the file path which the logger will use to send files
// @arg filename: the file name that the logger will output the files inside
// @return File: the structure which impements the interface that will be used to make logs
func SetLogger(logOut int) {

	switch logOut {
	case CONSOLE:
		logger = &console{}
	case FILE:
		logger = &file{
			filePath: os.Getenv("LOGFILE_PATH"),
			filename: os.Getenv("LOGFILE_NAME"),
		}
	case ELASTIC:
		logger = &elastic{
			url: os.Getenv("ELASTIC_URL"),
		}
	}
}

// Debug is the function which will output the Debuging message to the output logging interface
// @arg tag: a string which represents identification for the log message
// @arg message: the message which will be printed
// @arg a: the strings that will be formated into the message string
func Debug(tag string, message string, a ...interface{}) {

	// Extract Information and send it to the logger interface
	time, location, fullMessage := getLogInfo(tag, message, a...)
	logger.debug(time, location, fullMessage)

}

// Info is the function which will output the Information message to the output logging interface
// @arg tag: a string which represents identification for the log message
// @arg message: the message which will be printed
// @arg a: the strings that will be formated into the message string
func Info(tag string, message string, a ...interface{}) {

	// Extract Information and send it to the logger interface
	time, location, fullMessage := getLogInfo(tag, message, a...)
	logger.info(time, location, fullMessage)
}

// Warning is the function which will output the Warning message to the output logging interface
// @arg tag: a string which represents identification for the log message
// @arg message: the message which will be printed
// @arg a: the strings that will be formated into the message string
func Warning(tag string, message string, a ...interface{}) {

	// Extract Information and send it to the logger interface
	time, location, fullMessage := getLogInfo(tag, message, a...)
	logger.warning(time, location, fullMessage)
}

// Error is the function which will output the Error message to the output logging interface
// @arg tag: a string which represents identification for the log message
// @arg message: the message which will be printed
// @arg a: the strings that will be formated into the message string
func Error(tag string, message string, a ...interface{}) {

	// Extract Information and send it to the logger interface
	time, location, fullMessage := getLogInfo(tag, message, a...)
	logger.error(time, location, fullMessage)
}

// getLongInfo a helper function used to help the package interface methods
// to extract the information and get all the required information for better message representation
func getLogInfo(tag string, message string, a ...interface{}) (string, string, string) {

	// get the caller function filename and file line
	_, fileName, fileLine, ok := runtime.Caller(2) // the caller function back by 2 calls

	// check for the caller information
	if !ok {
		fileName = "Can't/be/resolved"
		fileLine = -1
	}

	// get the current time
	currentTime := time.Now()

	// compose the full message by merging the tag with the message
	fullMessage := fmt.Sprintf(message, a...)
	fullMessage = fmt.Sprintf("[%s] %s", tag, fullMessage)

	// formate the location of the error (filename, fileline)
	location := fmt.Sprintf("%s:%d", fileName, fileLine)

	// formate the current time in a human readable formate
	time := currentTime.Format("2006-01-02 15:04:05")

	return time, location, fullMessage
}
