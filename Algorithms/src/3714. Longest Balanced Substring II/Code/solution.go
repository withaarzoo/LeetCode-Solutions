package main

import (
	"fmt"
	"strconv"
)

func key(x, y int) string {
	return strconv.Itoa(x) + "#" + strconv.Itoa(y)
}

func longestBalanced(s string) int {
	n := len(s)
	a, b, c := 0, 0, 0
	ans := 0

	// longest single-char run
	run := 0
	var prev byte = 0
	for i := 0; i < n; i++ {
		if i == 0 || s[i] != prev {
			run = 1
		} else {
			run++
		}
		prev = s[i]
		if run > ans {
			ans = run
		}
	}

	map3 := make(map[string]int)    // (b-a, c-a)
	map_ab_c := make(map[string]int) // (b-a, c)
	map_ac_b := make(map[string]int) // (c-a, b)
	map_bc_a := make(map[string]int) // (c-b, a)

	map3[key(0,0)] = 0
	map_ab_c[key(0,0)] = 0
	map_ac_b[key(0,0)] = 0
	map_bc_a[key(0,0)] = 0

	for p := 1; p <= n; p++ {
		ch := s[p-1]
		if ch == 'a' {
			a++
		} else if ch == 'b' {
			b++
		} else {
			c++
		}

		k3 := key(b - a, c - a)
		if val, ok := map3[k3]; ok {
			if p - val > ans { ans = p - val }
		} else {
			map3[k3] = p
		}

		kab := key(b - a, c)
		if val, ok := map_ab_c[kab]; ok {
			if p - val > ans { ans = p - val }
		} else {
			map_ab_c[kab] = p
		}

		kac := key(c - a, b)
		if val, ok := map_ac_b[kac]; ok {
			if p - val > ans { ans = p - val }
		} else {
			map_ac_b[kac] = p
		}

		kbc := key(c - b, a)
		if val, ok := map_bc_a[kbc]; ok {
			if p - val > ans { ans = p - val }
		} else {
			map_bc_a[kbc] = p
		}
	}

	return ans
}

// quick test
func main() {
	fmt.Println(longestBalanced("abbac"))  // 4
	fmt.Println(longestBalanced("aabcc")) // 3
	fmt.Println(longestBalanced("aba"))   // 2
}
