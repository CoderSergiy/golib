package logging

import (
	"testing"
	"fmt"
//	"github.com/CoderSergiy/golib"
)

/****************************************************************************************
 *
 * Function : TestLogStruct
 *
 *  Purpose : Method to test CallMessage struct as CentralSystem
 *
 *   Return : True when generated request and response is parsed correctly, false otherwise
*/
func TestLogStruct (t *testing.T ) {

	log := LogConstructor("/tmp/logs")

	// Set max log file size by default
	if log.fileSize != 5000000 {
		t.Error (fmt.Printf("Wrong maxFileSize : '%v'", log.fileSize))
		return
	}

	t.Log(fmt.Printf("maxFileSize is correct '%v'", log.fileSize))


	log1 := LogConstructor("/tmp/logs", 125000)

	// Set max log file size by default
	if log1.fileSize != 125000 {
		t.Error (fmt.Printf("Wrong maxFileSize : '%v'", log1.fileSize))
		return
	}

	t.Log(fmt.Printf("maxFileSize is correct '%v'", log1.fileSize))
}