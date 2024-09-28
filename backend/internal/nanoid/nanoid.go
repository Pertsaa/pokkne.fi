package nanoid

import (
	"errors"
	"strings"

	nanoidgen "github.com/matoous/go-nanoid/v2"
)

const (
	alphabet = "0123456789abcdefghijklmnopqrstuvwxyz"
	length   = 12
)

func New() (string, error) { return nanoidgen.Generate(alphabet, length) }

func Must() string { return nanoidgen.MustGenerate(alphabet, length) }

func Validate(id string) error {
	if id == "" {
		return errors.New("public ID is required")
	}

	if len(id) != length {
		return errors.New("public ID must be 12 characters long")
	}

	if strings.Trim(id, alphabet) != "" {
		return errors.New("invalid character(s) in public ID")
	}

	return nil
}
