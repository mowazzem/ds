package main

import "github.com/mowazzem/ds/lists/singly"

type custom struct {
	name string
}

func main() {
	list := singly.NewList[custom]()
	c1 := custom{"m"}
	c2 := custom{"n"}
	list.Append(c1)
	list.Append(c2)
	list.Print()
}
