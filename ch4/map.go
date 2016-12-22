package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := make(map[string]int)
	ages["alice"] = 31
	ages["charlie"] = 34

	ages = map[string]int{
		"alice":   31,
		"charlie": 34,
	}

	delete(ages, "alice")
	ages["bob"]++

	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	// var names []string
	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

	var ages1 map[string]int
	fmt.Println(ages1 == nil)
	fmt.Println(len(ages1) == 0)
	ages1["carol"] = 21 //panic: assignment to entry in nil map

	age, ok := ages["bob"]
	if !ok {
		/* "bob" is not a key in this map; age == 0. */
	}

	if age, ok := ages["bob"]; !ok {
		/*...*/
	}

	mapEqual(map[string]int{"A": 0}, map[string]int{"B": 42})
}

func mapEqual(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
