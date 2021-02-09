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
)

// Console is the type which handle the console logging functionality
// this type shall implement the Log interface
type console struct {
}

// colors which will be used to color the message based on its level
var reset = "\033[0m"
var red = "\033[31m"
var yellow = "\033[33m"
var blue = "\033[34m"

// Debug is the function which will output the Debuging message to the output logging interface
// @arg tag: a string which represents identification for the log message
// @arg location: the location of the error in the project
// @arg message: the message which will be printed
func (c *console) debug(time string, location string, message string) {
	fmt.Printf(formatter("DEBUG", time, location, message) + "\n")

}

// Info is the function which will output the Information message to the output logging interface
// @arg tag: a string which represents identification for the log message
// @arg location: the location of the error in the project
// @arg message: the message which will be printed
func (c *console) info(time string, location string, message string) {
	fmt.Printf("%s %s %s\n", blue, formatter("INFO", time, location, message), reset)
}

// Warning is the function which will output the Warning message to the output logging interface
// @arg tag: a string which represents identification for the log message
// @arg location: the location of the error in the project
// @arg message: the message which will be printed
func (c *console) warning(time string, location string, message string) {
	fmt.Printf("%s %s %s\n", yellow, formatter("WARN", time, location, message), reset)
}

// Error is the function which will output the Error message to the output logging interface
// @arg tag: a string which represents identification for the log message
// @arg location: the location of the error in the project
// @arg message: the message which will be printed
func (c *console) error(time string, location string, message string) {

	fmt.Printf("%s %s %s\n", red, formatter("ERROR", time, location, message), reset)
}

// formate the log message to be suitable to be displayed on the console
func formatter(level string, time string, location string, fullMessage string) string {
	return fmt.Sprintf("%s | %s | %s | %s", time, level, location, fullMessage)

}
