// main.go
package q001

import (
	"fmt"
	"io"
	"unicode/utf8"
)

// DemoStringVRunesWithWriter demonstrates Strings v RUNES using a specified writer.
func DemoStringVRunesWithWriter(w io.Writer) {
	const s = "สวัสดี"
	fmt.Fprintln(w, "Len:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Fprintf(w, "%x ", s[i])
	}
	fmt.Fprintln(w)

	fmt.Fprintln(w, "Rune count:", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Fprintf(w, "%#U starts at %d\n", runeValue, idx)
	}

	fmt.Fprintln(w, "\nUsing DecodeRuneInString")
	for i := 0; i < len(s); {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Fprintf(w, "%#U starts at %d\n", runeValue, i)
		i += width

		examineRune(w, runeValue)
	}

}

func examineRune(w io.Writer, r rune) {
	if r == 't' {
		fmt.Fprintln(w, "found tee")
	} else if r == 'ส' {
		fmt.Fprintln(w, "found so sua")
	}
}
