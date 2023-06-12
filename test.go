package main

import "fmt"

func masin() {
	tes := [5]int{3,5,7,8,9,}

	fmt.Println(bins(tes, 8))

}

func bins(t [5]int, x int) int {
	var found int = -1
	var med int

	var kr int = 0
	var kn int = len(t) - 1

	for kr <= kn && found == -1 {
		med = (kr + kn) / 2
		fmt.Println(t[med])
		if x < t[med] {
			kn = med - 1
		} else if x > t[med] {
			kr = med + 1
		} else{
			found = med
		}
	}

	return found
}