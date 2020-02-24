package main

import (
	"fmt"
)

type node struct {
	Row int
	Col int
	Val int
}

func main() {
	var chess [11][11]int
	var nodeSlice []node
	chess[1][2] = 1
	chess[2][3] = 2

	item0 := node{
		Row: 11,
		Col: 11,
		Val: 0,
	}
	nodeSlice = append(nodeSlice, item0)

	for i, v := range chess {
		for j, v1 := range v {
			if v1 != 0 {
				item := node{
					Row: i,
					Col: j,
					Val: v1,
				}
				nodeSlice = append(nodeSlice, item)
			}
		}

	}

	var res [11][11]int

	for i, v := range nodeSlice {
		if i == 0 {
			continue
		} else {
			res[v.Row][v.Col] = v.Val
		}
	}
	fmt.Println(len(res))

}
