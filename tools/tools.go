/*	==========================================================================
	Golib Repo
	Filename: tools.go
	Owner: Sergiy Safronov
	Source : github.com/CoderSergiy/golib/tools
	Purpose: Utils methods
	=============================================================================
*/

package tools

/****************************************************************************************
 *
 * Function : EnsureSlashInEnd
 *
 *  Purpose : Ensure slesh is presents in the end of string
 *
 *    Input : url string - string to check
 *
 *	 Return : string with slesh in the end
 */
func EnsureSlashInEnd(url string) string {

	if len(url) == 0 {
		return (url)
	}

	if url[len(url)-1] == '/' {
		return url
	}

	return (url + "/")
}

/****************************************************************************************
 *
 * Function : Implode
 *
 *  Purpose : Convert string array to one line with separators
 *
 *    Input : array []string - income string
 *			  separator ...byte - set separator if need it (',' by default)
 *
 *	 Return : string
 */
func Implode(array []string, separator ...byte) string {

	var sep byte = ','
	if len(separator) > 0 {
		sep = separator[0]
	}

	var output string
	for _, word := range array {
		if len(output) > 0 {
			output = output + string(sep)
		}
		output = output + word
	}
	return output
}
