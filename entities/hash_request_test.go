package entities_test

import (
	"testing"

	"github.com/Jaytpa01/encryption-cli/entities"
	"github.com/Jaytpa01/encryption-cli/entities/flags"
	"github.com/stretchr/testify/assert"
)

func TestValidateAndSetDefaults(t *testing.T) {
	tests := []struct {
		hashRequest       entities.HashRequest
		expectedAlgorithm flags.Algorithm
		expectedEncoding  flags.Encoding
		errorExpected     bool
		description       string
	}{
		{
			hashRequest: entities.HashRequest{
				Algorithm: flags.BCRYPT,
				Encoding:  flags.BASE64,
			},

			expectedAlgorithm: flags.BCRYPT,
			expectedEncoding:  flags.BASE64,
			errorExpected:     true,
			description:       "providing algorith and encoding, therefore defaults shouldn't be set. an error is expected as no text or file is supplied",
		},

		{
			hashRequest: entities.HashRequest{
				TextToHash: "some text",
			},

			expectedAlgorithm: flags.SHA256,
			expectedEncoding:  flags.HEX,
			errorExpected:     false,
			description:       "no algorithm or encoding supplied, so defaults should be set. no expected error",
		},
	}

	for _, test := range tests {

		err := test.hashRequest.ValidateAndSetDefaults()
		if test.errorExpected {
			assert.NotNil(t, err, test.description)
		} else {
			assert.Nil(t, err, test.description)
		}

		assert.Equal(t, test.expectedAlgorithm, test.hashRequest.Algorithm)

		assert.Equal(t, test.expectedEncoding, test.hashRequest.Encoding)

	}
}
