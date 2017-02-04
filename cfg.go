package cfg

import (
	"strings"
	"sync"

	"bytes"

	cerr "github.com/go-hayden-base/err"
	fs "github.com/go-hayden-base/fs"
)

var _configInstance Config
var _once sync.Once

// Config redefined map[string]string to save configuration
type Config map[string]string

// SharedConfig return a singleton Config object
func SharedConfig() Config {
	_once.Do(func() {
		_configInstance = make(Config)
	})
	return _configInstance
}

// InitWithConfigFile be used to initialize Config object by file path
func (s Config) InitWithConfigFile(filePath string) error {
	if !fs.FileExists(filePath) {
		return cerr.NewErrMessage(cerr.ErrCodeFileNoSuchFile, "Config ile '"+filePath+"' not exists!")
	}

	for k := range s {
		delete(s, k)
	}

	fs.ReadLine(filePath, func(line string, finished bool, err error, stop *bool) {
		if !isEmpty(line) {
			s.fetchKeyAndValue(line)
		}
	})
	return nil
}

// Protect Methods
func (s Config) fetchKeyAndValue(line string) {
	trimLine := strings.TrimSpace(line)
	// remove code notes
	for {
		index := strings.Index(trimLine, "#")
		if index < 0 {
			break
		}
		trimLine = trimLine[0:index]
	}
	trimLine = strings.TrimSpace(trimLine)

	lineLen := len(trimLine)

	if lineLen == 0 {
		return
	}

	eqIndex := strings.Index(trimLine, "=")
	if eqIndex < 1 || eqIndex == lineLen-1 {
		return
	}

	keyString := trimLine[0:eqIndex]
	keyString = strings.TrimSpace(keyString)
	valueString := trimLine[eqIndex+1 : lineLen]
	valueString = strings.TrimSpace(valueString)

	if isEmpty(keyString) {
		return
	}

	if isEmpty(valueString) {
		delete(s, keyString)
		return
	}

	s[keyString] = valueString
}

func (s Config) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("{\n")
	for key, val := range s {
		buffer.WriteString("  " + key + ": " + val + "\n")
	}
	buffer.WriteString("}\n")
	return buffer.String()
}

func isEmpty(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}
