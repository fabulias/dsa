package main

import (
	"fmt"

	"gihtub.com/fabulias/dsa/datastructures/unionfind"
)

func main() {
	// Example usage
	uf := unionfind.New(5)
	uf.Union(0, 2)
	uf.Union(4, 2)
	uf.Union(3, 1)

	fmt.Println(uf.Find(4))
	fmt.Println(uf.Find(0) == uf.Find(2)) // true
	fmt.Println(uf.Find(3) == uf.Find(4)) // false
}
