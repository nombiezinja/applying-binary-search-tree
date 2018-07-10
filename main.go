package main

import (
	"fmt"

	bt "github.com/nombiezinja/binarytree"
)

type Bag struct {
	Colour string
}

func main() {
	// Start new tree
	things := bt.NewTree()

	// Add key value to tree
	things.Set(bt.FloatKey(2.25), []Bag{{Colour: "red"}, {Colour: "black"}, {Colour: "grey"}})
	things.Set(bt.FloatKey(1.25), []Bag{{Colour: "blue"}, {Colour: "purple"}})
	things.Set(bt.FloatKey(3.25), []Bag{{Colour: "green"}, {Colour: "aqua"}})
	things.Set(bt.FloatKey(4.25), []Bag{{Colour: "yellow"}})

	var i bt.Iterator = func(key bt.Comparable, value interface{}) {
		fmt.Println(key, value)
	}

	// Output:
	// 1.25 [{blue} {purple}]
	// 2.25 [{red} {black} {grey}]
	// 3.25 [{green} {aqua}]
	// 4.25 [{yellow}]

	things.Walk(i, true)

	// Output:
	// 4.25 [{yellow}]
	// 3.25 [{green} {aqua}]
	// 2.25 [{red} {black} {grey}]
	// 1.25 [{blue} {purple}]
	things.Walk(i, false)

	// Output:
	// Biggest:1.25 [{blue} {purple}]
	// Smallest:4.25 [{yellow}]
	fmt.Printf("Biggest:")
	fmt.Println(things.First())
	fmt.Printf("Smallest:")
	fmt.Println(things.Last())

	// Output:
	// true [{red} {black} {grey}]
	fmt.Println(things.Get(bt.FloatKey(2.25)))

	// Delete key from tree
	// Output:
	// 1.25 [{blue} {purple}]
	// 3.25 [{green} {aqua}]
	// 4.25 [{yellow}]
	things.Clear(bt.FloatKey(2.25))
	things.Walk(i, true)

	// Output:
	// 1.25 [{aquamarine} {maroon}]
	// 3.25 [{green} {aqua}]
	// 4.25 [{yellow}]
	// 5.25 [{red} {red}]
	// 6.25 [{white} {lemon}]
	// 7.25 [{brown} {mint}]
	// 9.25 [{purple} {yellow}]
	// 12.25 [{orange}]
	things.Set(bt.FloatKey(1.25), []Bag{{Colour: "aquamarine"}, {Colour: "maroon"}})
	things.Set(bt.FloatKey(5.25), []Bag{{Colour: "red"}, {Colour: "red"}})
	things.Set(bt.FloatKey(6.25), []Bag{{Colour: "white"}, {Colour: "lemon"}})
	things.Set(bt.FloatKey(0.25), []Bag{})
	things.Set(bt.FloatKey(9.25), []Bag{{Colour: "purple"}, {Colour: "yellow"}})
	things.Set(bt.FloatKey(12.25), []Bag{{Colour: "orange"}})
	things.Set(bt.FloatKey(7.25), []Bag{{Colour: "brown"}, {Colour: "mint"}})
	things.Walk(i, true)

	// Return value from tree by key
	// Output: &{<nil> 0xc420060270 5.25 [{red} {red}]}
	me := things.GetNode(bt.FloatKey(5.25))
	fmt.Println(me)
}
