package main

import (
	"fmt"
	"strconv"
	"strings"
)

func munculSekali(angka string) []int {
	strs := strings.Split(angka, "")
	ary := make([]int, len(strs))
	for i := range ary {
		ary[i], _ = strconv.Atoi(strs[i])
	}

	dict := make(map[int]int)
	for _, v := range ary {
		//fmt.Println(v)
		dict[v]++
	}

	res := make([]int, 0)
	for key, element := range dict {
		// fmt.Println("Key:", key, "=>", "Element:", element)

		if element == 1 {
			res = append(res, key)
		}
	}

	return res
}

func main() {
	fmt.Println(munculSekali("1234123"))    // [4]
	fmt.Println(munculSekali("76523752"))   // [6 3]
	fmt.Println(munculSekali("12345"))      // [1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) // []
	fmt.Println(munculSekali("0872504"))    // [8 7 2 5 4]
}
