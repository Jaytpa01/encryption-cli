package entities

import (
	"errors"

	"github.com/Jaytpa01/encryption-cli/entities/flags"
)

type HashRequest struct {
	Algorithm      flags.Algorithm
	Encoding       flags.Encoding
	TextToHash     string
	FileToHash     string
	OutputFilePath string
}

func (r *HashRequest) ValidateAndSetDefaults() error {
	if err := r.validate(); err != nil {
		return err
	}

	if r.Algorithm == "" {
		r.Algorithm.SetDefaultValue()
	}

	if r.Encoding == "" {
		r.Encoding.SetDefaultValue()
	}

	return nil
}

func (r *HashRequest) validate() error {
	// if both a text argument AND a file to hash has been
	// supplied throw an error
	// cobra should throw an error when this happens, but we
	// may as well check here
	if r.TextToHash != "" && r.FileToHash != "" {
		return errors.New("a text argument and a file to hash were both supplied. please only select one")
	}

	// if both fields empty, throw an error
	if r.TextToHash == "" && r.FileToHash == "" {
		return errors.New("no supplied arguments to hash. supply either a string argument or a file to hash")
	}

	return nil
}
