package jsonschema

import "fmt"

type parseError struct {
	// path is where the error occurred
	path Path
	msg string
}

func (e parseError) Error() string {
	return fmt.Sprintf("%s: %s", e.msg, e.path)
}

func ParseError(path Path, msg string) error {
	return parseError{
		path: path,
		msg:  msg,
	}
}

func ParseErrorf(path Path, format string, a ...interface{}) error {
	return parseError{
		path: path,
		msg:  fmt.Sprintf(format, a...),
	}
}
