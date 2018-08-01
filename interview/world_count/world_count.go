package main

import (
	"fmt"
	"sort"
	"strings"
)

type sortedMap struct {
	m map[string]int
	s []string
}

func (sm *sortedMap) Len() int {
	return len(sm.m)
}

func (sm *sortedMap) Less(i, j int) bool {
	if sm.m[sm.s[i]] == sm.m[sm.s[j]] {
		return sm.s[i] < sm.s[j]
	}
	return sm.m[sm.s[i]] > sm.m[sm.s[j]]
}

func (sm *sortedMap) Swap(i, j int) {
	sm.s[i], sm.s[j] = sm.s[j], sm.s[i]
}

func sortedKeys(m map[string]int) []string {
	sm := new(sortedMap)
	sm.m = m
	sm.s = make([]string, len(m))
	i := 0
	for key, _ := range m {
		sm.s[i] = key
		i++
	}
	sort.Sort(sm)
	return sm.s
}

func print_word(words map[string]int) {
	for _, word := range sortedKeys(words) {
		fmt.Println(word, words[word])
	}
}

func world_count(text string) map[string]int {
	words := make(map[string]int)
	for _, field := range strings.Split(text, " ") {
		words[strings.ToLower(field)]++
	}
	return words
}

func main() {
	text := "This is a big apple tree I love big big apple 42"
	words := world_count(text)
	print_word(words)
}
