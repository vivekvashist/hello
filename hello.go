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
}
