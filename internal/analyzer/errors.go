package analyzer

import (
	"fmt"
)

type FileNotFoundError struct {
	FilePath string
	Err      error
}

func (e *FileNotFoundError) Error() string {
	return fmt.Sprintf("fichier introuvable ou inaccessible: %s", e.FilePath)
}

func (e *FileNotFoundError) Unwrap() error {
	return e.Err
}

func NewFileNotFoundError(filePath string, err error) *FileNotFoundError {
	return &FileNotFoundError{
		FilePath: filePath,
		Err:      err,
	}
}

type ParseError struct {
	LogID   string
	Message string
	Err     error
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("erreur de parsing pour le log %s: %s", e.LogID, e.Message)
}

func (e *ParseError) Unwrap() error {
	return e.Err
}

func NewParseError(logID, message string, err error) *ParseError {
	return &ParseError{
		LogID:   logID,
		Message: message,
		Err:     err,
	}
} 