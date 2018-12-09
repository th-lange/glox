package scanner


import (
	"testing"
	"github.com/stretchr/testify/assert"
)



func TestScanner_Scan(t *testing.T) {
	type testStruct struct {
		name string
		input string
		expectedType TokenType
	}

	tts := GetTokenTypes()
	tests := make([]testStruct, len(tts), len(tts))

	for i, tt := range tts {
		tests[i] = testStruct{"Testing: " + tt.String, tt.String, tt.TokenTypes[0]}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scnr := Scanner{}
			scnr.Scan(tt.input)
			assert.Equal(t, scnr.Tokens[0].Type, tt.expectedType, tt.name)
		})
	}
}
