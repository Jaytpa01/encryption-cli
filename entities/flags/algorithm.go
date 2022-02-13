// algorithm.go defines the allowed values of the --algorithm flag
// used in the "hash" command

package flags

import (
	"fmt"
)

type Algorithm string

const (
	SHA256 Algorithm = "sha"
	BCRYPT Algorithm = "bcrypt"
)

func (a *Algorithm) String() string {
	return string(*a)
}

func (a *Algorithm) Set(value string) error {

	switch value {
	case string(SHA256), string(BCRYPT):
		*a = Algorithm(value)
		return nil

	default:
		return fmt.Errorf(`algorithm only accepts values: "%s", "%s"`, SHA256, BCRYPT)
	}
}

func (a *Algorithm) Type() string {
	return "Algorithm"
}

func (a *Algorithm) SetDefaultValue() {
	*a = SHA256
}
