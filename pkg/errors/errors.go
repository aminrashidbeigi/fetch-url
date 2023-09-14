package errors

import "errors"

var (
	ErrInvalidURL      = errors.New("invalid URL")
	ErrFetchError      = errors.New("error fetching web page")
	ErrCreateFileError = errors.New("error creating file")
	ErrSaveFileError   = errors.New("error saving file")
)
