package flags

import "fmt"

type Encoding string

const (
	HEX    Encoding = "hex"
	BASE64 Encoding = "base64"
	BINARY Encoding = "binary"
)

func (e *Encoding) String() string {
	return string(*e)
}

func (e *Encoding) Set(value string) error {
	switch value {
	case string(HEX), string(BASE64), string(BINARY):
		*e = Encoding(value)
		return nil

	default:
		return fmt.Errorf(`encoding only accepts values: "%s", "%s", "%s"`, HEX, BASE64, BINARY)
	}
}

func (e *Encoding) Type() string {
	return "Encoding"
}

func (e *Encoding) SetDefaultValue() {
	*e = HEX
}
