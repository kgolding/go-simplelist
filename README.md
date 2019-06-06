# go-simplelist

Simple, go routine safe list of strings with Add, Remove, Set methods and Exists function.

````
package main

import (
	"fmt"
	"github.com/kgolding/go-simplelist"
)

const (
	APPLE  = "Apple"
	ORANGE = "Orange"
	PEAR   = "Pear"
)

func main() {
	l := simplelist.New()

	l.Add(APPLE)
	l.Add(ORANGE)
	l.Set(PEAR, false)

	if l.Exists(APPLE) {
		fmt.Println("We have an APPLE")
	}
	if l.Exists(ORANGE) {
		fmt.Println("We have an ORANGE")
	}
	if !l.Exists(PEAR) {
		fmt.Println("We do not have a PEAR")
	}
}
````

https://play.golang.org/p/_GC8vNElNEC