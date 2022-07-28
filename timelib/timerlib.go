/*	==========================================================================
	Golib Repo
	Filename: timerlib.go
	Owner: Sergiy Safronov
	Purpose: Define EventTimer struct with print methods
	=============================================================================
*/

package timelib

import (
	"fmt"
	"time"
)

/****************************************************************************************
 *	Struct 	: EventTimer
 *
 * 	Purpose : Object handles timer for the process
 *
*****************************************************************************************/
type EventTimer struct {
	time time.Time
}

/****************************************************************************************
 *
 * Function : EventTimerConstructor (Constructor)
 *
 *  Purpose : Creates a new instance of the EventTimer object
 *
 *	Return : EventTimer object
 */
func EventTimerConstructor() EventTimer {
	eventTimer := EventTimer{}
	eventTimer.init()
	return eventTimer
}

/****************************************************************************************
 *
 * Function : EventTimer::init
 *
 * Purpose : Set default values for the EventTimer instance
 *
 *   Input : Nothing
 *
 *  Return : Nothing
 */
func (eventTimer *EventTimer) init() {
	eventTimer.time = time.Now()
}

/****************************************************************************************
 *
 * Function : EventTimer::PrintTimerString
 *
 * Purpose : Print timer value
 *
 *   Input : Nothing
 *
 *  Return : Timer value in text format for logging
 */
func (eventTimer *EventTimer) PrintTimerString() string {
	duration := time.Now().Sub(eventTimer.time)
	if int(duration.Hours()) > 0 {
		return fmt.Sprintf("%f hurs %f minutes %f seconds", duration.Hours(), duration.Minutes(), duration.Seconds())
	} else if int(duration.Minutes()) > 0 {
		return fmt.Sprintf("%f minutes %f seconds", duration.Minutes(), duration.Seconds())
	} else if duration.Seconds() > 0 {
		return fmt.Sprintf("%f seconds", duration.Seconds())
	}
	return fmt.Sprintf("%d Nanoseconds", duration.Nanoseconds())
}
