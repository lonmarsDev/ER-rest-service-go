/*Package log is a package which meant to be used for logging the different levels of information
* to different output like (console, file, elastic, .. etc) this library could be  extended in the future
* to add more outputs to this library. this library meant to isolate our code from the logging place
* this package also provides the levels of log such as (debug, warning, info, error)
*
* Author: MNP
* Creation by: 8/12/2020
 */
package log

import "fmt"

// File is the type which handle the console logging functionality
// this type shall implement the Log interface
type file struct {
	filePath string
	filename string
}

// Debug is the function which will output the Debuging message to the output logging interface
// @arg tag: a string which represents identification for the log message
// @arg location: the location of the error in the project
// @arg message: the message which will be printed
func (f *file) debug(time string, location string, message string) {
	if !f.isValidFileProperities() {
		return
	}

	// TODO: replace & implement writing a debug log in file
	fmt.Printf(formatter("DEBUG", time, location, message) + "\n")

}

// Info is the function which will output the Information message to the output logging interface
// @arg tag: a string which represents identification for the log message
// @arg location: the location of the error in the project
// @arg message: the message which will be printed
func (f *file) info(time string, location string, message string) {
	if !f.isValidFileProperities() {
		return
	}

	// TODO: replace & implement writing a Info in the file
	fmt.Printf("%s %s %s\n", blue, formatter("INFO", time, location, message), reset)
}

// Warning is the function which will output the Warning message to the output logging interface
// @arg tag: a string which represents identification for the log message
// @arg location: the location of the error in the project
// @arg message: the message which will be printed
func (f *file) warning(time string, location string, message string) {
	if !f.isValidFileProperities() {
		return
	}
	// TODO: replace & implement writing a Warining in the file
	fmt.Printf("%s %s %s\n", yellow, formatter("WARN", time, location, message), reset)
}

// Error is the function which will output the Error message to the output logging interface
// @arg tag: a string which represents identification for the log message
// @arg location: the location of the error in the project
// @arg message: the message which will be printed
func (f *file) error(time string, location string, message string) {
	if !f.isValidFileProperities() {
		return
	}
	// TODO: replace& implement writing a Error in the file
	fmt.Printf("%s %s %s\n", red, formatter("ERROR", time, location, message), reset)
}

func (f *file) isValidFileProperities() bool {

	if len(f.filePath) <= 0 {
		return false
	}

	if len(f.filename) <= 0 {
		return false
	}

	return true
}
