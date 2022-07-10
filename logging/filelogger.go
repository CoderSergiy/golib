/*	==========================================================================
	Golib
	Filename: filelogger.go
	Owner: Sergiy Safronov
	Purpose: File with methods to logging messages to the file.
			Including Logfile rotation feature.
			There are 4 types of messages: INFO, EVENT, WARNING, ERROR
	=============================================================================
*/

package logging

import (
	"fmt"
	"time"
	"os"
	"errors"
)

type Message_Type string

const (
	Message_Error		Message_Type = "ERROR"
	Message_Info		Message_Type = "INFO"
	Message_Warning		Message_Type = "WARNING"
	Message_Event		Message_Type = "EVENT"

	LOG_FILE_SIZE = 5000000
)


/****************************************************************************************
 *	Struct 	: Log
 * 
 * 	Purpose : Object handles the Log methods to create log file with rotation feature
 *
*****************************************************************************************/
type Log struct {
	fileName string
	fileSize int64
}

/****************************************************************************************
 *
 * Function : LogConstructor (Constructor)
 *
 *  Purpose : Creates a new instance of the Log object
 *
 *	Return : Log object
*/
func LogConstructor (fileName string, fileSize ...int64) Log {
	log := Log{}
	log.fileName = fileName
	// Set default fileSize
	log.fileSize = LOG_FILE_SIZE

	if len(fileSize) > 0 {
		// Set filesize from constructor the parameter
		log.fileSize = fileSize[0]
	}

	return log
}

/****************************************************************************************
 *
 * Function : Log::Info_Log
 *
 * Purpose : Logging INFO type message
 *
 *   Input : messageInput string - text of the message
 *           a ...interface{} - data for the message
 *
 *  Return : Nothing
*/
func (log *Log) Info_Log (messageInput string, a ...interface{}) {
	log.createLog(fmt.Sprintf(messageInput, a...), Message_Info)
}

/****************************************************************************************
 *
 * Function : Log::Warning_Log
 *
 * Purpose : Logging WARNING type message
 *
 *   Input : messageInput string - text of the message
 *           a ...interface{} - data for the message
 *
 *  Return : Nothing
*/
func (log *Log) Warning_Log (messageInput string, a ...interface{}) {
	log.createLog(fmt.Sprintf(messageInput, a...), Message_Warning)
}

/****************************************************************************************
 *
 * Function : Log::Error_Log
 *
 * Purpose : Logging ERROR type message
 *
 *   Input : messageInput string - text of the message
 *           a ...interface{} - data for the message
 *
 *  Return : Nothing
*/
func (log *Log) Error_Log (messageInput string, a ...interface{}) {
	log.createLog(fmt.Sprintf(messageInput, a...), Message_Error)
}

/****************************************************************************************
 *
 * Function : Log::Event_Log
 *
 * Purpose : Logging EVENT type message
 *
 *   Input : messageInput string - text of the message
 *           a ...interface{} - data for the message
 *
 *  Return : Nothing
*/
func (log *Log) Event_Log (messageInput string, a ...interface{}) {
	log.createLog(fmt.Sprintf(messageInput, a...), Message_Event)
}

/****************************************************************************************
 *
 * Function : Log::createLog
 *
 * Purpose : Base function to create logging message with provided type
 *
 *   Input : messageInput string - text of the message
 *           messageType Message_Type - type of the message
 *
 *  Return : Nothing
*/
func (log *Log) createLog(messageInput string, messageType Message_Type) {
	messageInput = time.Now().Format("2006-01-02 15:04:05.000") + " [" + string(messageType) + "] " + messageInput

	// Send line to the file
	fp := log.OpenCreateFile(log.fileName)
	if fp != nil {
		fmt.Fprintf(fp, "%s\n", messageInput)
    	fp.Close()
	}

	//Check logfile for rotaton
	if err := log.logRotation(log.fileName); err != nil {
		fmt.Printf("Log Fatal error '%v'", err)
	}
}

/****************************************************************************************
 *
 * Function : Log::logRotation
 *
 *  Purpose : Rename current filename when file reached maxFileSize
 *
 *  Return : Return error if cannot perform rotation, nil otherwise
*/
func (log *Log) logRotation(fileName string) error {
	// Check if log file exist
	stat, err := os.Stat(fileName)
    if err != nil {
       return err
    }

    // Check if log file size is ready to rotate
	if stat.Size() < LOG_FILE_SIZE {
		return nil
	}

	// Prepare new file name
	counter := 1
	newFileName := ""
	for {
		newFileName = fileName + "." + fmt.Sprint(counter)

	    if _, err := os.Stat(newFileName); err != nil {
			if os.IsNotExist(err) {
				// Found correct new file name (not exists in the folder)
				break
			}
	    }
		counter++

		if counter > 10000 {
			//Something wrong with log's filename
			return errors.New(fmt.Sprintf("Cannot find name for the file '%s'", fileName))
		}
	}

	// Rename current logfile to new
	if errRename := os.Rename(fileName, newFileName); errRename != nil {
		return errors.New(fmt.Sprintf("Cannot rename file '%s' to '%s' with error '%s'", fileName, newFileName, errRename))
	}

    return nil
}

/****************************************************************************************
 *
 * Function : Log::OpenCreateFile
 *
 *  Purpose : Open file and return file pointer to the log file to add message
 *
 *   Input : filename string - name of the current log file
 *
 *  Return : Return file pointer
*/
func (log *Log) OpenCreateFile(filename string) (*os.File) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Printf("Error to open/create file '%s' , err: '%s'\n", filename, err)
	    return nil
	}

	return f
}