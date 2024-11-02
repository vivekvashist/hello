package main // package declaration

// single line comment

/* multi
line
comment
*/

import (
	"fmt" // std lib package
)

func main() {
	fmt.Printf("%s\n", "lets go")
	s := "hello"
	fmt.Printf("%s\n", s)
	fmt.Printf("%T\n", s)
	b := s[0]
	fmt.Printf("%v\n", b)
	fmt.Printf("%T\n", b)
	fmt.Printf("%d\n", len(s))
	s1 := "😃 💁"
	fmt.Printf("%s\n", s1)
	s2 := s1[1:]
	fmt.Printf("%s\n", s2)
	s3 := s1[2:]
	fmt.Printf("%s\n", s3)
	fmt.Printf("%s, %s\n", s2, s3)
	fmt.Printf("%d\n", len(s1))

	var r rune = 1000 // int32 (4 bytes)
	fmt.Printf("%d\n", r)
	fmt.Printf("%T\n", r)

	var values [3]int
	values[0] = 2
	values[1] = 4
	values[2] = 4

	fmt.Printf("%v\n", values)
	fmt.Printf("%T\n", values)
	fmt.Printf("%d\n", len(values))
	fmt.Printf("%d\n", cap(values))

	var routers [3]string
	routers[0] = "cor1.pop1"
	routers[1] = "cor2.pop1"
	routers[2] = "cor3.pop1"

	fmt.Printf("%T\n", routers)
	fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	router := "bdr1.syd1"
	firstLetter := router[0]
	fmt.Printf("%b\n", firstLetter) // prints uint8=98
	fmt.Printf("%d\n", firstLetter) // prints 98 i.e ascii "b"
	fmt.Printf("%s\n", string(firstLetter))

	fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	// string is a sequence of bytes
	hostName := "sw1.pop1"
	// fmt.Printf("%s\n", hostName)
	// runehostName := []rune(hostName)
	// fmt.Printf("%d\n", runehostName) // prints [115 119 49 46 112 111 112 49]
	// bytehostName := []byte(hostName)
	// fmt.Printf("%s\n", bytehostName) // prints sw1.pop1
	// fmt.Printf("%T\n", hostName)
	// fmt.Printf("%T\n", runehostName)
	// r1 := rune(1111)
	// fmt.Printf("%T\n", r1)
	// fmt.Printf("%d\n", r1)
	fmt.Printf("%s\n", hostName)    // prints sw1.pop1
	fmt.Printf("%T\n", hostName)    // type string
	fmt.Printf("%d\n", hostName[0]) // 115
	fmt.Printf("%T\n", hostName[0]) // unint8
	for pos, char := range hostName {
		fmt.Printf("%d -> %d\n", pos, rune(char))
		fmt.Printf("%d -> %s\n", pos, string(char))
		fmt.Printf("%d -> %d\n", pos, byte(char))
	}
	// prints
	// 0-> 115
	// 1 -> 119
	// 2 -> 49
	// 3 -> 46
	// 4 -> 112
	// 5 -> 111
	// 6 -> 112
	// 7 -> 49

	fmt.Printf("%s\n", hostName) // prints sw1.pop1
	fmt.Printf("%d\n", []rune(hostName))
	fmt.Printf("%d\n", []byte(hostName))
}
