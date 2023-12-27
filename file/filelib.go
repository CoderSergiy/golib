package file

import "fmt"
import "os"


func GetFilePointer (LogFileName string) (*os.File){

    if FileExists (LogFileName) == false {
    	fp, err := CreateFile (LogFileName)
    	if err != nil {
    		fmt.Printf("Error to creata file '%s' , err: '%s'\n", LogFileName, err)
    		return nil
    	}
    	return fp
    }

	fp, err := OpenFile (LogFileName)
	if err != nil {
		fmt.Printf("Error to open file '%s' , err: '%s'\n", LogFileName, err)
		return nil
	}

    return fp

}

func FileExists(name string) bool {
    if _, err := os.Stat(name); err != nil {
       if os.IsNotExist(err) {
            return false
        }
    }
    return true
}

func CreateFile(name string) (*os.File, error) {
    fo, err := os.Create(name)
    if err != nil {
        return nil, err
    }
    defer func() {
        fo.Close()
    }()
    return fo, nil
}

func OpenFile(filename string) (*os.File, error) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
	    return nil, err
	}

	defer func() {
		f.Close()
	}()

	return f, nil
}

func OpenCreateFile(filename string) (*os.File) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Printf("Error to open/create file '%s' , err: '%s'\n", filename, err)
	    return nil
	}

	defer func() {
		//f.Close()
	}()

	return f
}