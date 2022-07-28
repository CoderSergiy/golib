package logging

import (
	"fmt"
	"testing"
)

/****************************************************************************************
 *
 * Function : TestLogStruct
 *
 *  Purpose : Method to test CallMessage struct as CentralSystem
 *
 *   Return : True when generated request and response is parsed correctly, false otherwise
 */
func TestLogStruct(t *testing.T) {

	log := LogConstructor("/tmp/logs", false)

	// Set max log file size by default
	if log.fileSize != 5000000 {
		t.Error(fmt.Printf("Wrong maxFileSize : '%v'", log.fileSize))
	} else {
		t.Log(fmt.Printf("maxFileSize is correct '%v'", log.fileSize))
	}

	duplicateOnTerminal := true
	log1 := LogConstructor("/tmp/logs", duplicateOnTerminal, 125000)

	// Check if max log filesize is not set by default and equal to set size in constructor
	if log1.fileSize != 125000 {
		t.Error(fmt.Printf("Wrong maxFileSize : '%v'", log1.fileSize))
	} else {
		t.Log(fmt.Printf("maxFileSize is correct '%v'", log1.fileSize))
	}

	// Check if flag set correctly in constructor
	if log1.duplicateOnTerminal == duplicateOnTerminal {
		t.Log(fmt.Printf("duplicateOnTerminal flag is set correctly: '%v'", duplicateOnTerminal))
	} else {
		t.Error(fmt.Printf("Wrong duplicateOnTerminal flag: '%v'", duplicateOnTerminal))
	}
}
