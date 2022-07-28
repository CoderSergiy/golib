/*	==========================================================================
	Golib Repo
	Filename: referencelib.go
	Owner: Sergiy Safronov
	Source : github.com/CoderSergiy/golib/tools
	Purpose: Methods to generate random reference
	=============================================================================
*/

package tools

import (
	"crypto/rand"
	"encoding/base64"
)

/****************************************************************************************
 *
 * Function : GenerateRandomBytes
 *
 *  Purpose : Generate Random Bytes for reference purpose
 *
 *    Input : n int - specify size of the buffer
 *
 *	 Return : []]byte - securely generated random bytes
 *			  error - if happened, nil otherwise
 */
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

/****************************************************************************************
 *
 * Function : GenerateRandomString
 *
 *  Purpose : Generate Random String for reference purpose
 *
 *    Input : n int - size of the returning string
 *
 *	 Return : string - securely generated random string
 *			  error - if happened, nil otherwise
 */
func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	bytes, err := GenerateRandomBytes(n)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes), nil
}

/****************************************************************************************
 *
 * Function : GenerateRandomStringURLSafe
 *
 *  Purpose : Generate URL-safe, base64 encoded random string
 *
 *    Input : n int - size of the returning string
 *
 *	 Return : string - securely generated random bytes
 *			  error - if happened, nil otherwise
 */
func GenerateRandomStringURLSafe(n int) (string, error) {
	b, err := GenerateRandomBytes(n)
	return base64.URLEncoding.EncodeToString(b), err
}
