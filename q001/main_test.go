// main_test.go
package q001

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDemoStringVRunesWithWriter(t *testing.T) {
	var buf bytes.Buffer
	DemoStringVRunesWithWriter(&buf)

	expectedOutput := "Len: 12\n" +
		"e0 b8 aa e0 b8 a7 e0 b8 b1 e0 b9 80 \n" +
		"Rune count: 4\n" +
		`U+0E2B 'ส' starts at 0` + "\n" +
		`U+0E27 'ว' starts at 3` + "\n" +
		`U+0E31 'ั' starts at 6` + "\n" +
		`U+0E40 'ด' starts at 9` + "\n" +
		"\nUsing DecodeRuneInString\n" +
		`U+0E2B 'ส' starts at 0` + "\n" +
		`U+0E27 'ว' starts at 3` + "\n" +
		`U+0E31 'ั' starts at 6` + "\n" +
		`U+0E40 'ด' starts at 9` + "\n" +
		"found so sua\n"

	// Compare the output
	assert.Equal(t, expectedOutput, buf.String())
}
