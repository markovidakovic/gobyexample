package main

import (
	"fmt"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	d, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(d)

	k, _ := strconv.Atoi("135") // convenience function for basic base-10 int parsing
	fmt.Println(k)

	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
