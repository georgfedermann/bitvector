package bitvectorlib

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small non-negative integers.
// The "zero value" represents the empty set.
type IntSet struct {
	words []uint64
}

// method Has reports whether the set contains the given non-negative value
func (intset *IntSet) Has(number int) bool {
	word, bit := number/64, uint(number%64)
	return word < len(intset.words) && intset.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set
func (intset *IntSet) Add(number int) {
	word, bit := number/64, uint(number%64)
	for word >= len(intset.words) {
		intset.words = append(intset.words, 0)
	}
	intset.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t
func (s *IntSet) UnionWith(t IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns a string representation of the bitvector
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
