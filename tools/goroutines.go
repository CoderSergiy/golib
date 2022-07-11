/*	==========================================================================
	Golib Repo
	Filename: goroutines.go
	Owner: Sergiy Safronov
	Purpose: Methods to work with Goroutines
	=============================================================================
*/

package tools

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

/****************************************************************************************
 *
 * Function : GetGoID
 *
 *  Purpose : Get ID of gorutine
 *
 *   Input : Nothing
 *
 *  Return : Return ID in int format
*/
func GetGoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}