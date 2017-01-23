package cfg

import (
	"encoding/json"
	"io/ioutil"

	cerr "github.com/go-hayden-base/err"
	fs "github.com/go-hayden-base/fs"
)

// ConfigFromJSON read json file to generate config
func ConfigFromJSON(filePath string, v interface{}) cerr.Err {
	if len(filePath) == 0 {
		return cerr.NewErrMessage(cerr.ErrCodeParamInvalid, "Parameter filePath can not be empty!")
	}
	if !fs.FileExists(filePath) {
		return cerr.NewErrMessage(cerr.ErrCodeFileNoSuchFile, "No such file ["+filePath+"]")
	}

	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return cerr.NewErr(cerr.ErrCodeUnknown, err)
	}

	if err := json.Unmarshal(bytes, v); err != nil {
		return cerr.NewErr(cerr.ErrCodeUnknown, err)
	}
	return nil
}
