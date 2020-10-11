package collections

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small non-negative integers.
type IntSet struct {
	words []uint64
	size  int
}

// Init used to create an empty "words"
func (s *IntSet) Init() {
	s.words = make([]uint64, 0, 0)
	s.size = 0
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	if s.words == nil {
		s.Init()
	}
	index, bit := x/64, x%64
	return index < len(s.words) && s.words[index]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	if s.words == nil {
		s.Init()
	}
	index, bit := x/64, x%64
	if index >= len(s.words) {
		s.words = capacity(s.words, index+1)
	}
	if s.words[index]&(1<<bit) == 0 {
		s.size++
	}
	s.words[index] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for index, word := range t.words {
		if index < len(s.words) {
			s.words[index] |= word
		} else {
			s.words = append(s.words, word)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j <= 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(',')
				}
				_, err := fmt.Fprintf(&buf, "%d", 64*i+j)
				if err != nil {
					break
				}
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Len return the number of elements
func (s *IntSet) Size() int {
	return s.size
}

// Remove removes x from the set
func (s *IntSet) Remove(x int) {
	if s.words != nil && s.Has(x) {
		index, bit := x/64, x%64
		flag := ^(1<<bit)
		s.words[index] &= uint64(flag)
		s.size--
	}
}

// Remove removes all elements from the set
func (s *IntSet) Clear() {
	s.Init()
}

// Clone return a deep copy of the set
func (s *IntSet) Clone() *IntSet {
	var t = IntSet{}
	t.words = make([]uint64, cap(s.words), cap(s.words))
	copy(t.words, s.words)
	t.size = s.size
	return &t
}

func capacity(words []uint64, size int) []uint64 {
	if size <= cap(words) {
		// There is room to grow.  Extend the slice.
		return words[:size]
	} else {
		// There is insufficient space.  Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		newSize := size
		if newSize < 2*len(words) {
			newSize = 2 * len(words)
		}
		newWords := make([]uint64, newSize, newSize)
		copy(newWords, words)
		return newWords
	}
}