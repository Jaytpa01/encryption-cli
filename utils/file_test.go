package utils_test

import (
	"testing"

	"github.com/Jaytpa01/encryption-cli/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetFileNameWithoutExtension(t *testing.T) {
	tests := []struct {
		inputFilePath    string
		expectedFileName string
	}{
		{
			inputFilePath:    "thisis/atest.exe",
			expectedFileName: "atest",
		},
		{
			inputFilePath:    "deargod pleasework.go",
			expectedFileName: "deargod pleasework",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expectedFileName, utils.GetFileNameWithoutExtension(test.inputFilePath))
	}
}
