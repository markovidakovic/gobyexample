package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// methos of base also become methods of container since it embeds base
type container struct {
	base
	str string
}

func main() {
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	fmt.Println("also num:", co.base.num)
	fmt.Println("describe:", co.describe())

	// embedding structs with methods may be used to bestow interface implementations onto other structs.
	// here we see that a container now implements the describer interface because it embeds base
	type describe interface {
		describe() string
	}
	var d describe = co
	fmt.Println("describer:", d.describe())
}
