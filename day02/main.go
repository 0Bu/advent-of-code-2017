package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	str := string(input)
	fmt.Printf("checksum is %d\n", checksum(str))
	fmt.Printf("sum of divisibles is %d\n", divsum(str))
}

func checksum(s string) (sum int) {
	for _, line := range strings.Split(s, "\n") {
		ints := str2int(strings.Split(line, "\t"))
		sort.Ints(ints)
		sum += ints[len(ints)-1] - ints[0]
	}
	return
}

func str2int(s []string) (a []int) {
	a = make([]int, len(s))
	for i, r := range s {
		a[i], _ = strconv.Atoi(string(r))
	}
	return
}

func divsum(s string) (sum int) {
	for _, line := range strings.Split(s, "\n") {
		ints := str2int(strings.Split(line, "\t"))
		sum += divisible(ints)
	}
	return
}

func divisible(ints []int) int {
	sort.Ints(ints)
	for i := 0; i < len(ints)-1; i++ {
		for j := i + 1; j < len(ints); j++ {
			if ints[j]%ints[i] == 0 {
				return ints[j] / ints[i]
			}
		}
	}
	panic("no divisibles found")
}
