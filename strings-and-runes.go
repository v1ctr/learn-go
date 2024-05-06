package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	const s = "สวัสดี"

	fmt.Println("Len:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// To count how many runes are in a string, we can use the utf8 package.
	// Note that the run-time of RuneCountInString depends on the size of the string,
	// because it has to decode each UTF-8 rune sequentially.
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// A range loop handles strings specially and decodes each rune along with its offset in the string.
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	// We can achieve the same iteration by using the utf8.DecodeRuneInString function explicitly.
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

func examineRune(r rune) {

	//Values enclosed in single quotes are rune literals.
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
