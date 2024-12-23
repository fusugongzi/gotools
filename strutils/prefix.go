package strutils

import (
	"github.com/fusugongzi/gotools/slices"
	"sort"
	"unicode/utf8"
)

type PrefixSearch struct {
	lens []int
	m    map[int][]string
}

func NewPrefixSearch(array []string) *PrefixSearch {
	ps := &PrefixSearch{
		lens: make([]int, 0),
		m:    make(map[int][]string, 0),
	}
	for _, s := range array {
		s := s
		sLen := utf8.RuneCountInString(s)
		if !slices.Contains(ps.lens, sLen) {
			ps.lens = append(ps.lens, sLen)
		}
		ps.m[sLen] = append(ps.m[sLen], s)
	}
	sort.SliceStable(ps.lens, func(i, j int) bool {
		return i < j
	})
	return ps
}

func (p *PrefixSearch) Search(target string) string {
	for _, l := range p.lens {
		l := l
		if len(target) < l {
			continue
		}
		targetPrefix := string([]rune(target)[0:l])
		for _, s := range p.m[l] {
			if s == targetPrefix {
				return s
			}
		}
	}
	return ""
}
