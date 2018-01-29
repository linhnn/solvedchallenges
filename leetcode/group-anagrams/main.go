//
// https://leetcode.com/problems/group-anagrams/description/
// https://play.golang.org/p/YDOD8LyE1sK
package main

import (
	"bytes"
	"fmt"
)

func main() {
	in := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(in))

	// In:       ["cab", "tin", "pew", "duh", "may", "ill", "buy", "bar", "max", "doc"]
	// Expected: [["doc"],["bar"],["buy"],["ill"],["may"],["tin"],["cab"],["pew"],["max"],["duh"]]

	in = []string{"cab", "tin", "pew", "duh", "may", "ill", "buy", "bar", "max", "doc"}
	fmt.Println(groupAnagrams(in))

	in = []string{"tao", "pit", "cam", "aid", "pro", "dog"}
	fmt.Println(groupAnagrams(in))

	fmt.Printf("%#v\n", groupAnagrams([]string{"", "cab", "", "", "abc"}))
}

func groupAnagrams(strs []string) [][]string {
	hashes := make(map[string][]string, len(strs)/2)
	for _, s := range strs {
		hash := h(s)
		hashes[hash] = append(hashes[hash], s)
	}

	res := make([][]string, 0, len(hashes))
	for _, strs := range hashes {
		res = append(res, strs)
	}

	return res
}

func h(s string) string {
	var b bytes.Buffer
	if s == "" {
		return b.String()
	}

	var alpha [26]int
	for ix := 0; ix < len(s); ix++ {
		alpha[s[ix]-'a']++
	}

	for n, rep := range alpha {
		if rep == 0 {
			continue
		}
		b.WriteString(fmt.Sprintf("%d%d", rep, n))
	}
	return b.String()
}
