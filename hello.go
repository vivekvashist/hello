package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// code from https://go.dev/blog/strings
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	fmt.Println("Println:")
	fmt.Println(sample)
	fmt.Printf("\n")
	fmt.Println("Byte loop:")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[1])
	}
	fmt.Printf("\n")
	fmt.Printf("\n")
	fmt.Println("Printf with %x:")
	fmt.Printf("%x\n", sample)
	fmt.Printf("\n")
	fmt.Println("Printf with % x:")
	fmt.Printf("% x\n", sample)
	fmt.Printf("\n")
	fmt.Println("Printf with %q:")
	fmt.Printf("%q\n", sample)
	fmt.Printf("\n")
	fmt.Println("Printf with %+q:")
	fmt.Printf("%+q\n", sample)
	fmt.Printf("\n")

	const placeOfInterest = `⌘`
	fmt.Printf("plain string: ")
	fmt.Printf("%s", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("quoted string: ")
	fmt.Printf("%+q", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("Hex bytes: ")
	for i := 0; i < len(placeOfInterest); i++ {
		fmt.Printf("%x ", placeOfInterest[i])
	}
	fmt.Printf("\n")

	const nihongo = "日本語"
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}

	for i := 0; i < len(nihongo); i++ {
		fmt.Printf("% x", nihongo[i])
	}

	fmt.Printf("\n")
	fmt.Printf("\n")

	for i, w := 0, 0; i < len(nihongo); i += w {
		runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
		fmt.Printf("%#U starts at byte postion %d\n", runeValue, i)
		w = width
	}
}
